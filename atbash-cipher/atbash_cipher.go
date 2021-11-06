package atbash

import (
	"strings"
)

// v2. Thanks @sjakobi.
// Atbash takes a string and returns it as an atbash cipher string.
func Atbash(s string) string {
	var atbash []rune

	for _, r := range s {
		/* 'z'-r+'a' ultimately subtracts the difference between 'z' and the letter in question.
		so 'Z'-'B'+'A' = 90 - 66 + 65 = 89 = 'Y'. So 'B' maps to 'Y'.
		Because A to Z and a to z the same number of characters, we can swap to lower case by
		subtracting capital letters from lower case 'z'. That moves to the 'left' by the same
		number of slots for either case. Then just add the appropriate case back.
		*/
		switch {
		case 'a' <= r && r <= 'z':
			atbash = append(atbash, 'z'-r+'a')
		case 'A' <= r && r <= 'Z':
			atbash = append(atbash, 'z'-r+'A')
		case '0' <= r && r <= '9':
			atbash = append(atbash, r)
		}
		// 'wrap' every 6 characters, where the 6th becomes a space, which is added once the
		// '5th' character is in place.
		if len(atbash)%6 == 5 {
			atbash = append(atbash, ' ')
		}
	}
	result := string(atbash)
	result = strings.TrimRight(result, " ")
	// return string(atbash)
	return result
}

// A pointlessly longer implementation.
// // atbashDict returns an English-language atbash dictionary.
// func atbashDict() map[rune]rune {
// 	var count rune
// 	dict := make(map[rune]rune)
// 	for r := 'a'; r <= 'z'; r++ {
// 		dict[r] = 'z' - count // counnt 'down' from the end as the count from 'a' goes 'up'
// 		count++
// 	}
// 	return dict
// }

// // parser takes a string and returns that string with only numbers and letters (and the letters
// // have been converted to lower case).
// func parser(s string) []rune {
// 	var output []rune
// 	for _, r := range s {
// 		if 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' {
// 			output = append(output, unicode.ToLower(r))
// 		} else if '0' <= r && r <= '9' {
// 			output = append(output, r)
// 		}
// 	}
// 	return output
// }

// // spacer returns a slice of rune with spaces inserted ever 'spaces' number of characters.
// func spacer(r []rune, spaces int) []rune {
// 	var output = []rune{}
// 	// With strings fewer than 'spaces' characters long, don't add spaces.
// 	if len(r) < spaces {
// 		return r
// 	}
// 	for i := 0; i <= len(r)-1; i += spaces {
// 		// If the index is with 'spaces' of the end, just append the remainder and exit.
// 		if i >= len(r)-spaces {
// 			output = append(output, ' ')
// 			output = append(output, r[i:]...)
// 			break
// 		}
// 		// Don't put a space at the start
// 		if i > 0 {
// 			output = append(output, ' ')
// 		}
// 		output = append(output, r[i:i+spaces]...)
// 	}
// 	return output
// }

// // Atbash takes a string and returns it as an atbash cipher string.
// func Atbash(s string) string {
// 	dict := atbashDict()
// 	sParsed := parser(s)
// 	var atbash []rune

// 	spaced := spacer(sParsed, 5)
// 	for _, r := range spaced {
// 		switch {
// 		case 'a' <= r && r <= 'z':
// 			atbash = append(atbash, dict[r])
// 		default:
// 			atbash = append(atbash, r)
// 		}
// 	}

// 	return string(atbash)
// }
