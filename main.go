package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Printf("help: Displays a help message\n")
	fmt.Printf("exit: Exit the Pokedex\n")
	fmt.Println("")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func getcommandMap() map[string]cliCommand {
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

func main() {
	commandMap := getcommandMap()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")

		scanner.Scan()
		input := scanner.Text()

		if cmd, exists := commandMap[input]; exists {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error", err)
			}
		} else {
			fmt.Println("Invalid input, only accepts help or exit")
		}
	}
}
