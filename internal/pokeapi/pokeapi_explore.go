package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPokemonInArea(c *Config, area string) ([]PokemonEncounters, error) {
	url := fmt.Sprintf("%s%s", LocationAreaAPI, area)
	if c.Cache != nil {
		if cachedData, ok := c.Cache.Get(url); ok {
			var explorationResponse ExplorationResponse
			if err := json.Unmarshal(cachedData, &explorationResponse); err != nil {
				return nil, fmt.Errorf("error unmarshalling cached exploration response: %v", err)
			}

			fmt.Println("Using cached exploration response")
			return explorationResponse.PokemonEncounters, nil
		}
	}

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

	var explorationResponse ExplorationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&explorationResponse); err != nil {
		return nil, fmt.Errorf("error decoding %v - %v", url, err)
	}

	if c.Cache != nil {
		resBody, err := json.Marshal(explorationResponse)
		if err != nil {
			return nil, fmt.Errorf("error marshalling response body - %v", err)
		}
		c.Cache.Add(url, resBody)
	}

	return explorationResponse.PokemonEncounters, nil
}
