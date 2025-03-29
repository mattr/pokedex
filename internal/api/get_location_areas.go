package api

import (
	"encoding/json"
	"errors"
	"github.com/mattr/pokedex/internal/cache"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(page *string, cache *cache.Cache) (*LocationAreas, error) {
	url := baseUrl + "/location-area"
	if page != nil {
		url = *page
	}

	if resp, ok := cache.Get(url); ok {
		return parseLocationAreasResult(resp)
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
	return parseLocationAreasResult(data)
}

func parseLocationAreasResult(data []byte) (*LocationAreas, error) {
	result := &LocationAreas{}
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
