package allyourbase

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Convert from source base to decimal (base 10 ) by multiplying each digit with the base raised to the power of the digit number (starting from right digit number 0):

// decimal = ∑(digit×basedigit number)
// Convert from decimal to destination base: divide the decimal with the base until the quotient is 0 and calculate the remainder each time. The destination base digits are the calculated remainders.

// toBaseTen converts []int of an arbitrary base to base 10.
func toBaseTen(inputBase int, inputDigits []int) []int {
	// baseTenDigits := make([]int, len(inputDigits))
	var baseTenDigits = []int{}
	var sum int

	// for i := len(inputDigits) - 1; i >= 0; i-- {
	for i := 0; i <= len(inputDigits)-1; i++ {
		sum += inputDigits[i] * int(math.Pow(float64(inputBase), float64(len(inputDigits)-1-i)))
	}

	stringDigits := fmt.Sprint(sum)
	for _, n := range stringDigits {
		baseTenDigits = append(baseTenDigits, int(n-'0'))
	}

	// baseTenDigits = append(baseTenDigits, sum)
	return baseTenDigits
}

// toBaseX converts []int from base ten to an arbitrary base X.
func toBaseX(outputBase int, inputDigits []int) []int {
	var stringDigits string
	var baseXDigits = []int{}
	var quotient int
	// var quotient = 1
	divisor := outputBase
	var remainder int

	for _, n := range inputDigits {
		stringDigits += strconv.Itoa(n)
	}
	joinedDigits, _ := strconv.Atoi(stringDigits)
	sliceJoinedDigits := []int{joinedDigits}

	for _, i := range sliceJoinedDigits {
		dividend := i
		for dividend > 0 {
			quotient = dividend / divisor
			// remainder += dividend % divisor
			remainder = dividend % divisor
			baseXDigits = append(baseXDigits, remainder)
			dividend = quotient
		}
	}
	// baseXDigits = append(baseXDigits, remainder)
	baseXDigitsReversed := []int{}
	for i := len(baseXDigits) - 1; i >= 0; i-- {
		baseXDigitsReversed = append(baseXDigitsReversed, baseXDigits[i])
	}

	// return baseXDigits
	return baseXDigitsReversed

}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	// outputDigits := make([]int, len(inputDigits))

	if inputBase <= 1 {
		return []int{}, errors.New("input base must be >= 2")
	} else if outputBase <= 1 {
		return []int{}, errors.New("output base must be >= 2")
	}

	var hasNonZero bool // track non-zero input.
	for _, digit := range inputDigits {
		switch {
		case digit <= -1 || digit >= inputBase:
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		case digit != 0:
			hasNonZero = true
		}
	}

	// Deal with all zero inputDigits
	if !hasNonZero {
		return []int{0}, nil
	}

	if inputBase == 10 {
		return toBaseX(outputBase, inputDigits), nil
	} else if outputBase == 10 {
		return toBaseTen(inputBase, inputDigits), nil
	} else {
		baseTen := toBaseTen(inputBase, inputDigits)
		return toBaseX(outputBase, baseTen), nil
	}
}
