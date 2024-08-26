package main

import (
	"errors"
	"fmt"
)

func commandInspect(mem *memory, args ...string) error {

	if len(args) != 1 {
		return errors.New("inspect requires a pokemon name as a parameter")
	}
	pokemonName := args[0]

	pokemon, ok := mem.capturedPokemon[pokemonName]
	if !ok {
		fmt.Printf("You have yet to capture %s!\n", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" -%v\n", t.Type.Name)
	}
	return nil
}
