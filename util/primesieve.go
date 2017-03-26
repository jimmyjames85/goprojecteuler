package util
// FROM: https://golang.org/s/prime-sieve
// A concurrent prime sieve
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- uint) {
	// Modified to skip over even numbers except for 2
	ch <- 2
	for i := uint(3); ; i+=2 {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan uint, out chan<- uint, prime uint) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes.
func GeneratePrimesWithSieve(ch chan<- uint) {
	ch0 := make(chan uint) // Create a new channel.
	go generate(ch0)      // Launch generate goroutine.
	for i := uint(0); ; i++ {
		prime := <-ch0
		ch <- prime
		ch1 := make(chan uint)
		go filter(ch0, ch1, prime)
		ch0 = ch1
	}
}
