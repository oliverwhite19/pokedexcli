package main

import (
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func initializeMemory() memory {

	initialLocations := "https://pokeapi.co/api/v2/location-area/"

	return memory{
		previousLocations: nil,
		nextLocations:     &initialLocations,
	}
}
