package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/benjaminafoster/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient      pokeapi.Client
	nextLocationsURL   *string
	prevLocationsURL   *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleanInputText := cleanInput((scanner.Text()))
		if len(cleanInputText) == 0 {
			continue
		}

		// get the first word from the cleaned input which is the command
		commandName := cleanInputText[0]

		commandRegistry := getCommands()

		command, exists := commandRegistry[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

		
	}
}

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	words := strings.Fields(lowerCaseText)
	return words
}

type cliCommand struct {
	name                string
	description         string
	callback            func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:           "map",
			description:    "Displays the next set of 20 Pokedex location areas",
			callback:       commandMap,
		},
		"mapb": {
			name:           "mapb",
			description:    "Displays the previous set of 20 Pokedex location areas",
			callback:       commandMapB,
		},
		"help": {
			name:           "help",
			description:    "Displays a help message",
			callback:       commandHelp,
		},
		"exit": {
			name:           "exit",
			description:    "Exit the Pokedex",
			callback:       commandExit,
		},
	}
}