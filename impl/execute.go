package impl

import (
	"fmt"
)

func Execute() {
	fmt.Println("Writing a simple go routine:")
	done := make(chan bool)
	go Hello(done)
	<-done

	number := 589
	sqCh := make(chan int)
	cuCh := make(chan int)
	go CalcSquares(number, sqCh)
	go CalcCubes(number, cuCh)
	squares, cubes := <-sqCh, <-cuCh
	fmt.Printf("Sum of Squares: %d and sum of Cubes: %d\n", squares, cubes)

	evenCh := make(chan bool)
	oddCh := make(chan bool)
	dones := make(chan bool)

	go PrintOdd(evenCh, oddCh)
	go PrintEven(evenCh, oddCh, dones)

	oddCh <- true
	<-dones

	num := 10
	squareChannel := make(chan int)
	go NumSquares(num, squareChannel)
	squareResult := <-squareChannel
	fmt.Println("Square of given number: ", squareResult)

	idCh := make(chan int, 5)

	for i := range 5 {
		go GoRoutineWithId(i, idCh)
	}

	for range 5 {
		id := <-idCh
		fmt.Println(id)
	}
	fmt.Println("Go routine execution completed!")

	PrintMessages()

	producerCh := make(chan int)
	go ProduceNumbers(producerCh)

	for range 10 {
		fmt.Println("Number: ", <-producerCh)
	}
}
