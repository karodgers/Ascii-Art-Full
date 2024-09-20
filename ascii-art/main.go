package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/handlers"
)

/*
the main function deals with ensuring there are correct number arguments passed in the command line.
since we have used flags to print in other formats we also ensure that
the flag argument is not printed and so we remove it,
*/

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: usage is; go run . \"sampletext\" ")
		os.Exit(1)
	}

	arguments := []string{}
	arguments = append(arguments, os.Args[1:]...)
	slice1 := []string{}

	if arguments[len(arguments)-1] == "-th" || arguments[len(arguments)-1] == "-sh" {

		slice1 = append(slice1, arguments[:len(arguments)-1]...)
		str := strings.Join(slice1, " ")
		text := strings.TrimSpace(str)

		asciiArt := handlers.ReadAsciiArt()

		if handlers.ContainsNonASCII(text) {
			fmt.Println("Print Errort: Non Ascii characters detected")
		} else {
			handlers.PrintAsciiArt(text, asciiArt)
		}
	} else {

		str := strings.Join(arguments, " ")
		text := strings.TrimSpace(str)
		asciiArt := handlers.ReadAsciiArt()

		if handlers.ContainsNonASCII(text) {
			fmt.Println("Print Error: Non Ascii characters detected")
		} else {
			handlers.PrintAsciiArt(text, asciiArt)
		}
	}
}

/*

How it works:

1. Command-line Arguments
   The program first checks if any command-line arguments have been provided.
   It does this by checking the length of `os.Args`, which is a slice containing all command-line arguments passed to the program.
   The first element of `os.Args` is the name of the program itself. If no additional arguments are provided, it prints an error message
   and exits.

2. Storing Arguments
   It initializes a slice of strings named `arguments` to store all arguments excluding the program name. This is accomplished by appending
   all elements of `os.Args`, starting from index 1, to the `arguments` slice.

3. Handling Flags
   The program then checks if the last argument is a flag (`-t` or `-sh`). These flags indicate a specific type of ASCII art that should be generated.

4. Preparation for ASCII Conversion
   If there is a flag:
   a. The program creates a new slice `slice1` and copies all arguments except for the last one (which is the flag) into it.
   b. It joins `slice1` into a single string `str` and trims any leading or trailing whitespace to get the final text `text`.
   c. The program calls `handlers.ReadAsciiArt()`, which likely reads in an ASCII art template from somewhere (the code for this isn't shown).
   d. It checks if `text` contains non-ASCII characters using `handlers.ContainsNonASCII(text)`. If it does, it prints a message and does not proceed with printing ASCII art.
   e. If the text is only ASCII characters, it then calls `handlers.PrintAsciiArt(text, asciiArt)` to print the ASCII representation of `text`.

5. No Flags
   If there is no flag, the program joins all arguments into a single string, trims it, and then follows the same steps as above to handle the conversion to ASCII art.

*/
