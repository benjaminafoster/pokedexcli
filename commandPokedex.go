package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}
	
	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	
	return nil
}