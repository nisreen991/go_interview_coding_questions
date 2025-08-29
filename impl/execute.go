package impl

import "fmt"

func Execute() {

	selectionarr := []int{13, 6, 10, 44, 136, 1, 8, 2}
	SelectionSort(selectionarr)
	fmt.Println("Array sorted using Selection Sort:", selectionarr)

}
