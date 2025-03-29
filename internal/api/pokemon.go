package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mattr/pokedex/internal/cache"
	"io"
	"net/http"
)

type PokemonResult struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func GetPokemon(name *string, cache *cache.Cache) (*PokemonResult, error) {
	if name == nil {
		return nil, errors.New("pokemon name is required")
	}

	url := fmt.Sprintf("%s/pokemon/%s/", baseUrl, *name)

	if result, ok := cache.Get(url); ok {
		return parsePokemonResult(result)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cache.Add(url, data)
	return parsePokemonResult(data)
}

func parsePokemonResult(data []byte) (*PokemonResult, error) {
	result := &PokemonResult{}
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
