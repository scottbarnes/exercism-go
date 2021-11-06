package raindrops

import "strconv"

// Convert returns a string that contains raindrop sounds "corresponding to certain potenntial factors".
func Convert(number int) string {
	var r string
	if number%3 == 0 {
		r += "Pling"
	}

	if number%5 == 0 {
		r += "Plang"
	}

	if number%7 == 0 {
		r += "Plong"
	}

	if r == "" {
		r += strconv.Itoa(number)
	}

	return r
}
