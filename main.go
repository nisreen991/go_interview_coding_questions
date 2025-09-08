package main

import "fmt"

func main() {
	arr := []int{5, 3, 2, 4, 1}

	ans := numberOfInversions(arr)

	fmt.Println("Number of inversions: ", ans)

}

func numberOfInversions(arr []int) int {
	return mergeSort(arr, 0, len(arr)-1)
}

/*
	Time Complexity - O(n logn) Space Complexity - O(n)

This algorithm works with Divide & Merge
*/
func mergeSort(arr []int, low int, high int) int {
	count := 0
	if low >= high {
		return count
	}
	mid := (low + high) / 2
	count += mergeSort(arr, low, mid)
	count += mergeSort(arr, mid+1, high)
	_, cnt := merge(arr, low, mid, high)
	count += cnt
	return count
}

func merge(arr []int, low int, mid int, high int) ([]int, int) {
	temp := []int{}

	left := low
	right := mid + 1
	count := 0
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left += 1
		} else {
			temp = append(temp, arr[right])
			count += (mid - left + 1)
			right += 1
		}
	}

	for left <= mid { // if there are any elements remaining on left half of the array
		temp = append(temp, arr[left])
		left += 1
	}

	for right <= high { // if there are any elements remaining on right half of the array
		temp = append(temp, arr[right])
		right += 1
	}

	for i := low; i <= high; i++ { // To copy elements from temp to original array
		arr[i] = temp[i-low]
	}

	return arr, count
}
