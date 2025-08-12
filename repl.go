package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Print(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
