package main

import (
	"fmt"
)

func commandInspect(cfg *config, args...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide the name of a pokemon")
	}

	targetPokemon := args[0]
	if pokemon, exists := caughtPokemon[targetPokemon]; exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pType := range pokemon.Types {
			fmt.Printf("  - %s\n", pType.Type.Name)
		}
		return nil
	} else {
		fmt.Printf("you have not caught %s\n", targetPokemon)
	}

	return nil
}