package impl

func PostfixToPrefix(exp string) string {
	stack := NewStack[string]()

	for _, ch := range exp {
		if IsOperand(ch) {
			stack.Push(string(ch))
		} else { // It's an operator
			if stack.Size() >= 2 {
				t1 := stack.Pop()
				t2 := stack.Pop()
				newStr := string(ch) + t2 + t1
				stack.Push(newStr)
			}
		}
	}
	return stack.Peek()
}
