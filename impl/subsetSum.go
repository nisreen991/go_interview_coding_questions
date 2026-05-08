package impl

func findSubsetSum(index int, sum int, arr []int, subsetSum *[][]int) {
	if index == len(arr) {
		*subsetSum = append(*subsetSum, []int{sum})
		return
	}

	findSubsetSum(index+1, sum+arr[index], arr, subsetSum)

	findSubsetSum(index+1, sum, arr, subsetSum)
}

func SubsetSum(arr []int) [][]int {
	var subsetSum [][]int

	findSubsetSum(0, 0, arr, &subsetSum)

	return subsetSum
}
