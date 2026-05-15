package service

import "fmt"

func ShowHelp() {
	fmt.Println("Usage: chit <subcommand> [arguments]")
	fmt.Println()
	fmt.Println("Subcommands:")
	fmt.Println("  create     Create a new chi project")
	fmt.Println("  generate   Generate CRUD from model structs")
	fmt.Println("  help       Show this help message")
}
