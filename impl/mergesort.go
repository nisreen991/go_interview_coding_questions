package impl

/* Time Complexity - O(n logn) Space Complexity - O(n)
This algorithm works with Divide & Merge */
func MergeSort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)
	Merge(arr, low, mid, high)
}

func Merge(arr []int, low int, mid int, high int) []int {
	temp := []int{}

	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left += 1
		} else {
			temp = append(temp, arr[right])
			right += 1
		}
	}

	for left <= mid { // if there are any elements remaining on left half of the array
		temp = append(temp, arr[left])
		left += 1
	}

	for right <= high { // if there are any elements remaining on right half of the array
		temp = append(temp, arr[right])
		right += 1
	}

	for i := low; i <= high; i++ { // To copy elements from temp to original array
		arr[i] = temp[i-low]
	}

	return arr
}
