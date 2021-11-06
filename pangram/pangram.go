// Package pangram returns true if a string uses all 26 English letters.
package pangram

import (
	"strings"
)

// IsPanagram takes a string, and for each character in the English alphabet,
// checks if it's in the provided string. If a charcter isn't, return false.
// If every character is in the string, return true.
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for c := 'a'; c <= 'z'; c++ {
		if strings.IndexRune(s, c) == -1 {
			return false
		}
	}

	return true
}
