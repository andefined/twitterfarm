package main

import (
	"log"

	"github.com/andefined/twitterfarm/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
