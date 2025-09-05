package main

import "fmt"

func main() {
	arr := []int{15, -2, 2, -8, 1, 7, 10}
	ans := longestSubarrayWithSumZero(arr)
	fmt.Println(ans)
}

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
