package accumulate

// Accumulate takes a slice of strings and a function (which itself takes a string
// as an argument) and returns a slice of strings.
func Accumulate(s []string, f func(string) string) []string {
	// Create an empty string slice the length of the input slice.
	var r = make([]string, len(s))
	// For each item in the input slice, set the result index to the corresponding
	// function value.
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}
