package main

type cliCommand struct {
	name string
	description string
	callback func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays a list of locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous page of locations",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays a list of Pokemon in a specified location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a Pokemon",
			callback: commandCatch,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays a list of caught Pokemon",
			callback: commandPokedex,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a caught Pokemon",
			callback: commandInspect,
		},
	}
}