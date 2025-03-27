package main

import (
	"fmt"
	"github.com/mattr/pokedex/internal/api"
)

func commandMap(cfg *config) error {
	if cfg.Next == nil {
		fmt.Println("You are on the last page")
		return nil
	}

	result, err := api.FetchLocations(*cfg.Next)
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
