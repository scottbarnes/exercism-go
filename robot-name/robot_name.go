package robotname

import (
	"math/rand"
)

type Robot struct {
	name string
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"

var usedNames = make(map[string]bool)

// Name() generates a name if there isn't one, and returns it otherwise.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = getUniqueName()
	}
	return r.name, nil
}

// getUniqueName returns an unused name as a string.
func getUniqueName() string {
	var n string
	for {
		n = nameMaker()
		if !isNameUsed(n) {
			break
		}
	}
	return n
}

// isNameUsed returns true if a name is used and false otherwise.
func isNameUsed(name string) bool {
	// If the name is in the map, return true because the name is used.
	if _, ok := usedNames[name]; ok {
		return true
	}
	// Add the used name to the map
	usedNames[string(name)] = true
	return false
}

// nameMaker creates a name of the format LLNNN (L=letter, N=number)
func nameMaker() string {
	var n []byte
	for i := 0; i < 2; i++ {
		n = append(n, letters[rand.Intn(len(letters))])
	}

	for i := 0; i < 3; i++ {
		n = append(n, numbers[rand.Intn(len(numbers))])
	}

	return string(n)
}

// Reset() generates a new unique name.
func (r *Robot) Reset() {
	r.name = getUniqueName()
}
