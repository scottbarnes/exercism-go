// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// RemoveNonAscii removes all non-ascii characters from s.
// Note: this keeps spaces.
func removeNonAscii(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if r == ' ' {
			builder.WriteRune(r)
		}
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func createAbbreviation(s []string) string {
	var builder strings.Builder
	for _, word := range s {
		builder.WriteByte(word[0])
	}

	return builder.String()
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	hyphenSplitWords := strings.Split(s, "-")
	rejoinedWords := strings.Join(hyphenSplitWords, " ")
	onlyAscii := removeNonAscii(rejoinedWords)
	titleWords := strings.ToTitle(onlyAscii)
	splitWords := strings.Fields(titleWords)

	return createAbbreviation(splitWords)
}
