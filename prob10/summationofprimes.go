package main

import (
	"fmt"
	"github.com/jimmyjames85/goprojecteuler/util"
)

// https://projecteuler.net/problem=10
//
// Summation of primes
// -------------------
//
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

func main() {

 	threshold := uint(2000000) // uint(10)

	ch := make(chan uint)
	go util.GeneratePrimes(ch)

	sum := uint64(0)

	for p := <-ch; p < threshold; p = <-ch {
		fmt.Printf("\r %d %%", 100*p/threshold)
		sum += uint64(p)
	}

	fmt.Printf("\rThe sum of all the primes below %d is %d.\n", threshold, sum)
}