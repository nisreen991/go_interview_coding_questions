package impl

func ProduceNumbers(ch chan int) chan int {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	return ch
}
