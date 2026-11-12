package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	numWorkers := 3
	jobs := make(chan int, 10)

	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

}

func worker(id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
	time.Sleep(1 * time.Second)
}
