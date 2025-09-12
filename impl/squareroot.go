package impl

func SquareRoot(n int) int {
	low, high := 1, n

	ans := 1

	for low <= high {
		mid := (low + high) / 2
		if (mid * mid) <= n {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
