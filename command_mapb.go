package main

import (
	"fmt"
	"github.com/mattr/pokedex/internal/api"
)

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	result, err := api.FetchLocations(*cfg.Previous)
	if err != nil {
		return err
	}

	// Update the config for the next fetch
	cfg.Next = result.Next
	cfg.Previous = result.Previous

	// Display the current page of results
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}
