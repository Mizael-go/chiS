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

func GenerateController(model string, directory string, fileName string) {
	controller := fmt.Sprintf(`
package controller

import (
	"encoding/json"
	"net/http"
	"%s/model"

	"github.com/go-chi/chi/v5"
)

func (c *Controller) GetAll%s(w http.ResponseWriter, r *http.Request) {
	models, err := c.R.GetAll%s()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models)
}

func (c *Controller) Get%sById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	model, err := c.R.Get%sById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

func (c *Controller) Create%s(w http.ResponseWriter, r *http.Request) {
	var model model.%s
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model, err := c.R.Create%s(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model)
}

func (c *Controller) Update%s(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var model model.%s
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.ID = id
	model, err := c.R.Update%s(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(model)
}

func (c *Controller) Delete%sById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.R.Delete%sById(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"message": "%s deleted successfully"})
}


	`, directory, model, model, model, model, model, model, model, model, model, model, model, model, model)

	controllerPath := filepath.Join("controller", fileName)
	ctrl, err := os.Create(controllerPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	ctrl.Write([]byte(controller))
	ctrl.Close()
}
