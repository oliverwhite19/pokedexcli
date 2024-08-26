package main

import (
	"errors"
	"fmt"
)

func commandExplore(mem *memory, args ...string) error {

	if len(args) != 1 {
		return errors.New("expore requires a location name as a parameter")
	}
	locationName := args[0]

	queryURL := "https://pokeapi.co/api/v2/location-area/" + locationName

	locationDetails, err := mem.pokeClient.GetLocation(queryURL)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
