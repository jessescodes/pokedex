// file containing the cli command logic

package main

import (
	"fmt"
	"io"
	"os"
)

type CLICommand struct {
	name        string
	description string
	callback    func(w io.Writer) error
}

var CommandRegistery = map[string]CLICommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func commandExit(w io.Writer) error {
	fmt.Fprintf(w, "Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(w io.Writer) error {
	fmt.Fprintf(w, "Welcome to the Pokedex!\n"+
		"Usage:\n"+
		"help: Displays a help message\n"+
		"exit: Exits the Pokedex\n")
	return nil
}
