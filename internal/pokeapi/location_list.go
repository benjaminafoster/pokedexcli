package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//ListLocations
func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check first if the URL is currently in the cache. If not, complete the new request
	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, nil
	}

	locationsResp := RespLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocations{}, nil
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}