package main

import (
	"errors"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
}

func executeCommand(command string) error {
	selectedCommand, ok := getCommands()[strings.TrimSpace(command)]
	if !ok {
		return errors.New("invalid command selected: try `help` to see available commands")
	}

	return selectedCommand.callback()
}
