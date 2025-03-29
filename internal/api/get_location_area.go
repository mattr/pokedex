package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mattr/pokedex/internal/cache"
	"io"
	"net/http"
)

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocationArea(name *string, cache *cache.Cache) (*LocationArea, error) {
	if name == nil {
		return nil, errors.New("location area name is required")
	}
	url := fmt.Sprintf("%s/location-area/%s/", baseUrl, *name)

	if resp, ok := cache.Get(url); ok {
		return parseLocationAreaResult(resp)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cache.Add(url, data)
	return parseLocationAreaResult(data)
}

func parseLocationAreaResult(data []byte) (*LocationArea, error) {
	result := &LocationArea{}
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
