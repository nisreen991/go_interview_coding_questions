package impl

func longestSubarrayWithSumZero(nums []int) int {
	sum := 0
	maxLen := 0
	prefixMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == 0 {
			maxLen = max(maxLen, i+1)
		} else {
			_, exists := prefixMap[sum]
			if exists {
				maxLen = max(maxLen, i-prefixMap[sum])
			} else {
				prefixMap[sum] = i
			}
		}
	}
	return maxLen
}
