package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func cleanInput(text string) []string {
	var cleanedWords []string

	words := strings.Fields(text) // isolate words in string and ignore leading and trailing whitespaces
	for _, word := range words {
		cleanedWords = append(cleanedWords, strings.ToLower(word)) // make each word lowercase
	}

	return cleanedWords
}

func startREPL(c *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		cleanCommand := cleanedInput[0]
		commandArgs := cleanedInput[1:]
		cmds := getCommands()
		cmd, ok := cmds[cleanCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(c, commandArgs...)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
}