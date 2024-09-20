package handlers

import (
	"fmt"
	"strings"
)

/*
This Go code defines a package named `handlers` that contains a single function, `PrintAsciiArt`.
This function is responsible for converting a given input text into ASCII art based on the provided ASCII representations
passed in as an `asciiArt` slice (array). The function does not actually produce the ASCII art representations itself but
is designed to handle input text manipulations and call the other function `PrintLineByLine`, to do the actual ASCII printing.
*/

func PrintAsciiArt(text string, asciiArt []string) {
	if text == "" {
		return
	}

	specialChars := map[string]string{
		"\\t": "Tab",
		"\\b": "Backspace",
		"\\v": "Vertical Tab",
		"\\0": "Null",
		"\\f": "Form Feed",
		"\\r": "Carriage Return",
	}

	for spChar, description := range specialChars {
		if strings.Contains(text, spChar) {
			fmt.Printf("Print Error: Special ASCII character (%s) or (%s) detected \n", spChar, description)
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
				PrintLineByLine(word, asciiArt)
			}
		}
	} else {

		lines := strings.Split(text, "\n")
		for _, line := range lines {
			if len(line) > 0 {
				PrintLineByLine(line, asciiArt)
			} else {
				fmt.Println()
			}
		}
	}
}

/*
How the function works:

1. 	The function first checks if the input `text` is empty.
	If it is, the function simply returns without doing anything.

2. 	If the `text` contains the special escape characters "\t" (tab) or "\b" (backspace),
	the function prints a message indicating that special ASCII characters were detected but
	are not included in the scope of the program, and then it returns.

3.	If the `text` contains the escape character sequence "\\n", which is intended to
	allow for explicit new lines within the text, the function splits the text into lines using
	this sequence as a delimiter. It then processes each word (line segment separated by "\\n"):
   - If a word is empty, it increments a counter and prints a new line as long as there are more lines to process.
   - For non-empty words, it calls an unimplemented function `PrintLineByLine` for printing the ASCII art for each word.

4. 	If the `text` does not contain "\\n", the function will split the input on actual new line characters ("\n")
	to separate lines of the input text. It then processes each line:
   - For non-empty lines, it calls `PrintLineByLine` to generate the ASCII art for each line.
   - For empty lines, it prints a new line character.
*/
