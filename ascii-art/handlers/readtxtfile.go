package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
in this section of the function we focus on aquiring the ascii art
in the font files and subsequent storage of the stream of data aquired
in a slice for later aquisition. We first create an instance of using
flags in order to print the ascii in various formats found in different files
*/

const (
	asciiArtHeight = 8
)

func ReadAsciiArt() []string {
	var args []string

	args = append(args, os.Args[1:]...)

	var filename string

	switch args[len(args)-1] {
	case "-th":
		filename = "thinkertoy.txt"
	case "-sh":
		filename = "shadow.txt"
	default:
		filename = "standard.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	lineCount := 0

	scanner := bufio.NewScanner(file)

	var asciiArt []string
	var artLines []string

	for scanner.Scan() {

		lineCount++

		lines := scanner.Text()

		if len(lines) == 0 {
			continue
		}
		artLines = append(artLines, lines)

		if len(artLines) == asciiArtHeight {
			asciiArt = append(asciiArt, strings.Join(artLines, "\n"))
			artLines = nil
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading ASCII art file:", scanner.Err())
		os.Exit(1)
	}
	if len(asciiArt) == 0 {
		fmt.Println("Error: The ASCII art file is empty.")
		os.Exit(1)
	}

	if lineCount != 855 {
		fmt.Println("Read Error: Do not change the content of the txt file")
		os.Exit(1)
	}

	return asciiArt
}

/*
How it works:
1. Conatamt declaration

const asciiArtHeight = 8

2. Processing Command-Line Arguments

var args []string
args = append(args, os.Args[1:]...)

This code segment retrieves command-line arguments passed to the program, skipping the first argument
(`os.Args[0]`), which is the name of the executable itself.

2. File Opening Based on the Last Argument

The function checks the last argument to determine which file to open (`thinkertoy.txt`, `shadow.txt`, or by default `standard.txt`).
It then opens the specified file for reading.

2. Reading from the File

It creates a `bufio.Scanner` to read the file line by line. It uses a loop to read each line using `scanner.Scan()`.
Inside this loop, it skips empty lines and appends non-empty lines to a temporary slice `artLines`. Once `artLines`
reaches a length equal to `asciiArtHeight`, it joins these lines into one string (representing a complete ASCII art block)
and appends it to the `asciiArt` slice. Then, it resets `artLines` for the next block.

3. Error Checking After Scanning

if err := scanner.Err(); nil != err {
	log.Fatal(err)
}

After the loop, it checks if any error occurred during scanning. If so, it logs and halts the program.

4. Return ASCII art

The function finally returns the `asciiArt` slice, which contains all the ASCII art blocks read from the file.
*/
