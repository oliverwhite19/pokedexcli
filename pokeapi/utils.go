package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getRequest(link string) (*json.Decoder, error) {

	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	decoder := json.NewDecoder(res.Body)

	return decoder, nil
}
