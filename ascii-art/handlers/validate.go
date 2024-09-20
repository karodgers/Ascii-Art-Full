package handlers

import (
	"unicode"
)

// Checks if a string contains non-ASCII characters i.e characters >127 in the ascii table.

func ContainsNonASCII(s string) bool {
	for _, char := range s {
		if char > unicode.MaxASCII {
			return true
		}
	}
	return false
}

/*

How it Works:

1. 	The code starts with importing the `unicode` package with the line `import "unicode"`.
	This package provides data and functions to test some properties of Unicode characters.
	Specifically, it provides the `unicode.MaxASCII` constant that represents the maximum value of ASCII characters.


3. 	Inside the function, there's a `for` loop with the line `for _, char := range s`. This loop iterates over each
	character (rune) in the string. The underscore `_` is used as a blank identifier, because the `range` keyword normally
	returns two values: the index and the value of the element at that index.

4. 	For each `char` (which is of type `rune` in Go, representing a Unicode code point), it checks whether the character's
	Unicode code point is greater than `unicode.MaxASCII` with the line `if char > unicode.MaxASCII`. The `unicode.MaxASCII`
	constant is the maximum value of an ASCII character, which is 127.

5. 	If any character code point is greater than this value, it means the character isn't part of the ASCII table and therefore
	the function returns `true` indicating the presence of non-ASCII characters in the string.

6. 	If the function iterates over the entire string without finding a character with a code point greater than `unicode.MaxASCII`,
	it means the string contains only ASCII characters. As a result, the function returns `false`.

*/
