package handlers

import (
	"fmt"
	"strings"
)

func PrintAsciiArt(text string, asciiArtMap map[rune]string, align string, width int) {
	if text == "" {
		return
	}

	specialCharacters := []string{"\\t", "\\b", "\\v", "\\0", "\\f", "\\r"}
	for _, invalid := range specialCharacters {
		if strings.Contains(text, invalid) {
			fmt.Printf("Error: unsupported %s special character detected.\n", invalid)
			return
		}
	}

	if strings.Contains(text, "\\n") {
		input := strings.Split(text, "\\n")
		count := 0
		for _, word := range input {
			if word == "" {
				count++
				if count < len(input) {
					fmt.Println()
				}
			} else if len(word) > 0 {
				PrintLineByLine(word, asciiArtMap, align, width)
			}
		}
	} else {
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			if len(line) > 0 {
				PrintLineByLine(line, asciiArtMap, align, width)
			} else {
				fmt.Println()
			}
		}
	}
}
