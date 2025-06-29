package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cookieninjalabs/bootdev-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
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
	}

	cfg := &pokeapi.Config{}

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
			err := command.callback(cfg)
			if err != nil {
				fmt.Println("Error: %s\n", err)
			}
		}
	}
}

func commandExit(c *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for k, v := range cmd {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(c *pokeapi.Config) error {
	url := c.Next
	if url == "" {
		url = pokeapi.LocationAreaAPI
	}

	locations, err := pokeapi.GetLocations(c, url)
	if err != nil {
		return fmt.Errorf("Error getting locations: %v", err)
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *pokeapi.Config) error {
	url := c.Prev
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := pokeapi.GetLocations(c, url)
	if err != nil {
		return fmt.Errorf("Error getting locations: %v", err)
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
