# Pokedex CLI

A command-line Pokedex application written in Go that allows you to explore Pokemon locations and encounters using
the [PokeAPI](https://pokeapi.co/).

## Features

- Interactive command-line interface
- Browse Pokemon locations with pagination
- Explore specific areas to find Pokemon
- Built-in caching system to reduce API calls
- Simple and intuitive commands

## Installation

### Prerequisites

- Go 1.16 or higher

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/cookieninjalabs/bootdev-pokedex.git
   ```

2. Navigate to the project directory:
   ```bash
   cd bootdev-pokedex
   ```

3. Build the application:
   ```bash
   go build
   ```

4. Run the application:
   ```bash
   ./bootdev-pokedex
   ```

## Usage

Once you start the application, you'll be presented with a prompt:

```
Pokedex >
```

You can enter commands at this prompt to interact with the Pokedex.

### Available Commands

- `help`: Displays a help message with all available commands
- `exit`: Exits the Pokedex application
- `map`: Displays the next 20 location areas
- `mapb`: Displays the previous 20 location areas
- `explore [area]`: Explores a specific location area and lists the Pokemon that can be found there

### Examples

#### Viewing locations

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
...
```

#### Exploring an area

```
Pokedex > explore eterna-city-area
Exploring eterna-city-area
Found Pokemon:
 - magikarp
 - gyarados
 - psyduck
 - golduck
 ...
```

## How It Works

The application uses the [PokeAPI](https://pokeapi.co/) to fetch data about Pokemon locations and encounters. It
implements a simple caching system to reduce the number of API calls and improve performance.

The cache automatically expires entries after 5 seconds to ensure that the data remains relatively fresh while still
providing performance benefits.

## Project Structure

- `main.go`: Entry point of the application, contains the REPL and command implementations
- `repl.go`: Contains utility functions for the REPL
- `internal/pokeapi/`: Package for interacting with the PokeAPI
    - `pokeapi_location_area.go`: Functions for fetching location areas
    - `pokeapi_explore.go`: Functions for exploring areas and finding Pokemon
- `internal/pokecache/`: Package implementing a simple caching system

## License

This project is licensed under the terms of the [license](LICENSE) included in the repository.

## Future ToDos

- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you
  are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances
  of catching pokemon
