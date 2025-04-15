package main

import (
	"fmt"
	"math/rand/v2"
	"github.com/benjaminafoster/pokedexcli/internal/pokeapi"
)

var caughtPokemon = make(map[string]pokeapi.Pokemon)
var catchAttempts = make(map[string]int)

func commandCatch(cfg *config, args...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	// Define the target Pokemon name
	targetPokemon := args[0]

	// Check if you already have the Pokemon in your Pokedex
	if _, exists := caughtPokemon[targetPokemon]; exists {
		fmt.Printf("you have already caught %s\n", targetPokemon)
		return nil
	}

	// Get the Pokemon details from the PokeAPI
	pokemonDeets, err := cfg.pokeapiClient.PokemonDeets(targetPokemon)
	if err != nil {
		return err
	}

	// Logic for determining successful catch
	baseExp := pokemonDeets.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", targetPokemon)
	var catchSuccess bool
	if catchAttempts[targetPokemon] < 3 {
		catchSuccess = attemptCatch(baseExp)
	} else {
		fmt.Printf("you have attempted catching %s too many time...", targetPokemon)
		return nil
	}

	// Logic for adding the pokemon to the pokedex or counting catch failures
	if catchSuccess {
		caughtPokemon[targetPokemon] = pokemonDeets
		fmt.Printf("%s was caught!\n", targetPokemon)
		fmt.Printf("%s added to Pokedex\n", caughtPokemon[targetPokemon].Name)
	} else {
		fmt.Printf("%s escaped!\n", targetPokemon)

		_,exists := catchAttempts[targetPokemon]
		if exists {
			catchAttempts[targetPokemon]++
		} else {
			catchAttempts[targetPokemon] = 1
		}
		
	}
	return nil
}

func attemptCatch(baseExp int) bool {
	maxExp := 300.0 // This represents the hardest Pokemon to catch
	minCatchChance := 0.1 // Minimum 10% chance to catch
	maxCatchChance := 0.9 // Maximum 90% chance to catch

	// Normalize baseExp into a probability range
	normalized := float64(baseExp) / maxExp
	chance := maxCatchChance - (normalized * (maxCatchChance - minCatchChance))

	roll := rand.Float64()

	return roll < chance
}
