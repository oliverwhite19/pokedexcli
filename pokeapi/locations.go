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

func GetLocations(link string) (locations, error) {

	var result locations
	decoder, err := getRequest(link)
	if err != nil {
		return locations{}, err
	}
	if err := decoder.Decode(&result); err != nil {
		return locations{}, err
	}

	return result, nil
}
