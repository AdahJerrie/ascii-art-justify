package main

import (
	"strings"
)

func BUildArt(input, align string, terminalWidth int, banner map[rune][]string) string {

	var b strings.Builder
	data := SplitInput(input)
	for i, char := range data {
		if char != "" {
			rendered := RenderLine(char, banner)
			alignAll := AlignText(rendered, char, align, terminalWidth)
			b.WriteString(strings.Join(alignAll, "\n"))
			b.WriteString("\n")
		} else if i < len(data)-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
