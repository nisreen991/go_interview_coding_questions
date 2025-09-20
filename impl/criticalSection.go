package impl

import (
	"fmt"
	"sync"
)

func CritialSectionExample() {
	var wg sync.WaitGroup
	var count int
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count = count + 1
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
