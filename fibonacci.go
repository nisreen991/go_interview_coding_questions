package main 

import "fmt" 

func main () {
	fmt.Println("Printing n fibonacci numbers:");
	n := 10
	fibonacci(n);
	fmt.Println("Execution completed");
}

func fibonacci(n int) {
	if n<=0 {
		fmt.Println("Enter a positive number");
		return
	}

	a,b := 0,1
	
	for i:=0; i<n; i++ {
		fmt.Println(a," ");
		a,b = b, a+b
	}
	fmt.Println()
}