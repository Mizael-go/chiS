package templates

import (
	"os"
	"path/filepath"

	"github.com/Mizael-go/chit/handler"
)

func WriteRepository(directory string) {
	repo := `
package repository

import "gorm.io/gorm"

type Repository struct{
	DB *gorm.DB
}

func NewRepository( db *gorm.DB ) *Repository{
	return &Repository{
		DB: db,
	}
}	
	`

	repoPath := filepath.Join(directory, "repository", "repository.go")
	repository, err := os.Create(repoPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	repository.Write([]byte(repo))
	repository.Close()
}
