package handlers

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Println("Usage: go run . --align=<type> <STRING> <BANNER>")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
	os.Exit(0)
}
