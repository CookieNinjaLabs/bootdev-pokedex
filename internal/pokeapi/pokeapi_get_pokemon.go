package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPokemon(pokemon string) (PokemonDetails, error) {
	if pokemon == "" {
		return PokemonDetails{}, fmt.Errorf("pokemon is empty")
	}

	client := &http.Client{}

	url := PokemonAPI + pokemon
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error creating request - %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error requesting %v - %v", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PokemonDetails{}, fmt.Errorf("error requesting %v - %v", url, res.Status)
	}

	var pokemonDetails PokemonDetails
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemonDetails); err != nil {
		return PokemonDetails{}, fmt.Errorf("error decoding %v - %v", url, err)
	}

	return pokemonDetails, nil
}
