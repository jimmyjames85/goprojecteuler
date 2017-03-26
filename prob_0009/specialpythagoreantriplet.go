package main

import "fmt"

// https://projecteuler.net/problem=9
//
// Special Pythagorean triplet
// ---------------------------
//
// A Pythagorean triplet is a set of three natural numbers, a < b < c , for which,
//
// a² + b² = c²
//
// For example, 3² + 4² = 9 + 16 = 25 = 5² .
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc .

func main() {
	ch := make(chan triplet)
	go generatePythagoreanTriplets(ch)

	var t triplet
	for t = <-ch; t.a+t.b+t.c != 1000; t = <-ch {
	}
	fmt.Printf("%d² + %d² = %d²\n\n", t.a, t.b, t.c)
	fmt.Printf("%d + %d + %d = 1000\n", t.a, t.b, t.c)
	fmt.Printf("%d * %d * %d = %d\n", t.a, t.b, t.c, (t.a * t.b * t.c))
}

func generatePythagoreanTriplets(ch chan<- triplet) {

	// Dickson's method
	// https://en.wikipedia.org/wiki/Formulas_for_generating_Pythagorean_triples
	//
	// 1: Let r be any even natural number
	// 2: For every factor-pair (s,t) such that st = r²/2, let
	//
	// 	x = r + s
	// 	y = r + t
	// 	z = r + s + t
	//
	// 3: Then (x,y,z) is a Pythagorean Triple, since:
	//
	// 	     x²       +        y²       =
	// 	  (r + s)²    +     (r + t)²	=
	// 	r² + 2rs + s² +  r² + 2rt + t²	=
	// 	r² + 2rs + s² + 2st + 2rt + t²	=	(From step 2)
	// 		 (r + s + t)²		= z²

	for r := uint(2); ; r += 2 {
		st := uint(r * r / 2)
		chP := make(chan factorPair)
		go generateFactorPairs(chP, st)
		for p, ok := <-chP; ok; p, ok = <-chP {
			ch <- triplet{a: r + p.s, b: r + p.t, c: r + p.s + p.t}
		}
	}
}

func generateFactorPairs(ch chan<- factorPair, n uint) {

	defer close(ch)
	if n <= 1 {
		ch <- factorPair{n, n}
	}

	seen := make(map[string]struct{})
	for i := uint(2); i <= n; i++ {
		if n%i == 0 {
			p := factorPair{(n / i), i}
			hash := fmt.Sprintf("(%d,%d)", p.s, p.t)
			if _, ok := seen[hash]; !ok {
				seen[hash] = struct{}{}
				ch <- p
			}
		}
	}
}

type factorPair struct {
	s uint
	t uint
}

type triplet struct {
	a uint
	b uint
	c uint
}
