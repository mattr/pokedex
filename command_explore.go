package main

import (
	"errors"
	"fmt"
	"github.com/mattr/pokedex/internal/api"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no name given to explore")
	}

	name := args[0]

	result, err := api.GetLocationArea(&name, cfg.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range result.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
