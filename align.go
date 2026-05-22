package main

import (
	"fmt"
	"strings"
)

func AlignText(rows []string, input, align string, terminalWidth int) []string {
	result := []string{}

	contentWidth := len(rows[0])

	padding := terminalWidth - contentWidth

	if padding < 0 {
		padding = 0
	}

	switch align {

	case "right":
		for _, row := range rows {
			padRight := strings.Repeat(" ", padding) + row
			result = append(result, padRight)
		}

	case "center":
		for _, row := range rows {
			padLeft := strings.Repeat(" ", padding/2) + row
			result = append(result, padLeft)
		}

	case "left", "":
		return rows

	case "justify":
		fmt.Printf("DEBUG: contentWidth=%d termWidth=%d padding=%d noOfGaps=%d spacesPerGap=%d remainder=%d charWidth=%d\n",
			contentWidth, terminalWidth, padding, len(input)-1, padding/(len(input)-1), padding%(len(input)-1), contentWidth/len(input))

		noOfGaps := len(input) - 1

		if noOfGaps < 1 {
			for _, row := range rows {
				padLeft := strings.Repeat(" ", padding/2) + row
				result = append(result, padLeft)
			}
			break
		}

		spacesPerGap := padding / noOfGaps
		//remainder
		remainder := padding % noOfGaps

		charWidth := contentWidth / len(input)

		for _, row := range rows {
			var line strings.Builder
			for i := 0; i < len(input); i++ {
				chunk := row[i*charWidth : (i+1)*charWidth]
				line.WriteString(chunk)
				fmt.Printf("DEBUG chunk %d: %q\n", i, chunk)

				if i < len(input)-1 {
					spaces := spacesPerGap
					if i < remainder {
						spaces++
					}
					line.WriteString(strings.Repeat(" ", spaces))
				}

			}
			result = append(result, line.String())
		}

	}

	return result
}
