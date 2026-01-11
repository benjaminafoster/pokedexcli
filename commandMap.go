package main

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		fmt.Println("Error occurred: ", err)
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.prevLocationUrl = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}