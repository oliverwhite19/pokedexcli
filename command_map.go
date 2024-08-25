package main

import (
	"errors"
	"fmt"

	"github.com/oliverwhite19/pokedexcli/pokeapi"
)

func baseMap(link string, mem *memory) error {
	result, err := pokeapi.GetLocations(link)

	if err != nil {
		return err
	}

	mem.nextLocations = result.Next
	mem.previousLocations = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil

}

func commandMap(mem *memory) error {
	return baseMap(*mem.nextLocations, mem)
}

func commandMapB(mem *memory) error {
	if mem.previousLocations == nil {
		return errors.New("no previous locations available")
	}
	return baseMap(*mem.previousLocations, mem)
}
