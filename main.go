package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Hello, World!\n")
}

func cleanInput(text string) []string {
	word := strings.Fields(text)
	return word
}
