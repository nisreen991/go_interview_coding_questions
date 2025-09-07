package impl

func numberofSubarraysWithSumK(arr []int, k int) int {
	prefixSum, count := 0, 0
	prefixMap := make(map[int]int)
	prefixMap[prefixSum] += 1 // {0,1}

	for i := 0; i < len(arr); i++ {
		prefixSum += arr[i]
		remove := prefixSum - k
		count += prefixMap[remove]
		prefixMap[prefixSum] += 1
	}
	return count
}
