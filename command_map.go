package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

type Response struct {
	Next           string                      `json:"next"`
	Previous       string                      `json:"previous"`
	Results        []map[string]string         `json:"results"`
}

func commandMap(c *Config) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"
	
	var choiceUrl string

	if c.Next != "" {
		choiceUrl = c.Next
	} else {
		choiceUrl = baseUrl
	}

	// submit an HTTP GET request and capture its response
	res, err := http.Get(choiceUrl)
	if err != nil {
		return fmt.Errorf("error submitting HTTP GET request: %w", err)
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()
	
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
	}
	
	// unmarshal the response into a Response struct
	responseStruct := Response{}
	if err := json.Unmarshal(body, &responseStruct); err != nil {
		return fmt.Errorf("error unmarshaling json data: %w", err)
	}

	// print the names of the location areas provided in the API response
	for _,result := range responseStruct.Results {
		fmt.Printf("%s\n", result["name"])
	}

	// update config struct
	c.Next = responseStruct.Next
	c.Previous = responseStruct.Previous

	return nil
}

func commandMapB(c *Config) error {
	
	var choiceUrl string

	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		choiceUrl = c.Previous
	}

	// submit an HTTP GET request and capture its response
	res, err := http.Get(choiceUrl)
	if err != nil {
		return fmt.Errorf("error submitting HTTP GET request: %w", err)
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()
	
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
	}
	
	// unmarshal the response into a Response struct
	responseStruct := Response{}
	if err := json.Unmarshal(body, &responseStruct); err != nil {
		return fmt.Errorf("error unmarshaling json data: %w", err)
	}

	// print the names of the location areas provided in the API response
	for _,result := range responseStruct.Results {
		fmt.Printf("%s\n", result["name"])
	}

	// update config struct
	c.Next = responseStruct.Next
	c.Previous = responseStruct.Previous

	return nil
}