package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Printf("Process %d started\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Process %d finished\n", i)
	wg.Done()
}
