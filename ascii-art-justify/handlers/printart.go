package handlers

import (
	"fmt"
	"strings"
)

const asciiHeight = 8

func PrintLineByLine(text string, asciiArtMap map[rune]string, align string, width int) {
	var asciiArtLines [asciiHeight]string
	wordSpaces := []int{}
	totalWidth := 0

	for _, char := range text {
		art, ok := asciiArtMap[char]
		if !ok {
			fmt.Println("char not found")
			return
		}
		lines := strings.Split(art, "\n")
		for j := 0; j < asciiHeight; j++ {
			if j < len(lines) {
				asciiArtLines[j] += lines[j]
			} else {
				asciiArtLines[j] += strings.Repeat(" ", len(lines[0]))
			}
		}
		if char == ' ' {
			wordSpaces = append(wordSpaces, totalWidth)
		}
		totalWidth += len(lines[0])
	}

	// Apply alignment
	switch align {
	case "center":
		padding := (width - totalWidth) / 2
		for i := range asciiArtLines {
			asciiArtLines[i] = strings.Repeat(" ", padding) + asciiArtLines[i]
		}
	case "right":
		padding := width - totalWidth
		for i := range asciiArtLines {
			asciiArtLines[i] = strings.Repeat(" ", padding) + asciiArtLines[i]
		}
	case "left":
		// No additional padding needed for left alignment
	case "justify":
		if totalWidth < width && len(wordSpaces) > 0 {
			extraSpaces := width - totalWidth
			spacesToAdd := extraSpaces / len(wordSpaces)
			remainder := extraSpaces % len(wordSpaces)

			for i := range asciiArtLines {
				newLine := []rune(asciiArtLines[i])
				offset := 0
				for j, pos := range wordSpaces {
					additionalSpaces := spacesToAdd
					if j < remainder {
						additionalSpaces++
					}
					adjustedPos := pos + offset
					newLine = append(newLine[:adjustedPos], append([]rune(strings.Repeat(" ", additionalSpaces)), newLine[adjustedPos:]...)...)
					offset += additionalSpaces
				}
				asciiArtLines[i] = string(newLine)
			}
		}
	}

	// Print the result
	for _, line := range asciiArtLines {
		fmt.Println(line)
	}
}


