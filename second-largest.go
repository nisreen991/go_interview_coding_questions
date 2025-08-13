package main

import "fmt"

func main() {
	fmt.Println("Printing second largest array element: ")
	var largestEle, secondLargestEle int
	arr := []int{7, 4, 8, 3, 9, 1, 2, 4}
	largestEle = largest(arr)
	secondLargestEle = secondLargest(arr)
	fmt.Println("Result : ", largestEle, secondLargestEle)
	fmt.Println("Execution completed!")
}

func largest(arr []int) int {
	largest := arr[0]

	for i := 0; i < len(arr); i++ {
		if largest < arr[i] {
			largest = arr[i]
		}
	}
	return largest
}

func secondLargest(arr []int) int {
	largest, secondLargest := arr[0], -1

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] < largest && arr[i] > secondLargest {
			secondLargest = arr[i]
		}
	}
	return secondLargest
}
