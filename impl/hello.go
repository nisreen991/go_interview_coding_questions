package impl

import (
	"fmt"
	"sync"
)

func Hello(done chan bool) {
	fmt.Println("Hello World Goroutine!")
	done <- true
}

func PrintMessages() {
	var wg sync.WaitGroup
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	wg.Add(3)

	go func() {

		fmt.Println("m1")
		<-ch1
		ch2 <- true

	}()

	go func() {

		fmt.Println("m2")
		<-ch2
		ch3 <- true

	}()

	go func() {

		fmt.Println("m3")
		<-ch3
		ch1 <- true

	}()

	ch1 <- true
}
