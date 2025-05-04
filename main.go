package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/linuxunil/pokedexcli"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	conf := Config{}
	for {
		fmt.Print("Pokedex > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		lineTokens := cleanInput(line)
		cmd := lineTokens[0]
		// args := lineTokens[1:]

		switch cmd {
		case "map":
			cmds[cmd].callback(&conf)
		case "exit":
			cmds[cmd].callback(&conf)
		default:
			commandHelp()
		}
	}
}
