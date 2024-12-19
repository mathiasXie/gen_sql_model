package ddl2struct

import (
	"encoding/json"
	"fmt"
	ddlparser "github.com/mathiasXie/gen_sql_model/parsers"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

// CreateTableListener Thanks for ByteDance
type CreateTableListener struct {
	*ddlparser.BaseDDLParserListener

	SourceFile  string
	PackageName string
	Imports     []string
	Structs     []*GoStruct
}

type GoStruct struct {
	// option
	EnableSQLNull bool // 如果启用了 enable_sql_null，则 int64 null 会被替换为 sql.NullInt64

	// meta
	TableName    string
	TableComment string
	PrimaryKey   string
	StructName   string
	Definitions  []*Definition
}

type Definition struct {
	Name               string // Go 中字段的名称
	Type               string // Go 中字段的类型
	SQLName            string // SQL 中字段的名称
	SQLType            string // SQL 中字段的类型
	Tag                string // 计算好的 struct tag, e.g.: `gorm:"column:page_id" json:"page_id"`
	NotNull            bool
	IsInlinePrimaryKey bool
	SQLComment         string // SQL COMMENT 注释，兼容
	InlineComment      string // inline 注释（暂未实现）
}

type TagGenerator struct {
	Prefix string
	Map    func(string, bool) string
}

type TableOption struct {
	Alias         string `json:"alias"`
	EnableSQLNull bool   `json:"enable_sql_null"`
}

// EnterTableName 生成新的 GoStruct，并将 table name 和 struct name 赋值给这个新的 GoStruct
func (c *CreateTableListener) EnterTableName(ctx *ddlparser.TableNameContext) {
	nextStruct := &GoStruct{}
	nextStruct.TableName = strings.Replace(ctx.GetText(), "`", "", -1)
	nextStruct.StructName = strcase.ToCamel(strings.Replace(ctx.GetText(), "`", "", -1))

	c.Structs = append(c.Structs, nextStruct)
}

// EnterCreateDefinition 只生成 Definition 占位，子 enter 事件负责把其他数据拼装进 Definition 中
func (c *CreateTableListener) EnterCreateDefinition(ctx *ddlparser.CreateDefinitionContext) {
	if ctx.GetField() == nil { // 说明是 tableConstraint, 跳过
		return
	}

	fieldName := strings.Replace(ctx.GetField().GetText(), "`", "", -1)
	currentStruct := c.Structs[len(c.Structs)-1]

	currentStruct.Definitions = append(currentStruct.Definitions, &Definition{
		Name: func(s string) string {
			if s == "" {
				return s
			}
			if !strings.Contains(s, "_") {
				return strings.Title(s)
			}
			arr := strings.Split(s, "_")
			for idx, word := range arr {
				arr[idx] = strings.Title(word)
			}
			return strings.Join(arr, "")
		}(fieldName), // the original `snake2CamelWithUpperStart()`
		SQLName: fieldName,
	})
}

// EnterDataType 生成 SQLType / Type + NotNull
func (c *CreateTableListener) EnterDataType(ctx *ddlparser.DataTypeContext) {
	currentStruct := c.Structs[len(c.Structs)-1]
	currentStruct.Definitions[len(currentStruct.Definitions)-1].SQLType = strings.ToLower(ctx.GetTypeName().GetText())
}

// EnterColumnConstraint PK 定义在 columnConstraint 上
func (c *CreateTableListener) EnterColumnConstraint(ctx *ddlparser.ColumnConstraintContext) {
	if strings.EqualFold(ctx.GetText(), "PrimaryKey") {
		currentStruct := c.Structs[len(c.Structs)-1]
		currentStruct.Definitions[len(currentStruct.Definitions)-1].IsInlinePrimaryKey = true
	}
}

// EnterNullNotNull 生成 SQLType / Type + NotNull
func (c *CreateTableListener) EnterNullNotNull(ctx *ddlparser.NullNotNullContext) {
	currentStruct := c.Structs[len(c.Structs)-1]
	currentStruct.Definitions[len(currentStruct.Definitions)-1].NotNull = strings.ToLower(ctx.GetText()) == "notnull"
}

// EnterComment 获取 Content 注释内容
func (c *CreateTableListener) EnterComment(ctx *ddlparser.CommentContext) {
	lastStruct := c.Structs[len(c.Structs)-1]
	lastStruct.Definitions[len(lastStruct.Definitions)-1].SQLComment = ctx.GetContent().GetText()[1 : len(ctx.GetContent().GetText())-1]
}

// EnterTableConstraint 获取 tableConstraint 中的 PK
func (c *CreateTableListener) EnterTableConstraint(ctx *ddlparser.TableConstraintContext) {
	if ctx.GetPk() == nil {
		return
	}

	lastStruct := c.Structs[len(c.Structs)-1]
	pk := ctx.IndexColumnNames().GetText()[1 : len(ctx.IndexColumnNames().GetText())-1]
	if len(strings.Split(pk, ",")) != 1 {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Multiple columns found in PK: %s", pk))
		os.Exit(0)
	}

	// tableConstraint 中定义的 PK
	lastStruct.PrimaryKey = strings.Replace(pk, "`", "", -1)
}

func (c *CreateTableListener) ExitCreateDefinitions(ctx *ddlparser.CreateDefinitionsContext) {
}

// EnterTableOption 解析 tableOption 中的 table name alias
func (c *CreateTableListener) EnterTableOption(ctx *ddlparser.TableOptionContext) {
	if ctx.GetTableComment() == nil {
		return
	}

	comment := ctx.GetTableComment().GetText()
	startPos, endPos := strings.Index(comment, "{"), -1

	if startPos == -1 {
		return
	}

	for i := startPos + 1; i < len(comment); i++ {
		if comment[i:i+1] == "}" {
			endPos = i + 1
			break
		}
	}

	if endPos == -1 {
		return
	}

	table := &TableOption{}
	err := json.Unmarshal([]byte(comment[startPos:endPos]), table)
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("json.Unmarshal table option failed: %s", err.Error()))
		return
	}

	if len(table.Alias) > 0 {
		c.Structs[len(c.Structs)-1].StructName = table.Alias
	}
	c.Structs[len(c.Structs)-1].EnableSQLNull = table.EnableSQLNull
	c.Structs[len(c.Structs)-1].TableComment = comment
}
