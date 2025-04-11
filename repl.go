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

		command := cleanInputText[0]

		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	words := strings.Fields(lowerCaseText)
	return words
}