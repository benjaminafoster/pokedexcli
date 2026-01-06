package pokepai


import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
	"log"
)


type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


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


func FetchLocationAreaData(url string) LocationAreaResponse {
	data := fetchApiData(url)

	locationResponse := LocationAreaResponse{}
	err := json.Unmarshal(data, &locationResponse)
	if err != nil {
		fmt.Println(err)
	}

	return locationResponse

}