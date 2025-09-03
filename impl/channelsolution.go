package impl

import (
	"fmt"
	"sync"
)

var z = 0

func incrementChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}

func ChannelSolutionMain() {
	var w sync.WaitGroup
	ch := make(chan bool, 1) //This buffered channel is used to ensure that only one Goroutine access the critical section of code which increments z. This is done by passing true to the buffered channel in line no. 11 just before z is incremented. Since the buffered channel has a capacity of 1, all other Goroutines trying to write to this channel are blocked until the value is read from this channel after incrementing z in line no. 13. Effectively this allows only one Goroutine to access the critical section.

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementChannel(&w, ch)
	}
	w.Wait()
	fmt.Println("Final Value of z is: ", z)
}
