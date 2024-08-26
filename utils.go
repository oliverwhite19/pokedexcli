package main

import (
	"strings"
	"time"

	"github.com/oliverwhite19/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func initializeMemory() memory {

	initialLocations := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	pokeDexClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	return memory{
		previousLocations: nil,
		nextLocations:     &initialLocations,
		pokeClient:        pokeDexClient,
		capturedPokemon:   make(map[string]pokeapi.ResponsePokemon),
	}
}
