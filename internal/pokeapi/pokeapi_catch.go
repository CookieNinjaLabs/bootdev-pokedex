package pokeapi

import (
	"fmt"
	"math/rand"
)

const (
	catchSuccess = "was caught!"
	catchFailure = "escaped!"
)

var Pokedex = make(map[string]PokemonDetails)

func catchPokemonResultMessage(pokemon string, success bool) (string, error) {
	if pokemon == "" {
		return "", fmt.Errorf("pokemon is empty")
	}
	if success {
		return fmt.Sprintf("%s %s", pokemon, catchSuccess), nil
	}
	return fmt.Sprintf("%s %s", pokemon, catchFailure), nil
}

func CatchPokemon(pokemon string) error {
	// check if Pokemon exists
	if pokemon == "" {
		return fmt.Errorf("pokemon is empty")
	}

	pokemonDetails, err := GetPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("error getting pokemon details - %v", err)
	}

	// Throw a ball
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	//if c.Cache != nil {
	//	resBody, err := json.Marshal(pokemonDetails)
	//	if err != nil {
	//		return fmt.Errorf("error marshalling response - %v", err)
	//	}
	//	c.Cache.Add(url, resBody)
	//}

	// check resuslt based on base xp
	// TODO: pokedex and difficulty
	// - change the roll, so that the difficulty is based on the base xp of the Pokemon
	// - add the caught pokemon to the Pokedex
	roll := rand.Intn(pokemonDetails.BaseExperience)
	threshold := pokemonDetails.BaseExperience / 3
	result := false
	if roll > threshold {
		result = true
		Pokedex[pokemon] = pokemonDetails
	}

	// print msg
	resultMsg, err := catchPokemonResultMessage(pokemon, result)
	if err != nil {
		return err
	}
	fmt.Println(resultMsg)
	return nil
}
