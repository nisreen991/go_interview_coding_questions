package main

import (
	"fmt"
	"math"
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

	rotatedArray := []int{3, 4, 5, 1, 2}
	fmt.Println("Number of array rotations:", findMin(rotatedArray))
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

func findMin(nums []int) int {
	mini := math.MaxInt
	n := len(nums)
	low := 0
	high := n - 1
	count := -1
	for low <= high {
		mid := (low + high) / 2
		if nums[low] <= nums[high] {
			if nums[low] < mini {
				count = low
				//mini = nums[low]
			}
			break
		}

		if nums[low] <= nums[mid] {
			if nums[low] < mini {
				count = low
				mini = nums[low]
			}
			//mini = int(math.Min(float64(mini), float64(nums[low])))
			low = mid + 1
		} else {
			high = mid - 1
			if nums[mid] < mini {
				count = mid
				mini = nums[mid]
			}
			//mini = int(math.Min(float64(mini), float64(nums[mid])))
		}
	}
	return count
}
