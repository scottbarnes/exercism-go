// Package etl transforms input from map[int][]string to map[string]int format. More specifically
// it transforms and old Scrabble (R) scoring format into a new one.
package etl

import (
	"strings"
)

// Transform takes and old scrabble score and returns a new one, with any letters
// make lower case. input is map[int][]string and output is map[string]int
func Transform(i map[int][]string) map[string]int {
	var r = make(map[string]int)
	for score, letters := range i {
		for _, letter := range letters {
			r[strings.ToLower(letter)] = score
		}
	}

	return r
}
