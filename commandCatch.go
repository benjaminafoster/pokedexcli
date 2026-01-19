package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("the catch command requires a pokemon name argument")
	}
	
	pokemonName := args[0]
	
	pokemonResp, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	result := rand.Intn(pokemonResp.BaseExperience)
	if result > 40 {
		fmt.Printf("%s excaped!\n", pokemonName)
		return nil
	}
	
	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.caughtPokemon[pokemonName] = pokemonResp
	
	return nil
}