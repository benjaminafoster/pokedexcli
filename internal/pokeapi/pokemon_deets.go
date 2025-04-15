package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) PokemonDeets(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonDeets := Pokemon{}
		err := json.Unmarshal(val, &pokemonDeets)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshaling json data from cache: %w", err)
		}

		return pokemonDeets, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error creating http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error completing http request: %w", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response data: %w", err)
	}

	pokemonDeets := Pokemon{}
	err = json.Unmarshal(dat, &pokemonDeets)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling json: %w", err)
	}

	c.cache.Add(url, dat)


	return pokemonDeets, nil
}