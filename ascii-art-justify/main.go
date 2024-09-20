package main

import (
	"fmt"
	"os"

	// "strings"

	"ascii/handlers"
)

func main() {
	args := os.Args[1:]
	fileName := ""
	align, argz, bannerFile, err := handlers.ParseFlag(args)
	if err != nil || len(argz) < 1 || align == "" {
		handlers.PrintUsage()
		return
	} else if len(argz) == 1 {
		fileName = "standard.txt"
	}

	if align != "center" && align != "left" && align != "right" && align != "justify" {
		handlers.PrintUsage()
		return
	}

	text := argz[0]
	if bannerFile != "" {
		fileName = bannerFile + ".txt"
	}
	if len(argz) > 1 {
		if len(argz) > 1 && len(argz) == 2 {
			if handlers.Checker(argz) {
				switch argz[len(argz)-1] {
				case "shadow", "standard", "thinkertoy":
					fileName = argz[len(argz)-1] + ".txt"
				default:
				}
			} else {
				handlers.PrintUsage()
			}
		} else {
			handlers.PrintUsage()
			return
		}
	}

	width, err := handlers.GetTerminalWidth()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	asciiArtMap, err := handlers.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading ASCII art:", err)
		return
	}

	if handlers.ContainsNonASCII(text) {
		fmt.Println("Error: Non-ASCII characters detected")
		return
	}

	handlers.PrintAsciiArt(text, asciiArtMap, align, width)
}
