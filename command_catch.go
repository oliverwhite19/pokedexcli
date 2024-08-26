package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(mem *memory, args ...string) error {

	if len(args) != 1 {
		return errors.New("expore requires a pokemon name as a parameter")
	}
	pokemonName := args[0]

	queryURL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonDetails, err := mem.pokeClient.GetPokemon(queryURL)

	if err != nil {
		return err
	}

	randomSeed := rand.Intn(pokemonDetails.BaseExperience)

	if randomSeed < 40 {
		fmt.Printf("%s was caught!\n", pokemonDetails.Name)
		mem.capturedPokemon[pokemonDetails.Name] = pokemonDetails
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)

	}
	return nil

}
