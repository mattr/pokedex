package main

import (
	"fmt"
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
	fmt.Println("Hello, World!")
}
