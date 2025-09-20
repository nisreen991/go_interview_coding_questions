package impl

import "fmt"

func PrintEven(evenCh chan bool, oddCh chan bool, done chan bool) {
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

func PrintOdd(evenCh chan bool, oddCh chan bool) {
	for i := 1; i <= 10; i += 2 {
		<-oddCh
		fmt.Println("Odd: ", i)
		evenCh <- true
	}
}
