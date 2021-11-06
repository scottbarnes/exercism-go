package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a string of (potentially valid) nucleotides.
type DNA string

// Counts returns a Histogram with the count of each valid nucleotide in a DNS strang,
// or an error if the nuclotide is invalid.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	// For each letter in the string, if the letter isn't in the map, return an error.
	// Otherwise, increment the letter's value in the Histogram map.
	for _, letter := range d {
		if _, ok := h[letter]; !ok {
			return h, fmt.Errorf("invalid nucleotide: %v", letter)
		}
		h[letter]++
	}
	return h, nil
}
