package impl

func Fibonacci(n int, ch chan int) {
	a, b := 0, 1

	for range n {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}
