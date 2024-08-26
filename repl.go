package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	reader := bufio.NewScanner(os.Stdin)
	mem := initializeMemory()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		text := reader.Text()

		commands := cleanInput(text)

		selectedCommand, ok := getCommands()[commands[0]]
		if !ok {
			fmt.Println("invalid command selected: try `help` to see available commands")
			os.Exit(1)
		}
		err := selectedCommand.callback(&mem, commands[1:]...)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
