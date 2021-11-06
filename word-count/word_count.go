// package wordcount takes a string and returns a Frequency type of map[string]int
package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

// WordCount takes a phrase string and returns a Frequency.
func WordCount(phrase string) Frequency {
	count := make(Frequency)

	// Split on non-letter, non-digit, non-apostrophe characters.
	parsed := strings.FieldsFunc(phrase, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	})

	for _, word := range parsed {
		word = strings.Trim(strings.ToLower(word), "'")
		count[word]++
	}

	return count
}
