package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 5, 7, 8}
	result := reverseArray(arr)
	fmt.Println("Reversed Array:", result)
	fmt.Println("Reversed in pairs:", reverseInPairs(arr))
}

func reverseArray(arr []int) []int {
	i, j := 0, len(arr)-1

	for i <= j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

func reverseInPairs(arr []int) []int {
	for i := 0; i < len(arr)-1; i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	return arr
}
