package main

import "fmt"

// https://projecteuler.net/problem=2
//
// Even Fibonacci numbers
// ----------------------
//
// Each new term in the Fibonacci sequence is generated by adding the previous two terms.
// By starting with 1 and 2, the first 10 terms will be:
// 
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// 
// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.

func main() {

	ch := make(chan int)
	go generateFib(ch)

	sum := 0
	for fib := <-ch; fib <= 4000000; fib = <-ch {
		if fib%2 == 0 {
			fmt.Printf("%d\n", fib)
			sum += fib
		}
	}
	fmt.Printf("sum = %d\n", sum)
}

func generateFib(ch chan<- int) {
	f1 := 1
	f2 := 1
	for {
		ch <- f1
		f1, f2 = f2, f1 + f2
	}
}