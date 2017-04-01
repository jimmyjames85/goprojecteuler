package util

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

	maxRange := n
	for i := uint(2); i < maxRange; i++ {
		if n%i == 0 {
			return false
		}
		maxRange = n / i
	}
	return true
}
