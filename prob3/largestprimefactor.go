package prob3

import (
	"fmt"

	"github.com/jimmyjames85/goprojecteuler/util"
)

// https://projecteuler.net/problem=3
//
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

func Run() {

	ch := make(chan int)
	go util.GeneratePrime(ch)

	num := 600851475143 //120069858

	q := num
	p := <-ch
	nl := false

	for q != 1 {
		if q%p == 0 {
			fmt.Printf("%d ", p)
			nl = true
			q /= p
		} else {
			if nl {
				fmt.Printf("\n")
				nl = false
			}
			p = <-ch
		}
	}
	fmt.Printf("\n\n")
	fmt.Printf("largest prime factor of %d is %d\n", num, p)
}
