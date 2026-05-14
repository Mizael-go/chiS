package templates

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mizael-go/chiS/handler"
)

func GenerateRepository(model string, directory string, fileName string) {
	repository := fmt.Sprintf(`
package repository

import (
	"%s/model"
)

func (r *Repository) GetAll%s() ([]model.%s, error) {
	var models []model.%s
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r *Repository) Get%sById(id string) (model.%s, error) {
	var model model.%s
	if err := r.DB.First(&model, id).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (r *Repository) Create%s(model model.%s) (model.%s, error) {
	if err := r.DB.Create(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (r *Repository) Update%s(model model.%s) (model.%s, error) {
	if err := r.DB.Where("id = ?", model.ID).Updates(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (r *Repository) Delete%sById(id string) error {
	var model model.%s
	if err := r.DB.Where("id = ?", id).Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

	`, directory, model, model, model, model, model, model, model, model, model, model, model, model, model, model)

	repoPath := filepath.Join("repository", fileName)
	repo, err := os.Create(repoPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	repo.Write([]byte(repository))
	repo.Close()
}
