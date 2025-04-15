package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args...string) error {
	if len(args) > 1 {
		return fmt.Errorf("no additional arguments accepted for pokedex command")
	}
	fmt.Println("Your Pokedex:")
	for k, _ := range caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}