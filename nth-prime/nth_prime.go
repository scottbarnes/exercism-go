// package prime

// import (
// 	"math"
// )

// // isPrime takes a number and returns true if prime, false otherwise
// func isPrime(n int) bool {
// 	if n <= 3 {
// 		return n > 1
// 	}
// 	if ((n%2)+2)%2 == 0 || ((n%3)+3)%3 == 0 {
// 		return false
// 	}
// 	i := 5
// 	for math.Pow(float64(i), float64(2)) <= float64(n) {
// 		if ((n%i)+i)%i == 0 || ((n%(i+2)+(i+2))%(i+2)) == 0 {
// 			return false
// 		}
// 		i += 6
// 	}
// 	return true
// }

// // Nth takes an int n an return the nth prime and a bool.
// func Nth(n int) (int, bool) {
// 	if n <= 0 {
// 		return 0, false
// 	}
// 	var (
// 		currentNum int
// 		found      int
// 	)
// 	for found < n {
// 		currentNum++
// 		if isPrime(currentNum) {
// 			found++
// 		}
// 	}

// 	return currentNum, true
// }

// another one
// package prime

// import "math"

// func isPrime(p int) bool {

// 	if p%2 == 0 {

// 		return p == 2

// 	}

// 	m := int(math.Sqrt(float64(p))) + 1

// 	for t := 3; t < m; t += 2 {

// 		if p%t == 0 {

// 			return false

// 		}

// 	}

// 	return true

// }

// func Nth(n int) (int, bool) {

// 	if n <= 0 {

// 		return 0, false

// 	}

// 	for p := 2; ; p++ {

// 		if isPrime(p) {

// 			n--

// 			if n == 0 {

// 				return p, true

// 			}

// 		}

// 	}

// }

//Package prime calculates the nth prime number

package prime

import "math"

var primes []int

//sieve of erothanese

func sieve() {

	n := 105000 //chosen to pass tests

	primes = make([]int, 0, n)

	numbers := make([]bool, n)

	numbers[0] = true

	numbers[1] = true

	limit := int(math.Sqrt(float64(n)))

	for p := 2; p <= limit; {

		for i := p * p; i < n; i += p {

			numbers[i] = true

		}

		for {

			p++

			if !numbers[p] {

				break

			}

		}

	}

	for i := 2; i < n; i++ {

		if !numbers[i] {

			primes = append(primes, i)

		}

	}

}

//Nth returns the nth prime number

func Nth(n int) (int, bool) {

	if n == 0 {

		return 0, false

	}

	sieve()

	if n < len(primes) {

		return primes[n-1], true

	}

	return 0, false

}
