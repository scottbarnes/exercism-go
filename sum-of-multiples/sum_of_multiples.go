package summultiples

// v2. Thanks @hippeus.
// SumMultiples returns the total sum of all multiples of a set of numbers, up to limit.
func SumMultiples(limit int, divisors ...int) int {

	if limit < 2 || len(divisors) == 0 {
		return 0
	}
	var factorsUsed = make(map[int]struct{}) // Slightly faster than bool, which has 1 byte overhead. Thanks @hippeus.
	var total int
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for multiple := divisor; multiple < limit; multiple += divisor {
			if _, ok := factorsUsed[multiple]; !ok {
				factorsUsed[multiple] = struct{}{}
				total += multiple
			}
		}
	}
	// }
	return total
}

// Using goroutines, but with no channels and only locking. This is significantly slower than without
// go routines, perhaps because of the locking/unlocking.

// var wg sync.WaitGroup

// type Counter struct {
// 	m     sync.RWMutex
// 	value int
// 	seen  map[int]bool // Map initalized in the NewCounter constructor.
// }

// // CheckAdd determines whether to add a multiple to the map of seen multiples.
// func (c *Counter) CheckAdd(n int) {
// 	c.m.RLock()
// 	if _, ok := c.seen[n]; !ok {
// 		c.m.RUnlock()
// 		c.m.Lock()
// 		// defer c.m.Unlock() // creates non-trivial overhead.
// 		if _, ok := c.seen[n]; !ok {
// 			c.seen[n] = true
// 			c.value += n
// 		}
// 		c.m.Unlock()
// 	} else {
// 		c.m.RUnlock()
// 	}
// }

// // NewCounter creates a new counter, and does so in a way that initializes the map to avoid
// // nil map errors from: total := counter
// func NewCounter() *Counter {
// 	var c Counter
// 	c.seen = make(map[int]bool)
// 	return &c
// }

// // SumMultiples takes a limit and divisors and returns the sum of all multiples of the divisors
// // up to but not including the limit.
// func SumMultiples(limit int, divisors ...int) int {
// 	total := NewCounter()
// 	for number := 1; number < limit; number++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			for _, divisor := range divisors {
// 				if divisor != 0 && i%divisor == 0 {
// 					total.CheckAdd(i)
// 				}
// 			}
// 		}(number)

// 	}

// 	wg.Wait()

// 	return total.value
// }

// // Goroutines again, but this time with a channel and no locking. Massively slower than the original.
// var wg sync.WaitGroup
// var wgResult sync.WaitGroup

// // var multiples = make(chan int)

// type Counter struct {
// 	value int
// 	seen  map[int]bool // Map initalized in the NewCounter constructor.
// }

// // CheckAdd determines whether to add a multiple to the map of seen multiples.
// func (c *Counter) CheckAdd(n int) {
// 	if _, ok := c.seen[n]; !ok {
// 		if _, ok := c.seen[n]; !ok {
// 			c.seen[n] = true
// 			c.value += n
// 		}
// 	}
// }

// // NewCounter creates a new counter, and does so in a way that initializes the map to avoid
// // nil map errors from: total := counter
// func NewCounter() *Counter {
// 	var c Counter
// 	c.seen = make(map[int]bool)
// 	return &c
// }

// // SumMultiples takes a limit and divisors and returns the sum of all multiples of the divisors
// // up to but not including the limit.
// func SumMultiples(limit int, divisors ...int) int {
// 	multiples := make(chan int)
// 	total := NewCounter()
// 	wg.Add(limit - 1)
// 	for number := 1; number < limit; number++ {
// 		go func(i int) {
// 			defer wg.Done()
// 			for _, divisor := range divisors {
// 				if divisor != 0 && i%divisor == 0 {
// 					multiples <- i
// 				}
// 			}
// 		}(number)

// 	}

// 	wgResult.Add(1)
// 	go func() {
// 		defer wgResult.Done()
// 		for multiple := range multiples {
// 			total.CheckAdd(multiple)
// 		}
// 	}()

// 	wg.Wait()
// 	close(multiples)
// 	wgResult.Wait()

// 	return total.value
// }
