package impl

import "math"

func NumberOfRotations(nums []int) int {
	mini := math.MaxInt
	n := len(nums)
	low := 0
	high := n - 1
	count := -1
	for low <= high {
		mid := (low + high) / 2
		if nums[low] <= nums[high] {
			if nums[low] < mini {
				count = low
			}
			break
		}

		if nums[low] <= nums[mid] {
			if nums[low] < mini {
				count = low
				mini = nums[low]
			}
			low = mid + 1
		} else {
			high = mid - 1
			if nums[mid] < mini {
				count = mid
				mini = nums[mid]
			}
		}
	}
	return count
}
