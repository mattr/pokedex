package main

import (
	"errors"
	"fmt"
	"github.com/mattr/pokedex/internal/api"
	"math/rand"
)

func CommandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("pokemon name is required")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	result, err := api.GetPokemon(&name, cfg.Cache)
	if err != nil {
		return err
	}

	exp := result.BaseExperience
	rnd := rand.Intn(1000)
	if rnd > exp {
		fmt.Printf("%s was caught!\n", name)
		cfg.Pokedex[name] = *result
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
