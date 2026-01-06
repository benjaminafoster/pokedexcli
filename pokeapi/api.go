package pokepai


import (
	"net/http"
	"io"
	"encoding/json"
	"log"
)


func fetchApiData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	
	return body
}


type LocationResponse struct {
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
}


type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func FetchLocationAreaData(url string) (LocationAreaResponse, error) {
	data := fetchApiData(url)

	locationResponse := LocationAreaResponse{}
	err := json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationResponse, nil

}