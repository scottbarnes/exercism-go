package lsproduct

import "fmt"

// multiplySlice takes a string of numbers and returns their product, or an error if
// a non-number is included.
func multiplySlice(s string) (int64, error) {
	var product int64 = 1
	for _, i := range s {
		if !('0' <= i && i <= '9') {
			return 0, fmt.Errorf("invalid number: %v", i)
		}
		product *= int64(i - '0')
	}
	return product, nil
}

// LargestSeriesProduct takes a string of digits and a span, and returns the highest product
// of any length span within the string.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	var highProduct int64

	if span > len(digits) || span < 0 {
		return 0, fmt.Errorf("invalid span value: %v", span)
	}

	for i := 0; i <= len(digits)-span; i++ {

		num, err := multiplySlice(digits[i : i+span])
		if err != nil {
			return 0, err
		}

		if num > highProduct {
			highProduct = num
		}
	}
	return highProduct, nil
}
