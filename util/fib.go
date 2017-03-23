package util

func GenerateFib(ch chan<- int) {
	f1 := 1
	f2 := 1
	for {
		ch <- f1
		f3 := f1 + f2
		f1 = f2
		f2 = f3
	}
}
