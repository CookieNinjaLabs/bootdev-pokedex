package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExplorationResponse struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}
type VersionEncounterDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}
type PokemonEncounters struct {
	Pokemon        Pokemon                   `json:"pokemon"`
	VersionDetails []VersionEncounterDetails `json:"version_details"`
}

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
