package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	url := "https://pokeapi.co/api/v2/location-area/"
	conf := Config{url, ""}
	for {
		fmt.Print("Pokedex > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		lineTokens := CleanInput(line)
		cmd := lineTokens[0]
		// args := lineTokens[1:]

		switch cmd {
		case "map":
			cmds[cmd].callback(&conf)
		case "mapb":
			cmds[cmd].callback(&conf)
		case "exit":
			cmds[cmd].callback(&conf)
		default:
			CommandHelp()
		}
	}
}
