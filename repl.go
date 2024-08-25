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

		commandName := cleanInput(text)[0]

		selectedCommand, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("invalid command selected: try `help` to see available commands")
			os.Exit(1)
		}
		err := selectedCommand.callback(&mem)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
