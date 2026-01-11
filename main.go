package main


func main() {
	config := Config{
		nextLocationUrl: "https://pokeapi.co/api/v2/location-area",
		prevLocationUrl: nil,
	}

	startREPL(config)
}