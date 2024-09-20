package handlers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) (map[rune]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	asciiMap := make(map[rune]string)

	currentChar := ' '
	currentArt := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentArt != "" {
				asciiMap[currentChar] = currentArt
				// Reset for the next character
				currentArt = ""
				currentChar++
			}
		} else {
			currentArt += line + "\n"
		}
	}

	// add the last character if the file doesn't end with an empty line
	if currentArt != "" {
		asciiMap[currentChar] = currentArt
	}

	return asciiMap, scanner.Err()
}
