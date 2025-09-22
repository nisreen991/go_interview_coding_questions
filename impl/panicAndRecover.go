package impl

import "fmt"

func divide(a, b int) int {
	return a / b
}

func recoverError() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic")
	}
}

func DivideMain() {
	defer recoverError()
	res := divide(10, 0)
	fmt.Println(res)
	fmt.Println("Divide Main finished execution")
}
