package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationAreaResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	
	if val, ok := c.cache.Get(url); ok {
		locationResponse := LocationAreaResponse{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	
	c.cache.Add(url, data)

	return locationResponse, nil

}

func (c *Client) GetLocation(location string) (LocationResponse, error) {
	url := baseUrl + "/location-area/" + location
	
	if val, ok := c.cache.Get(url); ok {
		locationResponse := LocationResponse{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return LocationResponse{}, err
		}
		return locationResponse, nil
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}
	
	locationInfo := LocationResponse{}
	err = json.Unmarshal(data, &locationInfo)
	if err != nil {
		return LocationResponse{}, err
	}
	
	c.cache.Add(url, data)
	
	return locationInfo, nil
}

func (c *Client) GetPokemonInfo(pokemon string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemon
	
	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	
	pokemonInfo := Pokemon{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}
	
	c.cache.Add(url, data)
	
	return pokemonInfo, nil
}
