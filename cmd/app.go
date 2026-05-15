package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/Mizael-go/chit/handler"
	"github.com/Mizael-go/chit/service"
)

func App() {
	pflag.BoolP("help", "h", false, "show help")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	args := pflag.Args()

	if len(args) == 0 || viper.GetBool("help") {
		service.ShowHelp()
		return
	}

	switch args[0] {
	case "create", "-c":
		if len(args) < 2 {
			fmt.Println("missing project name!")
			os.Exit(1)
		}
		projectName := args[1]
		viper.Set("project_name", projectName)

		db := handler.AskDB()
		viper.Set("db_driver", db)

		service.Create(
			viper.GetString("project_name"),
			viper.GetString("db_driver"),
		)

	case "generate", "-g":
		viper.SetConfigName(".chit")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				fmt.Println("config error:", err)
			}
		}
		service.Generate()

	case "help", "-h":
		service.ShowHelp()

	default:
		fmt.Printf("Invalid subcommand: %s\n", args[0])
		os.Exit(1)
	}
}
