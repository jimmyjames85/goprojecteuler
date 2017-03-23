package util
// FROM: https://golang.org/s/prime-sieve
// A concurrent prime sieve
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes.
func GeneratePrime(ch chan<- int) {
	ch0 := make(chan int) // Create a new channel.
	go generate(ch0)      // Launch generate goroutine.
	for i := 0; ; i++ {
		prime := <-ch0
		ch <- prime
		ch1 := make(chan int)
		go filter(ch0, ch1, prime)
		ch0 = ch1
	}
}
