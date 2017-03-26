package main

import "fmt"

// https://projecteuler.net/problem=6
//
// Sum square difference
// ---------------------
//
// The sum of the squares of the first ten natural numbers is,
// 
// 1^2 + 2^2 + ... + 10^2 = 385
// 
// The square of the sum of the first ten natural numbers is,
// 
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// 
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// 
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

func main() {
	n := 100
	d := squareOfSum(n) - sumOfSquares(n)
	fmt.Printf("The difference between the sum of the squares of the first %d natural numbers and the square of the sum is %d\n", n, d)
}

func sumOfSquares(n int) int {
	return n*(2*n+1)*(n+1)/6
}

func squareOfSum(n int) int {
	return n*(n+1)*n*(n+1)/4
}
