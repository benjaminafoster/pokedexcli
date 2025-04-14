package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location string) (LocationDeets, error) {
	url := locationURL + "/" + location

	if val, ok := c.cache.Get(url); ok {
		locationDeets := LocationDeets{}
		err := json.Unmarshal(val, &locationDeets)
		if err != nil {
			return LocationDeets{}, err
		}

		return locationDeets, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDeets{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDeets{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDeets{}, nil
	}

	locationDeets := LocationDeets{}
	err = json.Unmarshal(dat, &locationDeets)
	if err != nil {
		return LocationDeets{}, nil
	}

	c.cache.Add(url, dat)

	return locationDeets, nil
}


type LocationDeets struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}