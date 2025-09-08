package main

import "fmt"

func main() {

	arr := []int{4, 2, 6, 3, 1, 1}
	x, y := findMissingAndRepeatingNumber(arr)
	fmt.Println("Repeating and Missing Number are: ", x, y)

}

func findMissingAndRepeatingNumber(arr []int) (int, int) {
	x, y := 0, 0
	var s int   //sum of array elements
	var sn int  //sum of n natural numbers
	var s2 int  //sum of squares of array elements
	var s2n int // sum of squares of n natural numbers
	n := len(arr)
	for i := 0; i < len(arr); i++ {
		s += arr[i]
		s2 += arr[i] * arr[i]
	}

	sn = n * (n + 1) / 2
	s2n = n * (n + 1) * (2*n + 1) / 6

	val1 := s - sn            //x-y
	val2 := (s2 - s2n) / val1 // x + y

	x = (val1 + val2) / 2
	y = val2 - x

	return x, y
}
