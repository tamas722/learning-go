package concurrentprimes

import (
	"sort"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(n int) []int {
	if n < 2 {
		return []int{}
	}

	primesChan := make(chan int)
	var wg sync.WaitGroup

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if isPrime(num) {
				primesChan <- num
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(primesChan)
	}()

	var primes []int
	for p := range primesChan {
		primes = append(primes, p)
	}

	sort.Sort(sort.IntSlice(primes))
	return primes
}
