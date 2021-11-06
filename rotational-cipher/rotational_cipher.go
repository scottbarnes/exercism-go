package rotationalcipher

// v2. Thanks @marcopaganini.

/*
rotate takes a rune, subtracts the base, adds the shift key, and performs modulus 26 to rotate
through the alphabet.

For a rune, subtract the base so that the numbers correspond to 0 to 25 (approximating a to z
letters in the alphabet). Add the shiftkey to shift the letters forward by shift key, and then
use modulus 26 so that the numbers 'wrap' every 26 characters. i.e. 0 + 26 = 0, 0 + 25 = 25,
25 + 1 = 0. Then add the base back to get 'a', 'z', and 'a' respectively.
*/
func rotate(r rune, base rune, shiftkey int) rune {
	return (((r - base) + rune(shiftkey)) % 26) + base
}

// Convert a string using a rotational cypher based on the shiftkey.
func RotationalCipher(plain string, shiftkey int) string {
	var rotOutput []rune

	for _, r := range plain {
		switch {
		case 'a' <= r && r <= 'z':
			rotOutput = append(rotOutput, rotate(r, 'a', shiftkey))
		case 'A' <= r && r <= 'Z':
			rotOutput = append(rotOutput, rotate(r, 'A', shiftkey))
		default:
			rotOutput = append(rotOutput, r)
		}
	}
	return string(rotOutput)
}

// v1. 30843 ns/op
// // RotationalCipher takes a string and a shift and shifts the letters of the string as many values
// // as the value of the shift key
// func RotationalCipher(plain string, shiftkey int) string {
// 	var rotOutput string
// 	/* This is a terrible mess, but for each rune in the input, subtract 'a' or 'A' to get the character
// 	number to something between 0 and 25, add the shiftkey number of values, uses modulus 26 so that the
// 	count wraps at 26, and then add 'a' or 'A' back as needed to get the appropriate letter.
// 	*/
// 	for _, r := range plain {
// 		if unicode.IsLower(r) {
// 			rotOutput += fmt.Sprintf("%v", string(((((r - 'a') + rune(shiftkey)) % 26) + 'a')))
// 		} else if unicode.IsUpper(r) {
// 			rotOutput += fmt.Sprintf("%v", string(((((r - 'A') + rune(shiftkey)) % 26) + 'A')))
// 		} else {
// 			rotOutput += fmt.Sprintf("%v", string(r))
// 		}
// 	}
// 	return rotOutput
// }
