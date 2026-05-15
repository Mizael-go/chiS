package lib

import (
	// "fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/Mizael-go/chit/handler"
)

type ModelInfo struct {
	Name string
	Path string
}

func GetStruct(filename string) []ModelInfo {
	fset := token.NewFileSet()
	src, err := os.ReadFile(filename)
	if err != nil {
		handler.ErrorHandler(err)
		return []ModelInfo{}
	}
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		handler.ErrorHandler(err)
		return []ModelInfo{}
	}

	var models []ModelInfo

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

			structName := typeSpec.Name.Name
			mdl := ModelInfo{
				Name: structName,
				Path: filename,
			}
			models = append(models, mdl)
			// fmt.Printf("struct : %v\n", structName)
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
	return models
}
