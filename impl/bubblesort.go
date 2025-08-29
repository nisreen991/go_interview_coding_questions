package impl

/*Pushing max element to last by adjacent swaps
Worst and Average case TC - O(n*n) Best Case - O(n) */
func BubbleSort(arr []int) []int {
	n := len(arr)

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
