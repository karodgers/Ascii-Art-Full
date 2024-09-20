package handlers_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"ascii/handlers"
)

func TestPrintLineByLine(t *testing.T) {
	asciiArtMap := map[rune]string{
		'A': "  A  \n A A \nA   A\nAAAAA\nA   A\nA   A\nA   A\n     ",
		'B': "BBBB \nB   B\nB   B\nBBBB \nB   B\nB   B\nBBBB \n     ",
		' ': "     \n     \n     \n     \n     \n     \n     \n     ",
	}

	tests := []struct {
		text     string
		align    string
		width    int
		expected []string
	}{
		{
			text:  "A B",
			align: "left",
			width: 15,
			expected: []string{
				"A     BBBBB     ",
				" A A  B   B     ",
				"A   A B   B     ",
				"AAAAA BBBBB     ",
				"A   A B   B     ",
				"A   A B   B     ",
				"A   A BBBBB     ",
				"                ",
			},
		},
		{
			text:  "AB",
			align: "center",
			width: 15,
			expected: []string{
				"   A     BBBB   ",
				"  A A    B   B  ",
				" A   A   B   B  ",
				"AAAAA    BBBB   ",
				"A   A    B   B  ",
				"A   A    B   B  ",
				"A   A    BBBB   ",
				"                ",
			},
		},
		{
			text:  "A B",
			align: "right",
			width: 15,
			expected: []string{
				"     A     BBBBB",
				"     A A  B   B ",
				"    A   A B   B ",
				"    AAAAA BBBBB ",
				"    A   A B   B ",
				"    A   A B   B ",
				"    A   A BBBBB ",
				"                ",
			},
		},
		{
			text:  "A B",
			align: "justify",
			width: 15,
			expected: []string{
				"A      BBBBB     ",
				" A A   B   B     ",
				"A   A  B   B     ",
				"AAAAA  BBBBB     ",
				"A   A  B   B     ",
				"A   A  B   B     ",
				"A   A  BBBBB     ",
				"                ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.align, func(t *testing.T) {
			handlers.PrintLineByLine(tt.text, asciiArtMap, tt.align, tt.width)
		})
	}
}

func TestChecker(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{
			input:    []string{"text", "standard"},
			expected: true,
		},
		{
			input:    []string{"text", "shadow"},
			expected: true,
		},
		{
			input:    []string{"text", "thinkertoy"},
			expected: true,
		},
		{
			input:    []string{"text", "random"},
			expected: false,
		},
		{
			input:    []string{"text"},
			expected: false,
		},
		{
			input:    []string{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.input, " "), func(t *testing.T) {
			result := handlers.Checker(tt.input)
			if result != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestPrintAsciiArt(t *testing.T) {
	asciiArtMap := map[rune]string{
		'A': "  A  \n A A \nAAAAA\nA   A\nA   A",
		'B': "BBBB \nB   B\nBBBB \nB   B\nBBBB ",
	}

	tests := []struct {
		name        string
		text        string
		align       string
		width       int
		expectError bool
	}{
		{
			name:        "Empty text",
			text:        "",
			align:       "left",
			width:       10,
			expectError: false,
		},
		{
			name:        "Unsupported special character",
			text:        "\\tHello",
			align:       "left",
			width:       10,
			expectError: true,
		},
		{
			name:        "Single line of text",
			text:        "AB",
			align:       "left",
			width:       10,
			expectError: false,
		},
		{
			name:        "Multiple lines of text",
			text:        "A\nB",
			align:       "left",
			width:       10,
			expectError: false,
		},
		{
			name:        "Newline character in text",
			text:        "A\\nB",
			align:       "left",
			width:       10,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output of fmt.Printf
			originalStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Run the function
			handlers.PrintAsciiArt(tt.text, asciiArtMap, tt.align, tt.width)

			// Close the writer and reset os.Stdout
			w.Close()
			os.Stdout = originalStdout

			// Read the captured output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			outputStr := buf.String()

			// Check for expected errors or correct output
			if tt.expectError {
				if !strings.Contains(outputStr, "Error: unsupported") {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if strings.Contains(outputStr, "Error: unsupported") {
					t.Errorf("Unexpected error: %s", outputStr)
				}
			}
		})
	}
}
