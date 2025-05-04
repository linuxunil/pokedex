package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		lineTokens := cleanInput(line)
		switch lineTokens[0] {
		case "help":
		case "exit":
		default:
			fmt.Printf("Your command was: %v\n", lineTokens[0])
		}
	}
}
