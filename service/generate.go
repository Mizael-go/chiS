package service

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Mizael-go/chiS/handler"
	"github.com/Mizael-go/chiS/lib"
)

func Generate() {
	modelDir := "model"
	err := filepath.WalkDir(modelDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			handler.ErrorHandler(err)
			return err
		}

		if d.IsDir() || !strings.HasSuffix(strings.ToLower(p), ".go") {
			return nil
		}
		fmt.Println("Found go file in", p)
		lib.GetStruct(p)
		return nil
	})

	if err != nil {
		handler.ErrorHandler(err)
	}
}
