package util


func GenerateMultiplesOf(m int, ch chan <- int){

	if m<=0 {
		panic("expected m > 0")
	}

	for i:=1 ; ; i++{
		ch <- m*i
	}
}
