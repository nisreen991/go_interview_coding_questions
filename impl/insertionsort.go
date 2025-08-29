package impl

/* Take an element and place it in its right order
TC - Worst and Average Case - O(n*n)
Best Case - O(n) */
func InsertionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
	return arr
}
