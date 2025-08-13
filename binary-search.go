package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Returning index for the target element:")

	arr := []int{4, 5, 67, 2, 7, 2, 8, 16, 45} //array to be sorted for implementing binary search
	target := 3
	slices.Sort(arr)
	fmt.Println("Sorted Array: ", arr)
	var result = searchIndexPosition(arr, target)
	fmt.Println("Index : ", result)
}

func searchIndexPosition(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// if element is not in array, it will be inserted at index low
	return low
}
