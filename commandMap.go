package main

import (
	"fmt"
	"errors"
	"github.com/benjaminafoster/pokedexcli/pokeapi"
)

func commandMap() error {
	locationDataUrl := "https://pokeapi.co/api/v2/location-area"
	locationData, err := pokepai.FetchLocationAreaData(locationDataUrl)
}