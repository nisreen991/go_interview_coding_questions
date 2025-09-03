package impl

import (
	"fmt"
	"sync"
)

var y = 0

func incrementMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}
func MutexSolutionMain() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementMutex(&w, &m)
	}
	w.Wait()
	fmt.Println("Final value of y is: ", y)
}
