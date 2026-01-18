package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("explore command requires a location argument")
	}
	
	location := args[0]
	
	locationResp, err := cfg.pokeapiClient.GetLocation(location)
	if err != nil {
		return err
	}
	
	encounters := locationResp.PokemonEncounters
	
	for _, encounter := range encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	
	return nil
}
