package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	url := "https://pokeapi.co/api/v2/location-area/"
	conf := Config{url, "", url}
	for {
		fmt.Print("Pokedex > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		lineTokens := CleanInput(line)
		args := lineTokens[1:]

		if cmd, ok := cmds[lineTokens[0]]; ok {
			cmd.callback(&conf, args)
		} else {
			CommandHelp()
		}
	}
}
