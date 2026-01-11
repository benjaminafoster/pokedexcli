package main

import (
	"fmt"
	"github.com/benjaminafoster/pokedexcli/pokeapi"
)

func commandMap(c Config) error {
	if c.nextLocationUrl == nil {
		
	}
	locationAreaData, err := pokepai.FetchLocationAreaData(locationAreaUrl)
	if err != nil {
		fmt.Println("Error occurred: ", err)
	}

	for _, location := range locationAreaData.Results {
		fmt.Println(location.Name)
	}

	return nil
}