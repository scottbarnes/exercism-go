package isbn

// v2. Thanks @ paiv.

// IsValidISBN takes a string and checks for ISBN-10 validity.
func IsValidISBN(input string) bool {
	var (
		count int
		isbn  int
	)

	// Iterate through input, looking for digits. Each time a digit is found, count++. Use the incrementing
	// count to subtract from 10, so it subtracts 10, 9, 8, etc., as we work through the equation.
	for _, char := range input {
		switch {
		// this is faster than unicode.IsChar(char)
		case '0' <= char && char <= '9':
			isbn += (10 - count) * int(char-'0')
			count++
		// 0-based counting, so count = 9 is the 10th possibly valid character.
		case count == 9 && char == 'X':
			isbn += 10
			count++
		}
	}
	// If the count is < or > 10 (meaning more than 10 characters), then it returns false.
	return count == 10 && isbn%11 == 0
}
