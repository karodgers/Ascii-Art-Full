package pkg

import (
	"fmt"
	"strings"
)

func PrintAsciiArt(text string, p map[rune]string, color string, substring string) {
	if text == "" {
		return
	}

	invalidSubstrings := []string{"\\n", "\\t", "\\b", "\\v", "\\0", "\\f", "\\r"}
	for _, invalid := range invalidSubstrings {
		if strings.Contains(substring, invalid) {
			fmt.Printf("Error: %s special character in substring is not supported.\n", invalid)
			return
		}
	}
	specialCharatcters := []string{"\\t", "\\b", "\\v", "\\0", "\\f", "\\r"}
	for _, invalid := range specialCharatcters {
		if strings.Contains(text, invalid) {
			fmt.Printf("Error: %s special character in main string is not supported.\n", invalid)
			return
		}
	}

	text = strings.ReplaceAll(text, "\\n", "\n")

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			PrintArt(line, p, color, substring)
		}
	}
}
