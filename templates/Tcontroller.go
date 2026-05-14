package templates

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mizael-go/chiS/handler"
)

func WriteController(directory string) {
	controller := fmt.Sprintf(`
package controller

import (
	"%s/repository"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
	R  *repository.Repository
}

func NewController(db *gorm.DB, r *repository.Repository) *Controller {
	return &Controller{
		DB: db,
		R:  r,
	}
}
	
	`, directory)

	controllerPath := filepath.Join(directory, "controller", "controller.go")
	ctrl, err := os.Create(controllerPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	ctrl.Write([]byte(controller))
	ctrl.Close()
}
