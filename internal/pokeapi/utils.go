package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) getRequest(link string) ([]byte, error) {
	if val, ok := c.cache.Get(link); ok {
		return val, nil
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(link, dat)

	return dat, nil
}
