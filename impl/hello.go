package impl

import "fmt"

func Hello(done chan bool) {
	fmt.Println("Hello World Goroutine!")
	done <- true
}
