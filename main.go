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
		text, _ := reader.ReadString('\n')

		err := executeCommand(text)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
