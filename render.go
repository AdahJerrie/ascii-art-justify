package main

import (
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {
	words := SplitInput(input)
	output := []string{}
	var row strings.Builder
	for _, word := range words {
		for i := range 8 {
			for _, char := range word {
				row.WriteString(banner[char][i])
			}
			output = append(output, row.String())
			row.Reset()
		}
	}
	return output
}
