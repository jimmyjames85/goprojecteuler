package prob1

import "fmt"

// https://projecteuler.net/problem=1
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.

func Run() {

	ceiling := 1000 // 10

	multiples := make(map[int]struct{})
	ch3 := make(chan int)
	ch5 := make(chan int)

	go generateMultiplesOf(3, ch3)
	go generateMultiplesOf(5, ch5)

	complete3 := false
	complete5 := false

	for !(complete3 && complete5) {
		select {

		case m := <-ch3:
			if m < ceiling {
				multiples[m] = struct{}{}
			} else {
				complete3 = true
			}
		case m := <-ch5:
			if m < ceiling {
				multiples[m] = struct{}{}
			} else {
				complete5 = true
			}
		}
	}

	sum := 0
	for k, _ := range multiples {
		sum += k
	}

	fmt.Printf("The sum of all the multiples of 3 or 5 below 1000 is %d.\n", sum)

	//var sortme []int
	//for k, _ := range multiples {
	//	sortme = append(sortme, k)
	//}
	//sort.Ints(sortme)
	//for _, i := range sortme {
	//	fmt.Printf("%d\n", i)
	//}
}

func generateMultiplesOf(m int, ch chan<- int) {
	if m <= 0 {
		panic("expected m > 0")
	}
	for i := 1; ; i++ {
		ch <- m * i
	}
}
