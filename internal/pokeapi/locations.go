package pokeapi

import (
	"encoding/json"
)

type ResponseLocations struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

type ResponseLocation struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocations(link string) (ResponseLocations, error) {
	result, err := c.getRequest(link)
	if err != nil {
		return ResponseLocations{}, err
	}

	locationsResp := ResponseLocations{}
	err = json.Unmarshal(result, &locationsResp)
	if err != nil {
		return ResponseLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetLocation(link string) (ResponseLocation, error) {

	result, err := c.getRequest(link)
	if err != nil {
		return ResponseLocation{}, err
	}

	locationsResp := ResponseLocation{}
	err = json.Unmarshal(result, &locationsResp)

	if err != nil {
		return ResponseLocation{}, err

	}

	return locationsResp, nil

}
