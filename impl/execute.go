package impl

import "fmt"

func Execute() {
	arr1 := []int{15, -2, 2, -8, 1, 7, 10}
	longestSubarrayWithSumZero := longestSubarrayWithSumZero(arr1)
	arr2 := []int{4, 2, 2, 6, 4}
	numberofSubarraysWithXorK := numberofSubarraysWithXorK(arr2, 6)
	arr3 := []int{1, 2, 3}
	numberofSubarraysWithSumK := numberofSubarraysWithSumK(arr3, 3)
	fmt.Println("Longest Subarray With Sum Zero:", longestSubarrayWithSumZero)
	fmt.Println("Number of Subarrays With Xor K:", numberofSubarraysWithXorK)
	fmt.Println("Number of Subarrays With Sum K: ", numberofSubarraysWithSumK)
}
