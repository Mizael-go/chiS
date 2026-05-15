package service

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Mizael-go/chit/handler"
	"github.com/Mizael-go/chit/lib"
	"github.com/Mizael-go/chit/templates"
)

func Generate() {
	directory := lib.GetModuleName()
	modelDir := "model"
	fifoFile := "./generateModels.json"
	if _, err := os.Stat(fifoFile); os.IsNotExist(err) {
		initial := lib.FIFOModel{Visited: []string{}}
		data, _ := json.MarshalIndent(initial, "", "  ")
		os.WriteFile(fifoFile, data, 0644)
	}
	fifo, err := lib.FIFORead(fifoFile)
	if err != nil {
		handler.ErrorHandler(err)
		return
	}

	err = filepath.WalkDir(modelDir, func(p string, d fs.DirEntry, err error) error {

		if err != nil {
			handler.ErrorHandler(err)
			return err
		}

		if d.IsDir() || !strings.HasSuffix(strings.ToLower(p), ".go") {
			return nil
		}
		models := lib.GetStruct(p)
		for _, model := range models {
			if slices.Contains(fifo, model.Path) {
				continue
			}
			if err := lib.FIFOAppend(fifoFile, model.Path); err != nil {
				handler.ErrorHandler(err)
				return err
			}
			newPath := strings.Split(model.Path, "/")[1]
			templates.GenerateRepository(model.Name, directory, newPath)
			templates.GenerateController(model.Name, directory, newPath)
		}
		return nil
	})

	if err != nil {
		handler.ErrorHandler(err)
	}
}
