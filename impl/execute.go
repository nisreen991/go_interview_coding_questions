package impl

import (
	"fmt"
	"slices"
)

func Execute() {
	fmt.Println("Returning index for the target element:")

	arr := []int{4, 5, 67, 2, 7, 2, 8, 16, 45} //array to be sorted for implementing binary search
	target := 3
	slices.Sort(arr)
	fmt.Println("Sorted Array: ", arr)
	var result = SearchIndexPosition(arr, target)
	fmt.Println("Index : ", result)

	rotatedArray := []int{3, 4, 5, 1, 2}
	fmt.Println("Number of array rotations:", NumberOfRotations(rotatedArray))

	n := 82
	fmt.Println("SquareRoot of", n, "is: ", SquareRoot(n))
}
