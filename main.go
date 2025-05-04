package main

import (
	"bufio"
	"fmt"
	"os"

)

func main() {
	reader := bufio.NewReader(os.Stdin)
	conf := pokedexcli.Config{}
	// conf := Config{"", ""}
	for {
		fmt.Print("Pokedex > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		lineTokens := pokedexcli.CleanInput(line)
		cmd := lineTokens[0]
		// args := lineTokens[1:]

		switch cmd {
		case "map":
			pokedexcli.cmds[cmd].callback(&conf)
		case "mapb":
			pokedexcli.cmds[cmd].callback(&conf)
		case "exit":
			pokedexcli.cmds[cmd].callback(&conf)
		default:
			pokedexcli.commandHelp()
		}
	}
}
