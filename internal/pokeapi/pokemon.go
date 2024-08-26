package pokeapi

import "encoding/json"

type ResponsePokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func (c *Client) GetPokemon(link string) (ResponsePokemon, error) {
	result, err := c.getRequest(link)
	if err != nil {
		return ResponsePokemon{}, err
	}

	pokemonResponse := ResponsePokemon{}
	err = json.Unmarshal(result, &pokemonResponse)
	if err != nil {
		return ResponsePokemon{}, err
	}

	return pokemonResponse, nil
}
