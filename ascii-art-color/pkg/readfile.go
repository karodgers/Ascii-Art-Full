package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strings"
)

func ReadFile(fileName string) (map[rune]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return map[rune]string{}, fmt.Errorf("error 1 %w", err)
	}
	defer file.Close()
	asciiMap := map[rune]string{}
	theRune := ' '
	var res strings.Builder
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()

		if line == "" {
			if res.Len() > 0 {
				asciiMap[theRune] = res.String()
				res.Reset()
				theRune++
			}
		} else {
			res.WriteString(line + "\n")
		}
	}
	if res.Len() > 0 {
		asciiMap[theRune] = res.String()
	}
	return asciiMap, nil
}
