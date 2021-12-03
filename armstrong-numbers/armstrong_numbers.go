package armstrong

// v2. Thanks @ozan.
// IsNumber returns true if a number is an Amstrong number.
func IsNumber(n int) bool {
	var total int

	digitCount := 0
	for number := n; number > 0; number /= 10 {
		digitCount++
	}

	for number := n; number > 0; number /= 10 {
		digit := number % 10
		power := digit
		for i := 1; i < digitCount; i++ {
			digit *= power
		}
		total += digit
	}

	// Append is slower than dividing twice and incrementing.
	// numbers := []int{}

	// for number := n; number > 0; number /= 10 {
	// 	numbers = append(numbers, number%10)
	// }

	// length := len(numbers)

	// for _, number := range numbers {
	// 	power := number
	// 	// Repeated multiplication avoids math.Pow
	// 	for i := 1; i < length; i++ {
	// 		number *= power
	// 	}
	// 	total += number
	// }

	return int(total) == n
}
