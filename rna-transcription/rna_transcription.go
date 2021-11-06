package strand

import "strings"

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// v2
// ToRNA takes a string of DNA and returns its complement RNA string.
func ToRNA(dna string) string {
	return strings.Map(func(r rune) rune {
		return complements[r]
	}, dna)
}
