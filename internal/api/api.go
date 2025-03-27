package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

type LocationAreaResult struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func FetchLocations(page *string) (*LocationAreaResult, error) {
	url := baseUrl + "/location-area"
	if page != nil {
		url = *page
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

	result := &LocationAreaResult{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
