// package anagram takes a string and a []string of possible anagrams, and returns a []string of
// matching anagrams from the []string of possible anagrams.
package anagram

import (
	"strings"
	"unicode"
)

type counts [26]rune

// count takes a word string and returns a counts [26]rune. Specifically, it iterates through
// each letter, and, increments the value at the ndex corresponding to the letter (b/c "a" = 97,
// and subtracting 'a' subtracts 97, giving index 0 for "a", index 1 for "b", etc.)
func count(word string) counts {
	letterCount := [26]rune{}
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			letterCount[letter-'a']++
		}
	}
	return letterCount
}

// Detect takes a word string, a candidates []string, and returns a []string of anagrams
// from the candidates.
func Detect(word string, candidates []string) []string {
	var result []string
	lword := strings.ToLower(word)
	wc := count(lword)

	for _, candidate := range candidates {
		lcandidate := strings.ToLower(candidate)

		if len(word) != len(lcandidate) || strings.EqualFold(word, lcandidate) {
			continue
		}

		cc := count(lcandidate)

		if wc == cc {
			result = append(result, candidate)
		}
	}

	return result
}
