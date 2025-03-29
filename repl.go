package main

import (
	"bufio"
	"fmt"
	"github.com/mattr/pokedex/internal/api"
	"github.com/mattr/pokedex/internal/cache"
	"os"
	"strings"
)

type config struct {
	Cache               *cache.Cache
	Pokedex             map[string]api.Pokemon
	NextLocationURL     *string
	PreviousLocationURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	words := strings.Fields(lower)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch":   {name: "catch", description: "Attempt to catch a Pokemon", callback: CommandCatch},
		"exit":    {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"explore": {name: "explore", description: "Explore a location, accepts a name as argument", callback: commandExplore},
		"help":    {name: "help", description: "Show the help message", callback: commandHelp},
		"inspect": {name: "inspect", description: "Inspect a Pokemon in your Pokedex", callback: commandInspect},
		"map":     {name: "map", description: "Retrieve the next page of locations", callback: commandMap},
		"mapb":    {name: "mapb", description: "Retrieve the previous page of locations", callback: commandMapb},
		"pokedex": {name: "pokedex", description: "List the Pokemon in your Pokedex", callback: commandPokedex},
	}
}

func startRepl(cfg *config) {
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
				err := cmd.callback(cfg, words[1:])
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
