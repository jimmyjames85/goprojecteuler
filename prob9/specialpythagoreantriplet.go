package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=9
//
// Special Pythagorean triplet
// ---------------------------
//
// A Pythagorean triplet is a set of three natural numbers, a < b < c , for which,
//
// a^2 + b^2 = c^2
//
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2 .
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc .

func main() {

	c := make(chan pair)

	for r := 2; r <=20 ; r++{
		
	}

	go generateFactorPairs(c, 64)
	for p, ok :=<-c; ok; p, ok = <-c{
		fmt.Printf("%v\n", p)
	}
	return

	ch := make(chan trip)
	go generatePythagoreanTriplets(ch)

	for i := 0; i < 10; i++ {
		t := <-ch
		fmt.Printf("%d %d %d\n", t.a, t.b, t.c)
	}
}

func generatePythagoreanTriplets(ch chan<- trip) {

	// todo this generate 1,1  1,2 2,2 ... but skips 1,7
	a := 1
	b := 1

	increaseA := false

	for {
		cSquared := a*a + b*b
		if c, ok := naturalSqrt(cSquared); ok {
			ch <- trip{
				a: a,
				b: b,
				c: c,
			}
		}

		if increaseA {
			a++
		} else {
			b++
		}
		increaseA = !increaseA
	}

}

// todo ensure a<b<c
type trip struct {
	a int
	b int
	c int
}

func isPythagoreanTriple(a, b int) bool {
	if b < a {
		a, b = b, a
	}
	return false
}

func newPair(a, b int)  pair {
	if b < a {
		a, b = b, a
	}
	return pair{a: a, b: b}
}

type pair struct {
	a int
	b int
}

func (p pair) String() string{
	return fmt.Sprintf("(%d,%d)", p.a, p.b)
}

func generateFactorPairs(ch chan<- pair, n int) {

	seen := make(map[string]struct{})

	if n <= 0 {
		panic("expecting positive n")
	} else if n == 1 {
		ch <- pair{a: 1, b: 1}
		return
	}

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			p := newPair(i, (n / i))
			if _, ok := seen[p.String()] ; !ok{
				seen[p.String()] = struct{}{}
				ch <- p
			}
		}
	}

	close(ch)
}

// todo not sure this working properly...
func naturalSqrt(n int) (int, bool) {
	r := math.Sqrt(float64(n))

	if r != float64(int(r)) {
		return 0, false
	}

	//if r*r == float64(n){
	//	fmt.Printf("\tr = %f   r2 = %f     n = %d \n", r, r*r, n)
	//
	//}

	return int(r), float64(n) == r*r

}
