package main

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	locationResp, err := cfg.pokeapiClient.FetchLocationAreaData(cfg.nextLocationUrl)
	if err != nil {
		fmt.Println("Error occurred: ", err)
	}

	for _, location := range locationAreaData.Results {
		fmt.Println(location.Name)
	}

	return nil
}