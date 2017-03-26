package main

import (
	"fmt"

	"github.com/jimmyjames85/goprojecteuler/util"
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

	n := 10001
	ch := make(chan int)
	go util.GeneratePrime(ch)

	var p int
	for i := 0; i < n; i++ {
		fmt.Printf("\r %d %%", 100*i/n)
		p = <-ch
	}
	fmt.Printf("\rThe %d prime is %d\n", n, p)
}
