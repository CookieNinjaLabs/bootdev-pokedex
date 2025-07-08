package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/cookieninjalabs/bootdev-pokedex/internal/pokeapi"
	"github.com/cookieninjalabs/bootdev-pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config, args []string) error
}

var cmd map[string]cliCommand

func main() {
	cmd = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location and get a list of pokemon in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspectPokedex,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all Pokemon you caught in your Pokedex",
			callback:    commandPokedex,
		},
	}
	cache := pokecache.NewCache(5 * time.Second)

	cfg := &pokeapi.Config{
		Cache: cache,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command, ok := cmd[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			var args []string
			if len(cleanedInput[1:]) >= 0 {
				args = cleanedInput[1:]
			}

			err := command.callback(cfg, args)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}
	}
}

func commandExit(c *pokeapi.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *pokeapi.Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for k, v := range cmd {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(c *pokeapi.Config, args []string) error {
	url := c.Next
	if url == "" {
		url = pokeapi.LocationAreaAPI
	}

	locations, err := pokeapi.GetLocations(c, url)
	if err != nil {
		return fmt.Errorf("error getting locations: %v", err)
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *pokeapi.Config, args []string) error {
	url := c.Prev
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := pokeapi.GetLocations(c, url)
	if err != nil {
		return fmt.Errorf("error getting locations: %v", err)
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(c *pokeapi.Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <area>")
	}

	area := args[0]
	pokemonEncounters, err := pokeapi.GetPokemonInArea(c, area)
	if err != nil {
		return fmt.Errorf("error getting pokemon in area: %v", err)
	}

	fmt.Printf("Exploring %s\n", area)
	fmt.Println("Found Pokemon:")
	for _, encounter := range pokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *pokeapi.Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	pokemon := cleanInput(args[0])[0]
	err := pokeapi.CatchPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("error catching pokemon: %v", err)
	}

	return nil
}

func commandInspectPokedex(c *pokeapi.Config, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	pokemon := cleanInput(args[0])[0]

	pokedexEntry, ok := pokeapi.Pokedex[pokemon]
	if !ok {
		return fmt.Errorf("pokemon not found in pokedex")
	}

	var pokemonTypes string
	for _, pokemonType := range pokedexEntry.Types {
		pokemonTypes += fmt.Sprintf("  - %s\n", pokemonType.Type.Name)
	}

	fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
  -hp: %d
  -attack: %d
  -defense: %d
  -special-attack: %d
  -special-defense: %d
  -speed: %d
Types:
%s`,
		pokedexEntry.Name,
		pokedexEntry.Height,
		pokedexEntry.Weight,
		pokedexEntry.Stats[0].BaseStat,
		pokedexEntry.Stats[1].BaseStat,
		pokedexEntry.Stats[2].BaseStat,
		pokedexEntry.Stats[3].BaseStat,
		pokedexEntry.Stats[4].BaseStat,
		pokedexEntry.Stats[5].BaseStat,
		pokemonTypes)

	return nil
}

func commandPokedex(c *pokeapi.Config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("usage: pokedex")
	}
	if len(pokeapi.Pokedex) == 0 {
		fmt.Println("pokedex is empty")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokeapi.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
