package impl

import "fmt"

func GoRoutineWithId(id int, ch chan int) {
	fmt.Println("GO routine with ID: ", id)
	ch <- id
}
