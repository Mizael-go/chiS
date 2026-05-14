package service

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Mizael-go/chiS/handler"
	"github.com/Mizael-go/chiS/lib"
	"github.com/Mizael-go/chiS/templates"
)

func Generate() {
	directory := lib.GetModuleName()
	modelDir := "model"
	err := filepath.WalkDir(modelDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			handler.ErrorHandler(err)
			return err
		}

		if d.IsDir() || !strings.HasSuffix(strings.ToLower(p), ".go") {
			return nil
		}
		models := lib.GetStruct(p)
		for _, model := range models {
			newPath := strings.Split(model.Path, "/")[1]
			templates.GenerateRepository(model.Name, directory, newPath) 
		}
		return nil
	})

	if err != nil {
		handler.ErrorHandler(err)
	}
}
