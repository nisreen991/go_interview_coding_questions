package impl

import (
	"fmt"
	"slices"
)

func Execute() {
	fmt.Println("Returning index for the target element:")

	arr := []int{4, 5, 67, 2, 7, 2, 8, 16, 45} //array to be sorted for implementing binary search
	target := 3
	slices.Sort(arr)
	fmt.Println("Sorted Array: ", arr)
	var result = SearchIndexPosition(arr, target)
	fmt.Println("Index : ", result)

	rotatedArray := []int{3, 4, 5, 1, 2}
	fmt.Println("Number of array rotations:", NumberOfRotations(rotatedArray))

	n := 82
	fmt.Println("SquareRoot of", n, "is: ", SquareRoot(n))

	m, root := 69, 4
	fmt.Println(root, "th root of ", m, "is:", NthRoot(root, m))

	stalls := []int{0, 3, 4, 7, 9, 10}
	cows := 4
	fmt.Println("Aggressive Cows:", AggressiveCows(stalls, cows))

	matrix := [][]int{{0, 0, 1, 1, 1}, {0, 0, 0, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {0, 1, 1, 1, 1}}
	fmt.Println("Row with Max number of ones: ", MaxNumberofOnes(matrix))

	medianMatrix := [][]int{{1, 5, 7, 9, 11}, {2, 3, 4, 5, 10}, {9, 10, 12, 14, 16}}
	fmt.Println("Median of given matrix is:", MatrixMedian(medianMatrix))
}
