package impl

import (
	"fmt"
	"sync"
)

const (
	jobs    = 5
	workers = 3
)

func workerMain() {
	jobList := []int{1, 2, 3, 4, 5, 56, 7}
	jobCh := make(chan int, len(jobList))
	var wg sync.WaitGroup
	for id := range workers {
		wg.Add(1)
		go worker(id, jobCh, &wg)
	}
	wg.Wait()
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// for worker :=  range workers{
	// wg.Add(1)
	for job := range jobs {
		fmt.Println("Worker", id, "prints job:", job)
	}
	// }
}
