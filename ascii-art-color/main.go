package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-color/pkg"
)

func main() {
	args := os.Args[1:]

	color, substring, argz, bannerFile, err := pkg.ParseFlag(args)
	if err != nil || len(argz) < 1 || color == "" || len(argz) == 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	fileName := "standard.txt"
	text := argz[0]

	if bannerFile != "" {
		fileName = bannerFile + ".txt"
	}
	if len(argz) > 1 {
		if len(argz) > 1 && len(argz) == 3 {
			if pkg.Checker(argz) {
				switch argz[len(argz)-1] {
				case "shadow", "standard", "thinkertoy":
					fileName = argz[len(argz)-1] + ".txt"
					text = strings.Join(argz[:len(argz)-1], " ")
				default:
					text = strings.Join(argz, " ")
				}
			}
		} else {
			fmt.Printf("Invalid syntax in use of  banner flag: %s\n\nAvailable banners; standard, shadow & thinkertoy\n", argz[len(argz)-1])
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\" \"banner\"")
			return
		}
	}

	err = pkg.ValidateFileChecksum(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	theMap, err := pkg.ReadFile(fileName)
	if err != nil {
		fmt.Println("error")
		return
	}

	// If substring is empty, color the whole string
	if substring == "" {
		substring = text
	}

	// Use PrintAsciiArt instead of directly calling PrintArt
	pkg.PrintAsciiArt(text, theMap, color, substring)
}
