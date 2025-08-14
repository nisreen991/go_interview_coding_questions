package main

import "fmt"

func main() {
	fmt.Println("Printing n fibonacci numbers:")
	n := 10
	fibonacci(n)
	isFib := isFibonacci(n)
	fmt.Println("Is number present in Fibonacci series: ", isFib)
	fmt.Println("Execution completed")
}

func fibonacci(n int) {
	if n <= 0 {
		fmt.Println("Enter a positive number")
		return
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Println(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

// function to check if a number is present in Fibonacci series or not

func isFibonacci(num int) bool {
	if num <= 0 {
		return false
	}

	a, b := 0, 1

	for a <= num {
		if a == num {
			return true
		}
		a, b = b, a+b
	}
	return false
}
