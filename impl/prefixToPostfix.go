package impl

func PrefixToPostfix(exp string) string {

	stack := NewStack[string]()

	for i := len(exp) - 1; i >= 0; i-- {
		if IsOperand(rune(exp[i])) {
			stack.Push(string(exp[i]))
		} else { // It's an operator
			if stack.Size() >= 2 {
				t1 := stack.Pop()
				t2 := stack.Pop()
				newStr := t1 + t2 + string(exp[i])
				stack.Push(newStr)
			}
		}
	}
	return stack.Peek()
}
