package impl

/* 1. Pick a pivot and place it in its correct place in sorted array
2. Element smaller than pivot will be left on pivot else on right of pivot
3. Repeat 1. & 2.

Time Complexity : O(n logn)
Space Complexity: O(1)
*/
func QuickSort(arr []int, low int, high int) []int {
	if low < high {
		partition := pivot(arr, low, high)
		QuickSort(arr, low, partition-1)
		QuickSort(arr, partition+1, high)
	}
	return arr
}

func pivot(arr []int, low int, high int) int {
	pivotElement := arr[low]
	i, j := low, high

	for i < j {
		for arr[i] <= pivotElement && i <= high-1 {
			i += 1
		}

		for arr[j] > pivotElement && j >= low+1 {
			j -= 1
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
