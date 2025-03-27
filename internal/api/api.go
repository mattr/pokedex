package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

func FetchLocations(url string) (*LocationAreaResult, error) {
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

	locations := &LocationAreaResult{}
	err = json.Unmarshal(data, locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
