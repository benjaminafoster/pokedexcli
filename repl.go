package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	var cleanedWords []string

	words := strings.Fields(text) // isolate words in string and ignore leading and trailing whitespaces
	for _, word := range words {
		cleanedWords = append(cleanedWords, strings.ToLower(word)) // make each word lowercase
	}

	return cleanedWords
}

func startREPL() {
	fmt.Println("Here she goes")
}