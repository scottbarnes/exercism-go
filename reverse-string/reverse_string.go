// package reverse reverses a string.
package reverse

import "strings"

// Reverse takes a string and returns a reversed string.
func Reverse(s string) string {
	// Use strings.Builder so this doesn't choke as easily on large strings.
	var r strings.Builder
	// Work in runes so we can deal with multibyte unicode characters easily.
	runes := []rune(s)

	// Iterate in reverse through the runes.
	for i := len(runes) - 1; i >= 0; i-- {
		r.WriteRune(runes[i])
	}

	return r.String()
}
