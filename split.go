package main

import (
	"errors"
	"strings"
)

func SplitInput(text string) []string {
	return strings.Split(text, "\\n")
}

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return char, errors.New("invalid ascii char")
		}
	}
	return 0, nil
}
