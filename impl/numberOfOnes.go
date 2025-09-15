package impl

// return the row index containing max number of ones, this is n x m matrix
func MaxNumberofOnes(matrix [][]int) int {
	index, maxCount := -1, 0
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		countRow := m - lowerBound(matrix[i], 1)

		if countRow > maxCount {
			maxCount = countRow
			index = i
		}

	}
	return index
}

func lowerBound(arr []int, x int) int {
	low, high := 0, len(arr)-1
	ans := len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] >= x {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}
