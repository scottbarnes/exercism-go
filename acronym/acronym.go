// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

func getSplitWords(words string) []string {
	result := []string{}
	spaceSplitWords := strings.Split(words, " ")

	for _, word := range spaceSplitWords {
		if strings.Contains(word, "-") {
			hyphenSplitWords := strings.Split(word, "-")
			result = append(result, hyphenSplitWords...)
		} else {
			result = append(result, word)
		}
	}

	return result
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var builder strings.Builder

	cleanWords := strings.ReplaceAll(s, "_", "")
	splitWords := getSplitWords(cleanWords)
	titleWords := func() []string {
		result := []string{}
		for _, word := range splitWords {
			result = append(result, strings.ToTitle(word))
		}
		return result
	}()

	for _, v := range titleWords {
		if len(v) > 0 {
			firstLetter := v[0]
			builder.WriteByte(firstLetter)
		}
	}

	return builder.String()
}
