package allyourbase

import (
	"errors"
)

// v2. Thanks @ozan and @sount
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	var (
		inputBaseOutOfRange   = "input base must be >= 2"
		inputDigitsOutOfRange = "all digits must satisfy 0 <= d < input base"
		outputBaseOutOfRange  = "output base must be >= 2"
		outputDigits          = []int{}
		sum                   int
	)

	// Input validation
	if inputBase < 2 {
		return []int{0}, errors.New(inputBaseOutOfRange)
	}

	if outputBase < 2 {
		return []int{0}, errors.New(outputBaseOutOfRange)
	}

	// convert everything to base 10 (even things already in base 10)
	for _, digit := range inputDigits {
		sum = sum*inputBase + digit // Horner's Rule for converting base b to base 10.

		// Individual digit validation.
		if digit < 0 || digit >= inputBase {
			return []int{0}, errors.New(inputDigitsOutOfRange)
		}
	}

	// Catch empty input and all zero input.
	if sum == 0 {
		return []int{0}, nil
	}

	/* Convert from base 10 to base b by dividing the decimal (i.e. base 10 number)
	by the base until the quotient is 0. Calulate the remainder at each step and
	output it, in reverse order, to an []int. */
	dividend := sum
	for dividend > 0 {
		// Keep appending slice to the latest item, so the first item in is the last
		// in the list.
		outputDigits = append([]int{dividend % outputBase}, outputDigits...)
		dividend = dividend / outputBase
	}

	return outputDigits, nil
}
