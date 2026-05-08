package impl

func InfixToPrefix(exp string) {
	// Stack of runes for the infix to prefix conversion
	stack := NewStack[rune]()
	result := ""

	exp = reverseString(exp)
	exp = swapParentheses(exp)
	for _, ch := range exp {
		// If scanned char is operand, add it to the result string
		if IsOperand(ch) {
			result += string(ch)
		} else if ch == '(' { //If the scanned character is '(', push it to Stack
			stack.Push(ch)
		} else if ch == ')' { // If the scanned character is ')', pop from Stack until '(' is encountered
			for !stack.IsEmpty() && stack.Peek() != '(' {
				result += string(stack.Peek())
				stack.Pop()
			}
			stack.Pop()
		} else { //If an operator is scanned
			if ch == '^' {
				for !stack.IsEmpty() && (Priority(ch) <= Priority(stack.Peek())) {
					result += string(stack.Peek())
					stack.Pop()
				}
			} else {
				for !stack.IsEmpty() && (Priority(ch) < Priority(stack.Peek())) {
					result += string(stack.Peek())
					stack.Pop()
				}
			}
			stack.Push(ch) // Push the current operator to the stack
		}
	}

	for !stack.IsEmpty() {
		result += string(stack.Peek())
		stack.Pop()
	}

	result = reverseString(result)
	println("Prefix Expression: ", result)

}

func reverseString(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func swapParentheses(s string) string {
	runes := []rune(s)
	for i, ch := range runes {
		if ch == '(' {
			runes[i] = ')'
		} else if ch == ')' {
			runes[i] = '('
		}
	}
	return string(runes)
}
