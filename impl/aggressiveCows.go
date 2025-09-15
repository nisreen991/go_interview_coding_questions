package impl

func AggressiveCows(stalls []int, cows int) int {
	var ans int
	n := len(stalls)
	low, high := 1, stalls[n-1]-stalls[0]

	for low <= high {
		mid := (low + high) / 2
		if canBePlaced(stalls, mid, cows) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func canBePlaced(arr []int, dist int, cows int) bool {
	cntCows, last := 1, arr[0]
	for _, num := range arr {
		if num-last >= dist {
			cntCows += 1
			last = num
		}
	}
	if cntCows >= cows {
		return true
	} else {
		return false
	}
}
