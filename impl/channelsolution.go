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
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementChannel(&w, ch)
	}
	w.Wait()
	fmt.Println("Final Value of z is: ", z)
}
