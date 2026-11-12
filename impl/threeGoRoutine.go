// to print number 1 to 10 using 3 different go routines in order

package impl

import (
	"fmt"
	"sync"
)

func ThreeGoRoutine() {
	var wg sync.WaitGroup
	ch1, ch2, ch3 := make(chan bool), make(chan bool), make(chan bool)

	maxCounter := 10
	counter := 1
	wg.Add(3)

	go func() {
		defer wg.Done()
		for {
			<-ch1
			if counter > maxCounter {
				ch2 <- true
				return
			}
			fmt.Println("Go Routine 1:", counter)
			counter++
			ch2 <- true
		}

	}()

	go func() {
		defer wg.Done()
		for {
			<-ch2
			if counter > maxCounter {
				ch3 <- true
			}
			fmt.Println("Go Routine 2:", counter)
			counter++
			ch3 <- true
		}
	}()

	go func() {
		defer wg.Done()
		for {
			<-ch3
			if counter > maxCounter {
				ch1 <- true
			}
			fmt.Println("Go Routine 3:", counter)
			counter++
			ch1 <- true
		}
	}()

	ch1 <- true
	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)
}
