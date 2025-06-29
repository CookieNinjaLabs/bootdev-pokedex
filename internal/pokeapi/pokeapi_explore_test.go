package pokeapi

import (
	"testing"
)

func TestGetPokemonInArea(t *testing.T) {
	cases := []struct {
		area        string
		id          int
		name        string
		url         string
		description string
	}{
		{
			area: "canalave-city-area",
			id:   0,
			name: "tentacool",
			url:  "https://pokeapi.co/api/v2/pokemon/72/",
		},
		{
			area: "canalave-city-area",
			id:   1,
			name: "tentacruel",
			url:  "https://pokeapi.co/api/v2/pokemon/73/",
		},
		{
			area: "eterna-city-area",
			id:   0,
			name: "psyduck",
			url:  "https://pokeapi.co/api/v2/pokemon/54/",
		},
		{
			area: "eterna-city-area",
			id:   1,
			name: "golduck",
			url:  "https://pokeapi.co/api/v2/pokemon/55/",
		},
	}

	t.Run("Single Pokemon Names in an Area", func(t *testing.T) {
		config := &Config{}
		for _, c := range cases {
			pokemonEncounters, err := GetPokemonInArea(config, c.area)
			if err != nil {
				t.Fatalf("Expected no error in area %s, got %v", c.area, err)
			}

			pokemon := pokemonEncounters[c.id].Pokemon

			if pokemon.Name != c.name {
				t.Errorf("Expected %s, got %v", c.name, pokemon.Name)
			}
		}
	})

	t.Run("Multiple Pokemon Names in an Area", func(t *testing.T) {
		config := &Config{}
		for _, c := range cases {
			pokemonEncounters, err := GetPokemonInArea(config, c.area)
			if err != nil {
				t.Fatalf("Expected no error in area %s, got %v", c.area, err)
			}

			if len(pokemonEncounters) <= 1 {
				t.Errorf("Expected more than 1 pokemon in area %s, got %d", c.area, len(pokemonEncounters))
			}
		}
	})

}
