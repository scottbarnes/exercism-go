package lasagna

// TODO: define the 'OvenTime()' function
func OvenTime() int {
	return 40
}

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(i int) int {
	r := OvenTime() - i
	return r
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(l int) int {
	r := l * 2
	return r
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(l, m int) int {
	r := (l * 2) + m
	return r
}
