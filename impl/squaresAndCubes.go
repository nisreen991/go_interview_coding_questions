package impl

/*This program will print the sum of the squares and cubes of the individual digits of a number. We will structure the program such that the squares are calculated in a separate Goroutine, cubes in another Goroutine and the final summation happens in the main Goroutine. */

func CalcSquares(number int, squareOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		number = number / 10
		sum += (digit * digit)
	}
	squareOp <- sum
}

func CalcCubes(number int, cubeOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		number = number / 10
		sum += (digit * digit * digit)
	}
	cubeOp <- sum
}
