package main

import (
	"time"

	"github.com/benjaminafoster/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 10 * time.Minute)
	config := &Config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startREPL(config)
}
