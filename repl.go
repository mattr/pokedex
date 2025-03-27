package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	words := strings.Fields(lower)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help": {name: "help", description: "Show the help message", callback: commandHelp},
		"map":  {name: "map", description: "Retrieve the next page of locations", callback: commandMap},
		"mapb": {name: "mapb", description: "Retrieve the previous page of locations", callback: commandMapb},
	}
}

var firstPage = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

func startRepl() {
	cfg := &config{}
	cfg.Next = &firstPage

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			words := cleanInput(text)

			if len(words) == 0 {
				continue
			}

			cmd, ok := getCommands()[words[0]]
			if ok {
				err := cmd.callback(cfg)
				if err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}
