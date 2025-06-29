package main

import (
	"strings"
)

func cleanInput(input string) []string {
	if input == "" {
		return []string{}
	}
	trimmedInput := strings.TrimSpace(input)
	lowercaseInput := strings.ToLower(trimmedInput)
	result := strings.Split(lowercaseInput, " ")
	return result
}
