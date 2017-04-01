package util

import (
	"fmt"
	"time"
)

// FROM: http://stackoverflow.com/questions/18157322/find-the-sum-of-all-the-primes-below-two-million-project-euler-c
// See HiddenTyphoon's answer
func GeneratePrimes(ch chan<- uint) {
	// Modified not to iterate over even numbers
	ch <- 2
	for i := uint(3); ; i += 2 {
		if IsPrime(i) {
			ch <- i
		}
	}
}

func IsPrime(n uint) bool {

	// TODO Fermat's (Little) Theorem:
	// If p is a prime and if a is any integer, a^p = a (mod p)
	// In particular, if p does not divide a, then a^p-1 = 1 (mod p). ([proof])
	//
	// If p is a prime and if a is any integer, then a^p = a (mod p)
	// 	- can be rewritten as -
	// If a^p != a (mod p) then p is not prime or a is not an integer
	//
	// Quick Test for non-primality
	//
	// Let a = 2
	// Then a is an integer
	// Case 1: a^p != a (mod p)
	// 	   then p is not prime
	// Case 2: a^p = a (mod p)
	//  	   then p might be prime (continue to test2)
	//
	//if n <= 1 {
	//	return false
	//}
	//a := uint(2)
	//aToTheP := a << (n -1)
	////fmt.Printf("2^%d = %d\n", p, aToTheP)
	////fmt.Printf("2^%d mod %d = %d", p , p , aToTheP % p )
	////fmt.Printf("2 mod %d = %d",  p ,a % p )
	//
	//if (aToTheP % n) != ( a % n) {
	//	fmt.Printf(".")
	//	if n == uint(7){
	//		panic("should not return false")
	//	}
	//
	//	return false
	//}

	// https://primes.utm.edu/prove/merged.html
	// Fermat's theorem gives us a powerful test for compositeness:
	// Given n > 1, choose a > 1 and calculate a^(n-1) modulo n (there is a very easy way to do quickly by repeated squaring, see the glossary page "binary exponentiation").
	// If the result is not one modulo n, then n is composite.
	// If it is one modulo n, then n might be prime so n is called a weak probable prime base a (or just an a-PRP).
	// Some early articles call all numbers satisfying this test pseudoprimes, but now the term pseudoprime is properly reserved for composite probable-primes.

	// QUICK TEST #2nd attempt
	// In particular, if n does not divide a, then a^(n-1) = 1 (mod n). ([proof])
	//
	// if a^(n-1) != 1 (mod n) then n divides a
	//
	// let a=2
	// if  2^(n-1) % n == 1  then n might be prime

	fmt.Println()
	a := uint(1)

	for i := uint(2); i < 20; i++ {
		start := time.Now()
		calc := (a << i)
		mod := calc % i
		end := time.Since(start)
		fmt.Printf("2^%d = %d  \t %d  %v\n", i, calc, mod, end)
	}

	panic("uhoh")
	end := 3
	c := a << (n - 1) % n

	if c != 1 {
		fmt.Printf("%d\t%v\n", n, end)
		return false
	}
	fmt.Printf("%d\t%v\n", n, end)
	if n > 10 {
		panic("stoping")
	}

	max := n
	for i := uint(2); i < max; i++ {
		if n%i == 0 {
			return false
		}
		max = n / i
	}
	return true
}
