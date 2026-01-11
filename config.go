package main

import pokeapi "github.com/benjaminafoster/pokedexcli/pokeapi"

type Config struct {
	pokeapiClient pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
}