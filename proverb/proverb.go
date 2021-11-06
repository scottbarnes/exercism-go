// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb takes a slice of strings and returns a slice of strings in proverb format.
func Proverb(rhyme []string) []string {

	const (
		finalLine = "And all for the want of a %v."
		line      = "For want of a %v the %v was lost."
	)
	// Create a slice the length of the rhyme elements so there's no need to
	// expand+copy the slice with append().
	r := make([]string, len(rhyme))

	for i := range rhyme {
		// When at the end of the slice (or for the only element), set the index's
		// value to finalLine.
		if i == len(rhyme)-1 {
			r[i] = fmt.Sprintf(finalLine, rhyme[0])
		} else {
			r[i] = fmt.Sprintf(line, rhyme[i], rhyme[i+1])
		}
	}
	return r
}
