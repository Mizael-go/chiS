package lib

import (
	"fmt"
	"os"
	"strings"
)

func GetModuleName() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println("Erreur: impossible de lire go.mod")
		return ""
	}

	lines := strings.Split(string(data), "\n")
	firstLine := strings.TrimSpace(lines[0])
	module := strings.TrimPrefix(firstLine, "module")
	module = strings.TrimSpace(module)

	return module
}
