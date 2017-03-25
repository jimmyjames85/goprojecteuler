package main

import "fmt"

// https://projecteuler.net/problem=4
//
// Largest palindrome product
// --------------------------
//
// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	ch := make(chan (int))
	go generateProductsOf3DigitNumbers(ch)

	for {
		p := <-ch
		fmt.Printf("received: %d\n", p)

		if isPalindrome(p) {
			fmt.Printf("FOUND IT\n")
			return
		}
	}
}

func isPalindrome(n int) bool {
	// respect negative palindromes
	str := fmt.Sprintf("%d", abs(n))
	strLen := len(str)
	mid := strLen/2 - ((strLen + 1) % 2)

	for i := 0; i <= mid; i++ {
		if str[i] != str[strLen-1-i] {
			return false
		}
	}
	return true
}

func generateProductsOf3DigitNumbers(ch chan<- int) {

	chA := make(chan (int))
	chB := make(chan (int))

	go generateDecliningThreeDigitNumbers(chA)
	go generateDecliningThreeDigitNumbers(chB)

	a := <-chA
	b := <-chB

	readA := true

	// todo this should not check against 101
	// todo but rather check if the channel is closed ?... need to look up how chans work
	for a >= 101 && b >= 101 {
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
		ch <- a * b

		if readA {
			a = <-chA
		} else {
			b = <-chB
		}

		readA = !readA
	}

	for n1 := 999; n1 > 99; n1-- {
		for n2 := 999; n2 > 99; n2-- {
			ch <- 2
		}
	}
}

func generateDecliningThreeDigitNumbers(ch chan<- int) {
	for n := 999; n > 99; n-- {
		ch <- n
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -1 * n
}
