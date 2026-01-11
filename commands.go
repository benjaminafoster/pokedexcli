package main

type cliCommand struct {
	name string
	description string
	callback func(*Config) error
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
	}
}