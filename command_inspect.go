package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("please specify a pokemon to inspect")
	}

	name := args[0]
	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Forms:")
	for _, form := range pokemon.Forms {
		fmt.Printf("- %s\n", form.Name)
	}

	fmt.Println("Types:")
	for _, tt := range pokemon.Types {
		fmt.Printf("- %s\n", tt.Type.Name)
	}

	return nil
}
