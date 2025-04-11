package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func startRepl() {
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
			err := command.callback()
			if err != nil {
				fmt.Println("Unknown command")
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
	callback            func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:           "exit",
			description:    "Exit the Pokedex",
			callback:       commandExit,
		},
		"help": {
			name:           "help",
			description:    "Displays a help message",
			callback:       commandHelp,
		},
	}
}