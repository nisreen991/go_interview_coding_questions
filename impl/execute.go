package impl

import "fmt"

func Execute() {
	arr := []int{1, 2, 3, 4, 5}
	result := SubsetSum(arr)
	fmt.Println("Subset Sum: ", result)

}
