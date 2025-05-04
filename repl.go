package pokedex

import (
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/linuxunil/pokedex/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}

type Config struct {
	Next string
	Prev any
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
	"mapb": {
		name:        "mapb",
		description: "Display previous 20 locations",
		callback:    commandMapb,
	},
}

func CommandMap(conf *Config) error {
	locations, err := pokeapi.Locations(conf.Next)
	if err != nil {
		return err
	}
	fmt.Printf("Next: %s", locations.Next)

	conf.Next = locations.Next
	conf.Prev = locations.Previous

	for location := range locations.Results {
		fmt.Println(locations.Results[location])
	}
	return nil
}

func CommandMapb(conf *Config) error {
	locations, err := pokeapi.Locations(conf.Next)
	if err != nil {
		return err
	}

	conf.Next = locations.Next
	conf.Prev = locations.Previous

	for location := range locations.Results {
		fmt.Println(locations.Results[location])
	}

	return nil
}

func CommandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	for cmd := range cmds {
		fmt.Printf("Command: %s\nDescription: %s\n\n", cmds[cmd].name, cmds[cmd].description)
	}
	return nil
}

func CleanInput(text string) []string {
	fields := strings.Fields(strings.ToLower(text))
	return fields
}
