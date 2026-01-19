package main

import pokeapi "github.com/benjaminafoster/pokedexcli/internal/pokeapi"

type Config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	caughtPokemon   map[string]pokeapi.Pokemon
}
