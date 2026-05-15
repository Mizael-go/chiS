package cmd

import (
	"fmt"
	"os"

	"github.com/Mizael-go/chit/handler"
	"github.com/Mizael-go/chit/service"
)

func App() {
	switch os.Args[1] {
	case "create", "-c":
		if len(os.Args) < 3 {
			fmt.Println("missing project name!")
			service.ShowHelp()
			return
		}
		projectName := os.Args[2]
		db := handler.AskDB()
		service.Create(projectName, db)
		return
	case "help", "-h":
		service.ShowHelp()
		return
	case "generate", "-g":
		service.Generate()
		return
	default:
		fmt.Printf("Invalid flag %s", os.Args[1])
		service.ShowHelp()
		return
	}
}
