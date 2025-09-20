package impl

func NumSquares(number int, sqCh chan int) {
	sqCh <- number * number
}
