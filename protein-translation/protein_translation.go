// Package protein translate RNA sequences into proteins.
package protein

import (
	"errors"
)

var ErrStop = errors.New("stop codon")
var ErrInvalidBase = errors.New("invalid base")
var codonToAA = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGG": "Tryptophan",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA takes an RNA string and returns a []string of polypeptides and an error value.
func FromRNA(s string) ([]string, error) {
	var proteins []string

	// Process the RNA string a codon at a time and append to the protein slice.
	for i := 0; i < len(s); i += 3 {
		protein, err := FromCodon(s[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}

	return proteins, nil
}

// FromCodon takes a codon string and returns a polypeptide string and error value.
func FromCodon(s string) (string, error) {
	if codon, ok := codonToAA[s]; ok {
		if codon == "STOP" {
			return "", ErrStop
		}
		return codon, nil
	}
	return "", ErrInvalidBase
}
