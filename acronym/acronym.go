// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func getCleanWords(words string) string {
	var builder strings.Builder
	for _, r := range words {
		if unicode.IsSpace(r) || unicode.IsLetter(r) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func createAbbreviation(s []string) string {
	var builder strings.Builder
	for _, word := range s {
		r, _ := utf8.DecodeRuneInString(word)
		builder.WriteRune(r)
	}

	return builder.String()
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// A hyphen functions as a space in our acronym scheme.
	words := strings.ReplaceAll(s, "-", " ")
	cleanWords := getCleanWords(words)
	titleWords := strings.ToTitle(cleanWords)
	splitWords := strings.Fields(titleWords)

	return createAbbreviation(splitWords)
}
