package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides takes three floats and returns the Kind of triangle it is, if any.
func KindFromSides(a, b, c float64) Kind {
	switch {
	// Not a triangle
	case isNaN(a, b, c):
		return NaT
	case a+b <= c || b+c <= a || a+c <= b:
		return NaT
	// Equilateral
	case a == b && b == c:
		return Equ
	// Isosceles
	case a == b || b == c || a == c:
		return Iso
	// Scalene
	case a != b && b != c && a != c:
		return Sca
	default:
		return NaT
	}

}

// isNaN returns true if any of the numbers are NaN
func isNaN(nums ...float64) bool {
	for _, v := range nums {
		if v <= 0 || math.IsNaN(v) || math.IsInf(v, 0) {
			return true
		}
	}
	return false
}
