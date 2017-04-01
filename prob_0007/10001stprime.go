package main

import (
	"fmt"

	"github.com/jimmyjames85/goprojecteuler/util"
	"time"
)

// https://projecteuler.net/problem=7
//
// 10001st prime
// -------------
//
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

func main() {
	start := time.Now()
	n := uint(10001)
	ch := make(chan uint)
	go util.GeneratePrimes(ch)

	var p uint
	for i := uint(0); i < n; i++ {
		p = <-ch
	}
	fmt.Printf("\rThe %d prime is %d\n", n, p)
	fmt.Printf("%s\n", time.Since(start))
}
