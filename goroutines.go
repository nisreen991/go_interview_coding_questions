package main

import "fmt"

func main() {
	fmt.Println("Writing a simple go routine:")
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("Go routine execution completed!")
}

func hello(done chan bool) {
	fmt.Println("Hello World Goroutine!")
	done <- true
}
