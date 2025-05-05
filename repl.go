package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	pokeapi "github.com/linuxunil/pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config, opt ...[]string) error
}

type Config struct {
	Next string
	Prev string
	Url  string
}

var cmds = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    CommandExit,
	},
	"map": {
		name:        "map",
		description: "Display 20 locations",
		callback:    CommandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Display previous 20 locations",
		callback:    CommandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Show pokemon in the area",
		callback:    CommandExp,
	},
	"catch": {
		name:        "catch",
		description: "Attempt to catch a pokemon",
		callback:    CommandCatch,
	},
}
var pokedex = map[string]pokeapi.Pokemon{}

func CommandCatch(conf *Config, opt ...[]string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", opt[0][0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", opt[0][0])
	poke, err := pokeapi.Catch(url)
	if err != nil {
		return err
	}

	if rand.Int()%100 > poke.BaseExperience%100 {
		fmt.Printf("%v was caught!\n", poke.Name)
	} else {
		fmt.Printf("%v escaped\n", poke.Name)
	}
	return nil
}
func CommandExp(conf *Config, opt ...[]string) error {
	areas, err := pokeapi.Areas(conf.Url + opt[0][0])
	// fmt.Println(conf.Url + opt[0][0])
	if err != nil {
		return err
	}
	pokemon := areas.PokemonEncounters
	for p := range pokemon {
		fmt.Println(pokemon[p].Pokemon.Name)
	}
	return nil

}
func CommandMap(conf *Config, opt ...[]string) error {
	locations, err := pokeapi.Locations(conf.Next)
	if err != nil {
		return err
	}
	conf.Prev = conf.Next
	conf.Next = locations.Next

	for location := range locations.Results {
		fmt.Println(locations.Results[location].Name)
	}
	return nil
}

func CommandMapb(conf *Config, opt ...[]string) error {

	locations, err := pokeapi.Locations(conf.Prev)
	if err != nil {
		return err
	}

	conf.Next = locations.Next
	conf.Prev = locations.Previous
	for location := range locations.Results {
		fmt.Println(locations.Results[location].Name)
	}

	return nil
}

func CommandExit(conf *Config, opt ...[]string) error {
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
