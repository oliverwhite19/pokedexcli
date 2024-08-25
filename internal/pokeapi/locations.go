package pokeapi

type locations struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

func (c *Client) GetLocations(link string) (locations, error) {

	var result locations

	result, err := c.getRequest(link)
	if err != nil {
		return locations{}, err
	}

	return result, nil
}
