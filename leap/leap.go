/*
Package leap tells whether a year (int) is a leap year (bool).
*/
package leap

// v2. Thanks @danott. https://exercism.org/tracks/go/exercises/leap/solutions/48f5a7f1c9b086f99bf28168
// IsLeapYear taken a year as an int and returns a bool re whether the int is a leap year.
// See also @carylanou for a potentially more readable solution: https://exercism.org/tracks/go/exercises/leap/solutions/corylanou
func IsLeapYear(year int) bool {
	return isQuadrennialYear(year) && !isCentennialYear(year) || isQuadricentennialYear(year)
}

func isCentennialYear(year int) bool {
	return year%100 == 0
}

func isQuadrennialYear(year int) bool {
	return year%4 == 0
}

func isQuadricentennialYear(year int) bool {
	return year%400 == 0
}

// v1
// // IsLeapYear takes an int and returns a bool re whether a year is a leap year.
// func IsLeapYear(year int) bool {
// 	// If the year is evenly divisible by 4, it might be a leap year, but if not, it's not.
// 	if year%4 == 0 {
// 		switch {
// 		// Years evenly divisible by 4, 100, and 400 ARE leap years.
// 		case year%100 == 0 && year%400 == 0:
// 			return true
// 		// Years evenly divisible by just 4 and 100 are not leap years.
// 		case year%100 == 0:
// 			return false
// 		// Everything else evenly divisible by 4 is a leap year.
// 		default:
// 			return true
// 		}
// 	}
// 	return false
// }
