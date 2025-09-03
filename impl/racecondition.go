package impl

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func RaceConditionMain() {
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("Final Value of x:", x)
}

/*We spawn 1000 increment Goroutines from line no. 20 of the program above. Each of these Goroutines run concurrently and the race condition occurs when trying to increment x in line no. 11 as multiple Goroutines try to access the value of x concurrently. */
