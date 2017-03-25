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

// product is only used to send results through channels. a*b SHOULD equal p but is not enforced
type product struct {
	a int
	b int
	p int
}

func main() {
	findAll := false
	digits := 3

	ch := make(chan (product))
	go generateProductsOfNDigitNumbers(ch, digits)

	var p product
	var ok bool
	for p, ok = <-ch; ok && (findAll || !isPalindrome(p.p)); p, ok = <-ch {
		if findAll && isPalindrome(p.p) {
			fmt.Printf("%v\n", p)
		}
	}

	if !findAll {
		if ok && isPalindrome(p.p) {
			fmt.Printf("The largest palindrome made from the product of two %d-digit numbers is: %d * %d = %d\n", digits, p.a, p.b, p.p)
		} else {
			fmt.Printf("Failed to find palindrome.\n")
		}
	}

}

// generateProductsOfNDigitNumbers generates 999*999, 999*998, 998*998, 998*997, ... , 100*100
func generateProductsOfNDigitNumbers(ch chan<- product, digits int) {

	if digits < 0 {
		close(ch)
		return
	}

	start := 0
	for i := 0; i < digits; i++ {
		start = start*10 + 9
	}
	end := 1 + (start-9)/10

	generateDecliningThreeDigitNumbers := func(ch chan<- int) {
		for n := start; n >= end; n-- {
			ch <- n
		}
		close(ch)
	}

	chA := make(chan (int))
	chB := make(chan (int))
	go generateDecliningThreeDigitNumbers(chA)
	go generateDecliningThreeDigitNumbers(chB)

	readA := true
	a, aOk := <-chA
	b, bOk := <-chB
	for aOk && bOk {
		ch <- product{a: a, b: b, p: a * b}
		if readA {
			a, aOk = <-chA
		} else {
			b, bOk = <-chB
		}
		readA = !readA
	}
	close(ch)
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

// todo remove this and see if it's faster
func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -1 * n
}
