package impl

import "fmt"

func Execute() {
	expr := "a+b*(c^d-e)"
	fmt.Println("Infix Expression: ", expr)
	InfixToPostfix(expr)
	InfixToPrefix(expr)
	postfixExpr := "ab+cd^e-*"
	fmt.Println("Postfix Expression: ", postfixExpr)
	infixExpr := PostfixToInfix(postfixExpr)
	fmt.Println("Infix Expression: ", infixExpr)
}
