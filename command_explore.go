package main

import (
	"fmt"
)

//explore function
func commandExplore(cfg *config, args...string) error {
	// This command needs to get the second "parameter" in the line and run the GetPokemon command
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}
	location := args[0]
	locationDeets, err := cfg.pokeapiClient.ListPokemon(location)
	if err != nil {
		return err
	}

	pokemon := locationDeets.PokemonEncounters

	for _,val := range pokemon {
		fmt.Println(val.Pokemon.Name)
	}

	return nil
}


