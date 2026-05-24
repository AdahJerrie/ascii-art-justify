package main

import "strings"

func align(rows []string, input, align string, termWidth int) []string {
	result := []string{}

	contentWidth := len(rows[0])
	padding := termWidth - contentWidth

	switch align {
	case "left", "":
		return rows

	case "right":
		for _, row := range rows {
			toRight := strings.Repeat(" ", padding) + row
			result = append(result, toRight)
		}

	case "center":
		for _, row := range rows {
			center := strings.Repeat(" ", padding/2) + row
			result = append(result, center)
		}

	case "justify":
		noOfGaps := len(input) - 1
		if noOfGaps < 1 {
			for _, row := range rows {
				center := strings.Repeat(" ", padding/2) + row
				result = append(result, center)
			}
			break
		}
		spacePerGap := padding / noOfGaps
		remainder := padding % noOfGaps //remainder
		charWidth := noOfGaps / len(input)

		for _, row := range rows {
			var line strings.Builder
			for i := 0; i < len(input); i++ {
				chunk := row[i*charWidth : (i+1)*charWidth]
				line.WriteString(chunk)

				if i < len(input)-1 {
					spaces := spacePerGap
					if remainder > 0 {
						spaces++
						remainder--
					}
				}
			}
			result = append(result, line.String())
		}
	}
	return result
}
