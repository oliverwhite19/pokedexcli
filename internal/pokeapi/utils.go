package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) getRequest(link string) (locations, error) {
	if val, ok := c.cache.Get(link); ok {
		locationsResp := locations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return locations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return locations{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return locations{}, err
	}
	if res.StatusCode > 299 {
		return locations{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return locations{}, err
	}

	locationsResp := locations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return locations{}, err
	}

	c.cache.Add(link, dat)
	return locationsResp, nil
}
