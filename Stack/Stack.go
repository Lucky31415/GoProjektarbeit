package Stack

type Stack[T any] struct {
	stack []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		stack: []T{},
	}
}

func (s *Stack[T]) Push(val T) {
	s.stack = append(s.stack, val)
}

func (s *Stack[T]) Pop() T {
	var frontVal T
	frontVal, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return frontVal
}

func (s *Stack[T]) PopBack() T {
	var backVal T
	backVal, s.stack = s.stack[0], s.stack[1:]
	return backVal
}

func (s *Stack[T]) Top() T {
	return s.stack[0]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) GetSlice() []T {
	return s.stack
}
