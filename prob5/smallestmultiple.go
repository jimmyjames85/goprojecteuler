package main

import (
	"fmt"

	"github.com/jimmyjames85/goprojecteuler/util"
)

// https://projecteuler.net/problem=5
//
// Smallest multiple
// -----------------
//
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func main() {

	maxDivisor := 20

	// keys are primes, values are exponents
	factors := make(map[int]int)

	// prime decomposition of all number from 1 to max
	for d := 2; d <= maxDivisor; d++ {

		f := decompose(d)
		// p-prime e-exponent
		for p, e := range f {
			// if factors[p] > e then d divides the final product already
			factors[p] = max(factors[p], e)
		}
	}

	// calculate the final product
	product := 1
	for p, e := range factors {
		for i := 0; i < e; i++ {
			product *= p
		}
	}
	fmt.Printf("The smallest positive number that is evenly divisible by all of the numbers from 1 to %d is %d.\n", maxDivisor, product)
}

// decompose returns a map where the keys(p) are primes and values(e) are exponents such that the sum of all p^e = n
func decompose(n int) map[int]int {
	ret := make(map[int]int)
	ch := make(chan int)
	go util.GeneratePrime(ch)

	p := <-ch
	for n != 1 {
		if n%p == 0 {
			ret[p]++
			n /= p
		} else {
			p = <-ch
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
