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
	digits := uint(3)

	ch := make(chan (product))
	go generateProductsOfNDigitNumbers(ch, digits)

	var p, biggestPalindrome product
	var ok bool

	for p, ok = <-ch; ok; p, ok = <-ch {

		if p.isPalindrome() {
			if findAll {
				fmt.Printf("%v\n", p)
			}
			if p.p > biggestPalindrome.p {
				biggestPalindrome = p
			}
		}
	}

	if biggestPalindrome.p >= 0 {
		fmt.Printf("The largest palindrome made from the product of two %d-digit numbers is: %s\n", digits, biggestPalindrome)
	} else {
		fmt.Printf("Failed to find a palindrome.\n")
	}

}

func generateProductsOfNDigitNumbers(ch chan<- product, digits uint) {
	start, end := newNDigitRange(digits)
	for x := start; x >= end; x-- {
		for y := x; y >= end; y-- {
			ch <- product{
				a: x,
				b: y,
				p: x * y,
			}
		}
	}
	close(ch)
}

func newNDigitRange(digits uint) (int, int) {
	start := 0
	for i := uint(0); i < digits; i++ {
		start = start*10 + 9
	}
	end := 1 + (start-9)/10
	return start, end
}

func (p product) String() string {
	return fmt.Sprintf("[ %d * %d = %d ]", p.a, p.b, p.p)
}

func (p *product) isPalindrome() bool {
	return isPalindrome(p.p)
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
