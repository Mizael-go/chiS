package service

import "slices"

import "fmt"

var ParamsArgs = []string{"-c", "create", "-h", "--help"}

func ShowHelp() {
	fmt.Println("Usage: chit [flag] [optional]")
	fmt.Println("Available flags:")
	for _, flag := range ParamsArgs {
		if flag == "-c" || flag == "create" {
			fmt.Println("  " + flag + " : create new project")
		}
		if flag == "-h" || flag == "--help" {
			fmt.Println("  " + flag + " : show all commands")
		}
	}
}

func IsAllowed(flag string) bool {
	return slices.Contains(ParamsArgs, flag)
}

