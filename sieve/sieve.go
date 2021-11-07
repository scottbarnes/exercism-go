package sieve

// Thanks @jboelter.
// Sieve uses a sieve of Eratosthenes to return all primes between 2 and a limit.
func Sieve(limit int) []int {
	marked := make([]bool, limit+1) // Store marked non-primes.
	primes := make([]int, limit/2)  // No even primes > 2.
	primeIndex := 0                 // Index within the slice.

	// Continue when finding a marked number, as it's non-prime.
	for number := 2; number <= limit; number++ {
		if marked[number] {
			continue // it's a prime
		}

		// Add non-marked numbers as primes.
		primes[primeIndex] = number
		primeIndex++

		// For each non-marked number (i.e. each prime), mark all of its multiples,
		// as they are non-primes.
		for n := number + number; n <= limit; n += number {
			marked[n] = true
		}
	}

	return primes[:primeIndex]
}
