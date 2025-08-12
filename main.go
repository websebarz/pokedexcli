package main

import (
	"time"
)

func main() {
	pokeclient := pokeapi.NewClient(5, time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
