package main

import (
	"fmt"
	"errors"
)

func commandMapf(cfg *Config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.prevLocationUrl = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.prevLocationUrl = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}