package main

import (
	"github.com/oliverwhite19/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*memory, ...string) error
}

type memory struct {
	previousLocations *string
	nextLocations     *string
	pokeClient        pokeapi.Client
	capturedPokemon   map[string]pokeapi.ResponsePokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explores a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
