package impl

import "fmt"

func InfixToPostfix(exp string) {
	// Stack of runes for the infix to postfix conversion
	stack := NewStack[rune]()
	result := ""

	for _, ch := range exp {
		// If scanned char is operand, add it to the result string
		if IsOperand(ch) {
			result += string(ch)
		} else if ch == '(' { //If the scanned character is '(', push it to Stack
			stack.Push(ch)
		} else if ch == ')' { // If the scanned character is ')', pop from Stack until '(' is encountered
			for !stack.IsEmpty() && stack.Peek() != '(' {
				result += string(stack.Pop())
			}
			if !stack.IsEmpty() {
				stack.Pop() // Pop '(' from stack
			}
		} else { //If an operator is scanned
			for !stack.IsEmpty() && (Priority(ch) <= Priority(stack.Peek())) {
				result += string(stack.Pop())
			}
			stack.Push(ch) // Push the current operator to the stack
		}
	}

	//Pop all the remaining elements from stack
	for !stack.IsEmpty() {
		result += string(stack.Pop())
	}

	fmt.Println("Postfix Expression: ", result)
}

func IsOperand(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
