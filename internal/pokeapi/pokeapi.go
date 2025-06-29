package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	LocationAreaAPI = "https://pokeapi.co/api/v2/location-area/"
)

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Count   int            `json:"count"`
	Next    string         `json:"next"`
	Prev    string         `json:"previous"`
	Results []LocationArea `json:"results"`
}

type Config struct {
	Next string
	Prev string
}

func GetLocations(c *Config, url string) ([]LocationArea, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request - %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting %v - %v", url, err)
	}
	defer res.Body.Close()

	var locations LocationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return nil, fmt.Errorf("error decoding %v - %v", url, err)
	}

	// Update the config
	c.Next = locations.Next
	c.Prev = locations.Prev

	return locations.Results, nil
}
