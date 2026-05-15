package lib

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
)

type FIFOModel struct {
	Visited []string `json:"visited"`
}

func FIFORead(fifoPath string) ([]string, error) {
	data, err := os.ReadFile(fifoPath)
	if err != nil {
		if os.IsNotExist(err) || len(data) == 0 {
			return []string{}, nil
		}
		return nil, err
	}

	if len(strings.TrimSpace(string(data))) == 0 {
		return []string{}, nil
	}

	var models FIFOModel
	if err := json.Unmarshal(data, &models); err != nil {
		return nil, err
	}

	return models.Visited, nil
}

func FIFOAppend(fifoPath string, fifoContent string) error {
	data, err := os.ReadFile(fifoPath)
	var models FIFOModel

	if err != nil || len(data) == 0 || len(strings.TrimSpace(string(data))) == 0 {
		models = FIFOModel{Visited: []string{}}
	} else {
		if err := json.Unmarshal(data, &models); err != nil {
			return err
		}
	}

	if slices.Contains(models.Visited, fifoContent) {
		return nil
	}

	models.Visited = append(models.Visited, fifoContent)

	file, err := json.MarshalIndent(models, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fifoPath, file, 0644)
}
