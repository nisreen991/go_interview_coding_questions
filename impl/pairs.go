package impl

func FindPairs(head *Node, sum int) [][]int {
	result := [][]int{}

	left := head
	right := FindTailOfDll(head)

	for left.Data < right.Data {
		pairs := []int{}

		if left.Data+right.Data == sum {
			pairs = append(pairs, left.Data)
			pairs = append(pairs, right.Data)
			result = append(result, pairs)
			left = left.Next
			right = right.Prev
		} else if left.Data+right.Data < sum {
			left = left.Next
		} else {
			right = right.Prev
		}
	}

	return result
}
