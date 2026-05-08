package impl

import "fmt"

func Execute() {
	expr := "a+b*(c^d-e)"
	fmt.Println("Infix Expression: ", expr)
	InfixToPostfix(expr)
}
