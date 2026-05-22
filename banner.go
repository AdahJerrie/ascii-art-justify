package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("error")
	}

	splited := strings.Split(string(data), "\n")
	asciiMap := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		char := rune(i)
		start := (i - 32) * 9
		if start+9 > len(splited) {
			return nil, errors.New("Invalid banner format")
		}
		completelines := splited[start+1 : start+9]
		asciiMap[char] = completelines
	}

	return asciiMap, nil
}

// package main

// import "strings"

// func StringToArt(input string) string {
// 	if input == "" {
// 		return ""
// 	}

// 	digits := map[rune][]string{
// 		'0': {
// 			" _ ",
// 			"| |",
// 			"| |",
// 			"|_|",
// 			"   ",
// 		},
// 		'1': {
// 			"   ",
// 			" | ",
// 			" | ",
// 			" | ",
// 			"   ",
// 		},
// 		'2': {
// 			" _ ",
// 			" _|",
// 			"|_ ",
// 			" _|",
// 			"   ",
// 		},
// 		'3': {
// 			" _ ",
// 			" _|",
// 			" _|",
// 			" _|",
// 			"   ",
// 		},
// 		'4': {
// 			"   ",
// 			"|_|",
// 			"  |",
// 			"  |",
// 			"   ",
// 		},
// 		'5': {
// 			" _ ",
// 			"|_ ",
// 			" _|",
// 			"|_|",
// 			"   ",
// 		},
// 		'6': {
// 			" _ ",
// 			"|_ ",
// 			"|_|",
// 			"|_|",
// 			"   ",
// 		},
// 		'7': {
// 			" _ ",
// 			"  |",
// 			"  |",
// 			"  |",
// 			"   ",
// 		},
// 		'8': {
// 			" _ ",
// 			"|_|",
// 			"|_|",
// 			"|_|",
// 			"   ",
// 		},
// 		'9': {
// 			" _ ",
// 			"|_|",
// 			" _|",
// 			" _|",
// 			"   ",
// 		},
// 	}

// 	// validate: only digits and \n allowed
// 	for _, ch := range input {
// 		if ch != '\n' && (ch < '0' || ch > '9') {
// 			return ""
// 		}
// 	}

// 	segments := strings.Split(input, "\n")
// 	var result strings.Builder

// 	for i, segment := range segments {
// 		if segment == "" {
// 			// empty segment from \n — write a blank line
// 			if i < len(segments)-1 {
// 				result.WriteString("\n")
// 			}
// 			continue
// 		}

// 		// build 5 rows for this segment
// 		for row := 0; row < 5; row++ {
// 			for _, ch := range segment {
// 				result.WriteString(digits[ch][row])
// 			}
// 			result.WriteString("\n")
// 		}
// 	}

// 	return result.String()
// }
