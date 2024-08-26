package main

import "fmt"

func commandPokedex(mem *memory, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range mem.capturedPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
