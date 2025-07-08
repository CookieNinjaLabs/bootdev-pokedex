package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	LocationAreaAPI = "https://pokeapi.co/api/v2/location-area/"
)

type Config struct {
	Next  string
	Prev  string
	Cache CacheInterface
}

type CacheInterface interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
}

func GetLocations(c *Config, url string) ([]LocationArea, error) {
	// Check cache first
	if c.Cache != nil {
		if cachedData, ok := c.Cache.Get(url); ok {
			var locations LocationResponse
			if err := json.Unmarshal(cachedData, &locations); err != nil {
				return nil, fmt.Errorf("error unmarshalling cached locations: %v", err)
			}
			// Update config with cached values
			c.Next = locations.Next
			c.Prev = locations.Prev

			fmt.Println("Using cached locations")
			return locations.Results, nil
		}
	}

	// Make API requests if not in cache
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

	// Add to cache
	if c.Cache != nil {
		resBody, err := json.Marshal(locations)
		if err != nil {
			return nil, fmt.Errorf("error marshalling response body - %v", err)
		}
		c.Cache.Add(url, resBody)
	}

	// Update the config
	c.Next = locations.Next
	c.Prev = locations.Prev

	return locations.Results, nil
}
