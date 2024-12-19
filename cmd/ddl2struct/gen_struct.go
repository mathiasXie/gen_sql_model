package ddl2struct

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mathiasXie/gen_sql_model/parsers"
	"github.com/mathiasXie/gen_sql_model/utils/gofmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
	"unsafe"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Thanks for ByteDance
const ddl2StructTemplate = `// Code generated from [[ .SourceFile ]] by cbt. DO NOT EDIT.

package [[ .PackageName ]]

[[ if .Imports ]]
import(
[[- range .Imports ]]
	"[[ . ]]"
[[ end -]]
)
[[ end -]]

[[ range $s := .Structs ]]
func ([[ .StructName ]]) TableName() string {
	return "[[ .TableName ]]";
}

type [[ $s.StructName ]] struct {
[[- range $d := $s.Definitions ]]
	[[ $d.Name ]] [[ $d.Type ]] [[ $d.Tag ]] [[ if $d.SQLComment ]] // [[ $d.SQLComment ]] [[ end -]]
[[ end ]]
}

func New[[ .StructName ]]() *[[ .StructName ]] {
	return &[[ .StructName ]]{}
}
[[ end -]]
`

func ProcessDDL(sourceFile, packageName string) {
	if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("ddl 文件不存在: %s", sourceFile))
		os.Exit(-1)
	}

	fs, err := antlr.NewFileStream(sourceFile)
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("antlr.NewFileStream err: %s", err.Error()))
		os.Exit(-1)
	}

	p := ddlparser.NewDDLParser(antlr.NewCommonTokenStream(ddlparser.NewDDLLexer(fs), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true

	// 模板和 antlr 共用一个 model
	listener := CreateTableListener{
		SourceFile:  filepath.Base(sourceFile),
		PackageName: packageName,
	}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Root())

	// 将 SQLType 翻译成 Type
	processTypes(&listener)
	// import 去重
	processImports(&listener)
	// 字节对齐
	processStructAlignment(&listener)
	// comments 中自定义 tags
	processTags(&listener)
	// comments 中自定义 go type
	processTypeAliases(&listener)

	t, err := template.New("ddl2Struct").Delims("[[", "]]").Parse(ddl2StructTemplate)
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("初始化模板出错: %s", err.Error()))
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, listener); err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("渲染模板出错: %s", err.Error()))
		os.Exit(-1)
	}

	fmt.Println(gofmt.GoFmtInMem(buf.String()))
}

func processTypes(l *CreateTableListener) { // cbt_skip
	for _, s := range l.Structs {
		for _, d := range s.Definitions {
			if d.SQLType == "tinyint" {
				if s.EnableSQLNull && !d.NotNull {
					d.Type = "sql.NullInt64"
				} else {
					d.Type = "int8"
				}
			} else if d.SQLType == "smallint" {
				if s.EnableSQLNull && !d.NotNull {
					d.Type = "sql.NullInt64"
				} else {
					d.Type = "int16"
				}
			} else if d.SQLType == "int" || d.SQLType == "integer" || d.SQLType == "mediumint" {
				if s.EnableSQLNull && !d.NotNull {
					d.Type = "sql.NullInt64"
				} else {
					d.Type = "int32"
				}
			} else if d.SQLType == "bigint" {
				if s.EnableSQLNull && !d.NotNull {
					d.Type = "sql.NullInt64"
				} else {
					d.Type = "int64"
				}
			} else if d.SQLType == "bit" ||
				d.SQLType == "bool" ||
				d.SQLType == "boolean" {
				if s.EnableSQLNull && !d.NotNull {
					d.Type = "sql.NullInt64"
				} else {
					d.Type = "uint8"
				}
			} else if d.SQLType == "double" ||
				d.SQLType == "float" ||
				d.SQLType == "fixed" ||
				d.SQLType == "numeric" ||
				d.SQLType == "decimal" ||
				d.SQLType == "dec" {
				if s.EnableSQLNull && !d.NotNull {
					d.Type = "sql.NullFloat64"
				} else {
					d.Type = "float64"
				}
			} else if d.SQLType == "enum" ||
				d.SQLType == "set" ||
				d.SQLType == "char" ||
				d.SQLType == "character" ||
				d.SQLType == "varchar" ||
				d.SQLType == "tinytext" ||
				d.SQLType == "text" ||
				d.SQLType == "mediumtext" ||
				d.SQLType == "longtext" ||
				d.SQLType == "nchar" ||
				d.SQLType == "nvarchar" {
				if s.EnableSQLNull && !d.NotNull {
					d.Type = "sql.NullString"
				} else {
					d.Type = "string"
				}
			} else if d.SQLType == "timestamp" ||
				d.SQLType == "date" ||
				d.SQLType == "datetime" ||
				d.SQLType == "year" ||
				d.SQLType == "time" {
				d.Type = "time.Time"
			} else if d.SQLType == "binary" ||
				d.SQLType == "varbinary" ||
				d.SQLType == "tinyblob" ||
				d.SQLType == "blob" ||
				d.SQLType == "mediumblob" ||
				d.SQLType == "longblob" {
				d.Type = "[]byte"
			}
		}
	}
}

func processImports(l *CreateTableListener) {
	imports, importMapping := make([]string, 0), make(map[string]bool)
	for _, goStruct := range l.Structs {
		for _, definition := range goStruct.Definitions {
			if definition.Type == "time.Time" {
				if _, exists := importMapping["time"]; !exists {
					importMapping["time"] = true
					imports = append(imports, "time")
				}
			} else if definition.Type == "sql.NullBool" ||
				definition.Type == "sql.NullFloat64" ||
				definition.Type == "sql.NullInt64" ||
				definition.Type == "sql.NullString" {
				if _, exists := importMapping["database/sql"]; !exists {
					importMapping["database/sql"] = true
					imports = append(imports, "database/sql")
				}
			}
		}
	}
	l.Imports = imports
}

func processStructAlignment(l *CreateTableListener) {
	// 字节对齐
	sizeMap := map[string]uintptr{
		"int8":            unsafe.Sizeof(int8(0)),
		"uint8":           unsafe.Sizeof(uint8(0)),
		"int16":           unsafe.Sizeof(int16(0)),
		"uint16":          unsafe.Sizeof(uint16(0)),
		"int32":           unsafe.Sizeof(int32(0)),
		"uint32":          unsafe.Sizeof(uint32(0)),
		"int":             unsafe.Sizeof(int64(0)), // warning!!! 统一按照 64 算
		"int64":           unsafe.Sizeof(int64(0)),
		"uint64":          unsafe.Sizeof(uint64(0)),
		"float64":         unsafe.Sizeof(float64(0)),
		"sql.NullBool":    unsafe.Sizeof(sql.NullBool{}),
		"sql.NullFloat64": unsafe.Sizeof(sql.NullFloat64{}),
		"sql.NullInt64":   unsafe.Sizeof(sql.NullInt64{}),
		"sql.NullString":  unsafe.Sizeof(sql.NullString{}),
		"time.Time":       unsafe.Sizeof(time.Time{}),
		"string":          unsafe.Sizeof(""),
		"[]byte":          unsafe.Sizeof([]byte("")),
	}

	for _, s := range l.Structs {
		sort.Slice(s.Definitions, func(i, j int) bool {
			iSize, ok := sizeMap[s.Definitions[i].Type]
			if !ok {
				_, _ = os.Stderr.WriteString(fmt.Sprintf("Type not found in sizeMap: %s", s.Definitions[i].Type))
				os.Exit(-1)
			}

			jSize, ok := sizeMap[s.Definitions[j].Type]
			if !ok {
				_, _ = os.Stderr.WriteString(fmt.Sprintf("Type not found in sizeMap: %s", s.Definitions[j].Type))
				os.Exit(-1)
			}

			if iSize != jSize {
				return iSize < jSize // 先按照 size
			} else if s.Definitions[i].Type != s.Definitions[j].Type {
				return s.Definitions[i].Type < s.Definitions[j].Type // 再按照类型
			} else {
				return s.Definitions[i].Name < s.Definitions[j].Name // 最后按照 name
			}
		})
	}
}

func processTags(l *CreateTableListener) {
	for _, s := range l.Structs {
		tagGenerators := make([]*TagGenerator, 0)

		// GORM tag generator
		tagGenerators = append(tagGenerators, &TagGenerator{
			Prefix: "gorm",
			Map: func(fieldName string, isInlinePrimary bool) string {
				if isInlinePrimary || s.PrimaryKey == fieldName {
					return fmt.Sprintf("column:%s;primary_key", fieldName)
				}
				return fmt.Sprintf("column:%s", fieldName)
			},
		})

		// json tag generator
		tagGenerators = append(tagGenerators, &TagGenerator{
			Prefix: "json",
			Map: func(fieldName string, isInlinePrimary bool) string {
				return fieldName
			},
		})

		// ...其他默认的 tag generators 可以补充在这里

		// inline comment 中的特殊标志 tags:JSON 会覆盖掉默认的逻辑
		const customTagBegin = "tags:{"
		for _, definition := range s.Definitions {
			// 解析出当前 definition 中定义的 tag 覆盖规则（如果有的话）
			inlineCommentTags := make(map[string]string)
			startPos, endPos := strings.Index(definition.SQLComment, customTagBegin), -1

			if startPos >= 0 {
				for i := startPos + 1; i < len(definition.SQLComment); i++ {
					if definition.SQLComment[i:i+1] == "}" {
						endPos = i + 1
						break
					}
				}
			}

			if endPos != -1 {
				err := json.Unmarshal([]byte(definition.SQLComment[startPos+len(customTagBegin)-1:endPos]), &inlineCommentTags)
				if err != nil {
					_, _ = os.Stderr.WriteString(fmt.Sprintf("json.Unmarshal inline comment failed: %s", err.Error()))
				}
			}

			tags := make([]string, 0)
			for _, tg := range tagGenerators {
				if t, override := inlineCommentTags[tg.Prefix]; override {
					tags = append(tags, tg.Prefix+`:"`+t+`"`) // inline comment
				} else {
					tags = append(tags, tg.Prefix+`:"`+tg.Map(definition.SQLName, definition.IsInlinePrimaryKey)+`"`) // default
				}
			}

			definition.Tag = func() string {
				if len(tags) == 0 {
					return ""
				}
				return "`" + strings.Join(tags, " ") + "`"
			}()
		}
	}
}

func processTypeAliases(l *CreateTableListener) {
	// inline comment 中的特殊标志 type: 会覆盖掉默认的类型
	const customTypeBegin = "type:"

	for _, s := range l.Structs {
		for _, d := range s.Definitions {
			startPos, endPos := strings.Index(d.SQLComment, customTypeBegin), -1
			if startPos >= 0 {
				for i := startPos + 1; i < len(d.SQLComment); i++ {
					if d.SQLComment[i:i+1] == "," {
						endPos = i + 1
						break
					}
				}

				if endPos == -1 {
					endPos = len(d.SQLComment)
				}

				d.Type = d.SQLComment[startPos+len(customTypeBegin) : endPos]
			}
		}
	}
}
