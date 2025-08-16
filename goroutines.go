package main

import "fmt"

func main() {
	fmt.Println("Writing a simple go routine:")
	done := make(chan bool)
	go hello(done)
	<-done

	number := 589
	sqCh := make(chan int)
	cuCh := make(chan int)
	go calcSquares(number, sqCh)
	go calcCubes(number, cuCh)
	squares, cubes := <-sqCh, <-cuCh
	fmt.Printf("Sum of Squares: %d and sum of Cubes: %d\n", squares, cubes)
	fmt.Println("Go routine execution completed!")
}

func hello(done chan bool) {
	fmt.Println("Hello World Goroutine!")
	done <- true
}

/*This program will print the sum of the squares and cubes of the individual digits of a number. We will structure the program such that the squares are calculated in a separate Goroutine, cubes in another Goroutine and the final summation happens in the main Goroutine. */

func calcSquares(number int, squareOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		number = number / 10
		sum += (digit * digit)
	}
	squareOp <- sum
}

func calcCubes(number int, cubeOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		number = number / 10
		sum += (digit * digit * digit)
	}
	cubeOp <- sum
}
