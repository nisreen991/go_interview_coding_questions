package impl

import "fmt"

func Execute() {

	selectionarr := []int{13, 6, 10, 44, 136, 1, 8, 2}
	SelectionSort(selectionarr)
	fmt.Println("Array sorted using Selection Sort:", selectionarr)

	bubblearr := []int{13, 46, 24, 53, 20, 9}
	BubbleSort(bubblearr)
	fmt.Println("Array sorted using Bubble Sort: ", bubblearr)

	insertionarr := []int{14, 9, 15, 12, 6, 8, 13}
	InsertionSort(insertionarr)
	fmt.Println("Array sorted using Insertion Sort:", insertionarr)

	mergearr := []int{3, 1, 2, 4, 1, 5, 2, 6, 4}
	n := len(mergearr)
	MergeSort(mergearr, 0, n-1)
	fmt.Println("Array sorted using Merge Sort: ", mergearr)

}
