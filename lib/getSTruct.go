package lib

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/Mizael-go/chiS/handler"
)

func GetStruct(filename string) {
	fset := token.NewFileSet()
	src, err := os.ReadFile(filename)
	if err != nil {
		handler.ErrorHandler(err)
	}
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		handler.ErrorHandler(err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return true
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// structType, ok := typeSpec.Type.(*ast.StructType)
			// if !ok {
			// 	continue
			// }

			structName := typeSpec.Name.Name
			fmt.Printf("struct : %v\n", structName)
			// for _, field := range structType.Fields.List {
			// 	//if field have name
			// 	if field.Names != nil {
			// 		for _, name := range field.Names {
			// 			fmt.Printf("%v : %v\n", name.Name, field.Type)
			// 			// fieldType := fmt.Sprintf("%s", field.Type)
			// 			// tag := ""
			// 			// if field.Tag != nil {
			// 			// 	tag = field.Tag.Value
			// 			// }
			// 			// fmt.Printf("  - %s %s %s\n", name.Name, fieldType, tag)
			// 		}
			// 	}
			// }
		}
		return true
	})

}
