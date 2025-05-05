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
	"inspect": {
		name:        "inspect",
		description: "See pokemon stats",
		callback:    CommandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "See Pokedex",
		callback:    CommandPokedex,
	},
}
var pokedex = map[string]pokeapi.Pokemon{}

func CommandPokedex(conf *Config, opt ...[]string) error {
	fmt.Println("Your Pokedex")
	for p := range pokedex {
		fmt.Printf(" - %v\n", pokedex[p].Name)
	}
	return nil
}
func CommandInspect(conf *Config, opt ...[]string) error {
	type statKey int
	const (
		hp statKey = iota
		atk
		def
		satk
		sdef
		spd
	)
	/*Name: pidgey
	  Height: 3
	  Weight: 18
	  Stats:
	    -hp: 40
	    -attack: 45
	    -defense: 40
	    -special-attack: 35
	    -special-defense: 35
	    -speed: 56
	  Types:
	    - normal
	    - flying
	*/

	if poke, ok := pokedex[strings.ToLower(opt[0][0])]; ok {
		fmt.Printf("Name: %v\n", poke.Name)
		fmt.Printf("Height: %v\n", poke.Height)
		fmt.Printf("Weight: %v\n", poke.Weight)
		fmt.Printf("Stats:\n")
		fmt.Printf("-hp: %v\n", poke.Stats[hp].BaseStat)
		fmt.Printf("-attack: %v\n", poke.Stats[atk].BaseStat)
		fmt.Printf("-defense: %v\n", poke.Stats[def].BaseStat)
		fmt.Printf("-special-attack: %v\n", poke.Stats[satk].BaseStat)
		fmt.Printf("-special-defense: %v\n", poke.Stats[sdef].BaseStat)
		fmt.Printf("-speed: %v\n", poke.Stats[spd].BaseStat)
		fmt.Printf("Types:\n")
		for t := range poke.Types {
			fmt.Printf("- %v\n", poke.Types[t].Type.Name)
		}
	} else {
		fmt.Printf("You haven't caught a %v\n", opt[0][0])
	}

	return nil

}
func CommandCatch(conf *Config, opt ...[]string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", opt[0][0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", opt[0][0])
	poke, err := pokeapi.Catch(url)
	if err != nil {
		return err
	}

	if _, ok := pokedex[strings.ToLower(poke.Name)]; !ok {
		if rand.Int()%100 > poke.BaseExperience%100 {
			fmt.Printf("%v was caught!\n", poke.Name)
			pokedex[strings.ToLower(poke.Name)] = poke
		} else {
			fmt.Printf("%v escaped\n", poke.Name)
		}
	} else {
		fmt.Printf("You already have a %v", poke.Name)
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
