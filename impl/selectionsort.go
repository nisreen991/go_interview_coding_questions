package impl

/* Best, Avg and Worst case time complexity is O(n*n)
This algorithm works by selecting the minimum element of the array and swap */
func SelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}

	return arr
}
