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
	prefixExpr := "*-+ab^cde"
	fmt.Println("Prefix Expression: ", prefixExpr)
	infixExpr2 := PrefixToInfix(prefixExpr)
	fmt.Println("Infix Expression: ", infixExpr2)

	prefixExpr2 := PostfixToPrefix(postfixExpr)
	fmt.Println("Prefix Expression: ", prefixExpr2)

	postfixExpr2 := PrefixToPostfix(prefixExpr)
	fmt.Println("Postfix Expression: ", postfixExpr2)
}
