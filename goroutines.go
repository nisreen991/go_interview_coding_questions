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

	evenCh := make(chan bool)
	oddCh := make(chan bool)
	dones := make(chan bool)

	go printOdd(evenCh, oddCh)
	go printEven(evenCh, oddCh, dones)

	oddCh <- true
	<-dones
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

func printEven(evenCh chan bool, oddCh chan bool, done chan bool) {
	for i := 2; i <= 10; i += 2 {
		<-evenCh
		fmt.Println("Even: ", i)
		if i == 10 {
			done <- true
		} else {
			oddCh <- true
		}
	}
}

func printOdd(evenCh chan bool, oddCh chan bool) {
	for i := 1; i <= 10; i += 2 {
		<-oddCh
		fmt.Println("Odd: ", i)
		evenCh <- true
	}
}
