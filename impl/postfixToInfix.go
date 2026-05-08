package impl

func PostfixToInfix(exp string) string {
	// Stack of strings for the postfix to infix conversion
	stack := NewStack[string]()

	for _, ch := range exp {
		// If scanned char is operand, push it to the stack
		if IsOperand(ch) {
			stack.Push(string(ch))
		} else {
			t1 := stack.Peek()
			stack.Pop()
			t2 := stack.Peek()
			stack.Pop()
			newExp := "(" + t2 + string(ch) + t1 + ")"
			stack.Push(newExp)
		}
	}
	return stack.Peek()
}
