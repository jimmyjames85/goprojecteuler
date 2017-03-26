package main

import (
	"fmt"
	"strconv"
	"sync"
)

// https://projecteuler.net/problem=8
//
// Largest product in a series
// ---------------------------
//
// The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.
//
// 73167176531330624919225119674426574742355349194934
// 96983520312774506326239578318016984801869478851843
// 85861560789112949495459501737958331952853208805511
// 12540698747158523863050715693290963295227443043557
// 66896648950445244523161731856403098711121722383113
// 62229893423380308135336276614282806444486645238749
// 30358907296290491560440772390713810515859307960866
// 70172427121883998797908792274921901699720888093776
// 65727333001053367881220235421809751254540594752243
// 52584907711670556013604839586446706324415722155397
// 53697817977846174064955149290862569321978468622482
// 83972241375657056057490261407972968652414535100474
// 82166370484403199890008895243450658541227588666881
// 16427171479924442928230863465674813919123162824586
// 17866458359124566529476545682848912883142607690042
// 24219022671055626321111109370544217506941658960408
// 07198403850962455444362981230987879927244284909188
// 84580156166097919133875499200524063689912560717606
// 05886116467109405077541002256983155200055935729725
// 71636269561882670428252483600823257530420752963450
//
// Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

func main() {

	seriesLength := 13

	seqStr := `
73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450`

	greatest := product{}

	seq := strToSeq(seqStr)
	wg := sync.WaitGroup{}
	for i := 0; i <= len(seq)-seriesLength; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()

			product := 1
			var multipliers []int

			for offset := 0; offset < seriesLength; offset++ {
				multipliers = append(multipliers, seq[start+offset])
				product *= seq[start+offset]

				if product == 0 {
					// 0 times anything is 0...
					return
				}
			}

			// check if we are the greatest
			greatest.Lock()
			if product > greatest.product {
				// update
				greatest.product = product
				greatest.multipliers = multipliers
			}
			greatest.Unlock()

		}(i)
	}
	wg.Wait()

	fmt.Printf("The %d adjacent digits in the 1000-digit number that have the greatest product is \n\n\t", seriesLength)

	for i, m := range greatest.multipliers {
		if i == 0 {
			fmt.Printf("%d ", m)
		} else {
			fmt.Printf("* %d ", m)
		}
	}

	fmt.Printf("= %d\n", greatest.product)
}

// product is a way to keep track of the product of subsets from the sequence.
// The `product` _should_ equal the product of all the `multipliers`...but this is not enforced.
// Inherits sync.Mutex so we can lock/unlock in our goroutines from main
type product struct {
	sync.Mutex

	multipliers []int
	product int
}

//strToSeq parses only digits from the string `seq` and returns a slice of the parsed digits
func strToSeq(seq string) []int {
	var ret []int
	for _, d := range seq {
		i, err := strconv.Atoi(fmt.Sprintf("%c", d))
		if err == nil {
			ret = append(ret, i)
		}
	}
	return ret
}