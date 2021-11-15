package perfect

import "errors"

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("input must be positive")

// Classify takes an int64 and returns a Classification based on whether the
// number, based on its aliquot sum, perfect, deficient, or abundant.
func Classify(n int64) (Classification, error) {

	if n <= 0 {
		return 0, ErrOnlyPositive
	} else if n == 1 {
		return ClassificationDeficient, nil
	}

	var aliqoutSum int64 = 1
	for i := int64(2); i*i < n; i++ {
		if n%i == 0 {
			aliqoutSum += i
			aliqoutSum += n / i
		}
	}

	switch {
	case n == aliqoutSum:
		return ClassificationPerfect, nil
	case aliqoutSum < n:
		return ClassificationDeficient, nil
	case aliqoutSum > n:
		return ClassificationAbundant, nil
	default:
		return 0, ErrOnlyPositive
	}
}
