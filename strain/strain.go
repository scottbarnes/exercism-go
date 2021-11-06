package strain

type Ints []int
type Lists [][]int
type Strings []string

// (Ints) Keep returns a slice of ints where are true, of function f
func (i Ints) Keep(f func(int) bool) Ints {
	var r Ints
	// For each element in the slice, test if it's true of function f, and append if so.
	for _, v := range i {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// (Ints) Discard returns a slice of ints which are false, of function f
func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(i int) bool {
		return !f(i)
	})
}

// (Lists) Keep returns a slice of a slice of ints which are true, of function f.
func (l Lists) Keep(f func([]int) bool) Lists {
	var r Lists
	for _, v := range l {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func (s Strings) Keep(f func(string) bool) Strings {
	var r Strings
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}
