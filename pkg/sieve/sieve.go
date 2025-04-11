package sieve

import (
	"errors"
	"math"
)

//go:generate mockgen -source ./$GOFILE -destination ./mock_$GOPACKAGE.go -package $GOPACKAGE

// import "math"

type Sieve interface {
	NthPrime(n int64) (int64, error)
}

type SieveService struct{}

func NewSieve() Sieve {
	return SieveService{}
}

func (s SieveService) NthPrime(n int64) (int64, error) {
	if n < 0 {
		return 0, errors.New("invalid input")
	}

	if n == 0 {
		return 2, nil
	}

	//start with an estimated limit and increase if needed
	limit := estimateLimit(n)
	primes := []int64{}

	for {
		primes = findPrimes(limit)
		if int64(len(primes)) > n {
			return primes[n], nil
		}
		// if we did not go high enough to find the nth prime, up the limit and try again
		limit = int64(float64(limit) * 1.2)
	}
}

func findPrimes(limit int64) []int64 {
	isComposite := make([]bool, limit) // all marked as false by default, composite will be marked as true
	isComposite[0] = true
	isComposite[1] = true

	// mark multiples of 2 as composite (except 2 itself)
	for i := int64(2); i < limit; i += 2 {
		if i != 2 {
			isComposite[i] = true
		}
	}

	// since we already marked all even numbers except for 2 as composites, we can loop through only odd numbers to check for primes
	for i := int64(3); i*i < limit; i += 2 {
		if !isComposite[i] { // if number is prime, mark all (non-even) multiples of that number as composite
			for j := int64(i * i); j < limit; j += 2 * i {
				isComposite[j] = true
			}
		}
	}
	primes := []int64{}
	for i := int64(0); i < int64(len(isComposite)); i++ {
		if !isComposite[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func estimateLimit(n int64) int64 {
	if n < 6 {
		return int64(15)
	}
	// using n * log(n) + n * log(log(n)) as a rough approximation
	estimate := float64(n)*math.Log(float64(n)) + float64(n)*math.Log(math.Log(float64(n)))
	return int64(estimate)
}
