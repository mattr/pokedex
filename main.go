package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(input string) []string {
	var list []string
	for _, token := range strings.Split(input, " ") {
		token = strings.TrimSpace(token)
		if token != "" {
			list = append(list, strings.ToLower(token))
		}
	}
	return list
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			input := cleanInput(text)
			fmt.Printf("Your command was: %s\n", input[0])
		}
	}
}
