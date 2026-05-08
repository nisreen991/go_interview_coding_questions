package impl

// Generic Stack type using type parameters
type Stack[T any] struct {
	elements []T
}

// Creates a new generic Stack instance
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

// Push Data onto the Stack
func (s *Stack[T]) Push(data T) {
	s.elements = append(s.elements, data)
}

// Pop Element
func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	n := len(s.elements) - 1
	data := s.elements[n]
	s.elements = s.elements[:n]
	return data
}

// Peek at the Top element without removing it
func (s *Stack[T]) Peek() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if Stack is Empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

// Priority returns priority of the operators for infix, postfix and prefix operations
func Priority(ch rune) int {
	switch ch {
	case '^':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	default:
		return -1
	}
}
