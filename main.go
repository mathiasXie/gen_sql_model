package main

import (
	"flag"
	"github.com/mathiasXie/gen_sql_model/cmd/ddl2struct"
	"os"
)

func main() {

	ddlPath := flag.String("ddlpath", "", "path/to/ddl")
	targetPackage := flag.String("package", "", "target package")
	flag.Parse()
	if ddlPath == nil || targetPackage == nil || *ddlPath == "" || *targetPackage == "" {
		_, _ = os.Stderr.WriteString("ddlpath and package should not be nil\n")
		os.Exit(-1)
		return
	}
	ddl2struct.ProcessDDL(*ddlPath, *targetPackage)
}
