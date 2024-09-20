package handlers

import (
	"fmt"
	"strings"
)

/*
in this function we are trying to extract the index (ascii index) of the individual characters that appears on the
string we are passsing in the command line so that we can match it with the positional index of that letter in the
slice which we had stored the ascii art from the ascii banner file
*/

func PrintLineByLine(text string, asciiArt []string) {
	for i := 0; i < asciiArtHeight; i++ {

		for _, char := range text {

			// Convert special characters to their corresponding index

			index := strings.IndexAny(" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", string(char))

			if index != -1 && index < len(asciiArt) {
				fmt.Print(strings.Split(asciiArt[index], "\n")[i]) // Print the ith line of the ASCII art for the current character
			}
		}

		fmt.Println()

	}
}

/*
How the function works:

1. Iteration Over ASCII Art Lines:

    for i := 0; i < asciiArtHeight; i++

    The loop iterates based on `asciiArtHeight` (which is a constant 8) representing the number of lines each ASCII art character spans.
    This loop is meant to go through each line of the ASCII representation that needs to be printed for the entire text.

2. Processing Each Character in `text`:

    for _, char := range text

    For every character in the provided `text` string, this inner loop handles how the character is represented in ASCII art based on the `asciiArt` mapping.

3. Finding Character Index:

    index := strings.IndexAny("...set of ascii charachters...", string(char))

    Using `strings.IndexAny`, it determines the position of the current character (`char`) in a specific set of known characters
    (covering spaces, symbols, numbers, and letters). This index helps find the matching ASCII art in the `asciiArt` slice.

4. Check Index and Print:

    if index != -1 && index < len(asciiArt)

    If `index` is valid (i.e., the character is found in the defined set, and it's within the bounds of `asciiArt`), the corresponding
    ASCII art line is printed:

    fmt.Print(strings.Split(asciiArt[index], "\n")[i])

    This line splits the ASCII art string for the found character by new lines and selects the specific line (`i`) which needs to be printed
    at this iteration.

5. New Line After Each Text Line:

    fmt.Println()

    After processing each character for the current line (handled by the outer loop `i`), a new line is initiated to correctly lay out the ASCII art.
*/
