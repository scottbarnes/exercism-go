// Package encode contains solution to Run Length Encoding task from https://exercism.io/my/tracks/go

package encode

import (
	"strconv"
	"strings"
)

// V2. Thanks @metalim.

type builder struct {
	strings.Builder
}

// writeEncoded is a method to write a strings.Builder, and I suppose 'hide' the level of
// nesting of conditional clauses and to potentially make the code easier to read.
func (b *builder) writeEncoded(count int, r rune) {
	switch {
	case count == 1:
		b.WriteRune(r)
	case count > 1:
		b.WriteString(strconv.Itoa(count))
		b.WriteRune(r)
	}
}

// RunLengthEncode encodes a string in Run Length Encoding. Takes a string and returns a string.
func RunLengthEncode(s string) string {
	var prev rune
	var count int
	var encoded builder

	for _, r := range s {
		if prev == r { // && count > 0 {
			count++
		} else {
			encoded.writeEncoded(count, prev)
			count = 1
			prev = r
		}
	}
	// write the last entry, i.e. at index len(s)-1, as the above loop will only write that
	// to prev and count but not to the builder.
	encoded.writeEncoded(count, prev)
	return encoded.String()
}

// RunLengthDecode decodes a Run Length Encoded string.
func RunLengthDecode(s string) string {
	var count int
	var decoded builder

	for _, r := range s {
		switch {
		case '0' <= r && r <= '9':
			// case unicode.IsDigit(r):
			// Turn single runes into numbers.
			// E.g., if there is a 2, set count = 2 (count * 0 is 0, + 2 = 2). If the next character is a letter,
			// switch goes to default and that letter is written twice, and stops when count is 0 again. But if
			// the next character is a 4, it will multiply the 2 by 10 to get 20, and add the 4, for 24.
			count = count*10 + int(r-'0')
		// For non-numbers, use count to write the appropriate number of characters.
		default:
			if count == 0 {
				count = 1
			}
			for ; count > 0; count-- {
				decoded.WriteRune(r)
			}
		}
	}
	return decoded.String()
}
