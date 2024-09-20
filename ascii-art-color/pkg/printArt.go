package pkg

import (
	"fmt"
	"strings"
)

const asciiHeight = 8

func PrintArt(s string, p map[rune]string, color string, substring string) {
	for i := 0; i < asciiHeight; i++ {
		var b strings.Builder
		startIndex := 0
		for {
			index := strings.Index(s[startIndex:], substring)
			if index == -1 {
				b.WriteString(getArtLine(s[startIndex:], p, i, ""))
				break
			}
			b.WriteString(getArtLine(s[startIndex:startIndex+index], p, i, ""))
			b.WriteString(getArtLine(substring, p, i, color))
			startIndex += index + len(substring)
		}
		fmt.Println(b.String())
	}
}

func getArtLine(s string, p map[rune]string, lineNum int, color string) string {
	var b strings.Builder
	for _, v := range s {
		art, ok := p[v]
		if !ok {
			fmt.Println("char not found")
			return ""
		}
		artLines := strings.Split(art, "\n")
		if len(artLines) <= lineNum {
			continue
		}
		if color != "" {
			b.WriteString(applyColor(artLines[lineNum], color))
		} else {
			b.WriteString(artLines[lineNum])
		}
	}
	return b.String()
}

func applyColor(s, color string) string {
	// Implement color application based on your chosen color system
	// This is a placeholder implementation using ANSI color codes
	colorCodes := map[string]string{
		"black":     "\033[38;5;16m",
		"red":       "\033[38;5;196m",
		"green":     "\033[38;5;22m",
		"blue":      "\033[38;5;21m",
		"yellow":    "\033[38;5;226m",
		"magenta":   "\u001b[35m",
		"cyan":      "\u001b[36m",
		"white":     "\033[38;5;231m",
		"orange":    "\u001b[38;5;208m",
		"purple":    "\033[38;5;55m",
		"teal":      "\033[38;5;23m",
		"silver":    "\033[38;5;145m",
		"gray":      "\033[38;5;240m",
		"brown":     "\033[38;5;94m",
		"pink":      "\u001b[38;5;207m",
		"olive":     "\u001b[38;5;58m",
		"navy":      "\u001b[38;5;18m",
		"turquoise": "\u001b[38;5;80m",
		"lime":      "\033[38;5;46m",
		"indigo":    "\u001b[38;5;54m",
		"lavender":  "\u001b[38;5;183m",
		"charteuse": "\033[33m\033[34m",
		"salmon":    "\033[38;5;209m",
		"peach":     "\033[33m\033[96m",
		"seafoam":   "\033[32m\033[96m",
		"fuchsia":   "\033[38;5;201m",
		"violet":    "\033[33m\033[95m",
		"aqua":      "\033[38;5;51m",
		"maroon":    "\033[38;5;52m",
	}
	reset := "\033[0m"
	if code, ok := colorCodes[color]; ok {
		return code + s + reset
	}
	return s
}
