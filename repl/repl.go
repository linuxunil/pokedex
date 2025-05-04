package pokedexcli

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}
type Area struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type Config struct {
	Next string
	Prev string
}

var cmds = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Display 20 locations",
		callback:    commandMap,
	},
}

func commandMap(conf *Config) error {
	locations := Locations()
	return nil
}

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	for cmd := range cmds {
		fmt.Printf("Command: %s\nDescription: %s\n\n", cmds[cmd].name, cmds[cmd].description)
	}
	return nil
}

func cleanInput(text string) []string {
	fields := strings.Fields(strings.ToLower(text))
	return fields
}
