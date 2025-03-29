package main

import (
	"fmt"
	"github.com/mattr/pokedex/internal/api"
)

func commandMap(cfg *config, args []string) error {
	result, err := api.GetLocationAreas(cfg.NextLocationURL, cfg.Cache)
	if err != nil {
		return err
	}

	// Update the config for the next fetch
	cfg.NextLocationURL = result.Next
	cfg.PreviousLocationURL = result.Previous

	// Display the current page of results
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.PreviousLocationURL == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	result, err := api.GetLocationAreas(cfg.PreviousLocationURL, cfg.Cache)
	if err != nil {
		return err
	}

	// Update the config for the next fetch
	cfg.NextLocationURL = result.Next
	cfg.PreviousLocationURL = result.Previous

	// Display the current page of results
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}
