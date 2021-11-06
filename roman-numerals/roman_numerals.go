package romannumerals

import (
	"fmt"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

// slice of arabicToRoman structs to hold the mapping while keeping the order.
var dict = []arabicToRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(i int) (string, error) {
	// Check for invalid inputs.

	if i <= 0 || i > 3000 {
		return "fail", fmt.Errorf("invalid input number: %v", i)
	}

	/* For each number i, see if its value is 'bigger' than a roman numeral. If it is, append
	that roman numeral to the result, and subtract its corresponding arabic number from i,
	and repeat until the numbers are exhausted.
	*/

	var result string

	for _, val := range dict {
		for i >= val.arabic {
			result += val.roman
			i -= val.arabic
		}
	}

	return result, nil
}
