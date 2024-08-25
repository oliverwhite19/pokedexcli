package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		reader.Scan()

		text := reader.Text()

		commandName := cleanInput(text)[0]

		err := executeCommand(commandName)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
