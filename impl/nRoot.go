package impl

//This fucntion calculates m^n
func MidN(mid int, n int, m int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ans * mid
		if ans > m { //to avoid overflow with large values of m, the statement is written
			return 2
		}
	}
	if ans == mid {
		return 1
	}
	return 0
}

func NthRoot(n int, m int) int {

	low, high := 1, m

	for low <= high {
		mid := (low + high) / 2
		if MidN(mid, n, m) == 1 {
			return mid
		} else if MidN(mid, n, m) == 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
