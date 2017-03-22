package prob2

import "fmt"

// Problem Stmt:
//
// Each new term in the Fibonacci sequence is generated by adding the previous two terms.
// By starting with 1 and 2, the first 10 terms will be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.

func fibChan(n int, ch chan<- int) {
	if n < 0 {
		ch <- 0
		close(ch)
		return
	}

	last := 1
	cur := 1
	i := 0
	for {
		ch <- last
		next := last + cur
		last = cur
		cur = next
		i++
	}
	close(ch)
	return
}

func Run() {
	const n = 45

	fibCh := make(chan int)

	go fibChan(n, fibCh)

	opened := true
	var fib int

	fib, opened = <-fibCh
	sum := 0
	for opened {
		if fib > 4000000 {
			break
		}
		if fib%2 == 0 {
			fmt.Printf("%d\n", fib)
			sum += fib
		}
		fib, opened = <-fibCh
	}
	fmt.Printf("sum = %d\n", sum)
}
