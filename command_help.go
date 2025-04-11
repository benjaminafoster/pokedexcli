package main

import (
	"fmt"
)

func commandHelp(c *Config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _,cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}