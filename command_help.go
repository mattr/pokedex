package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Usage: command [args]")
	fmt.Println()
	fmt.Println("Available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
