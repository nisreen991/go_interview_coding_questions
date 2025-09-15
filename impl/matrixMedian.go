package impl

import "math"

func upperBound(arr []int, x int) int {
	var ans int
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > x {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func countSmallEqual(arr [][]int, x int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		count += upperBound(arr[i], x)
	}
	return count
}

func MatrixMedian(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	req := (m * n) / 2
	low, high := math.MaxInt, math.MinInt

	for i := 0; i < m; i++ {
		low = min(low, matrix[i][0])
		high = max(high, matrix[i][m-1])
	}

	for low <= high {
		mid := (low + high) / 2
		smEq := countSmallEqual(matrix, mid)
		if smEq <= req {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
