package Stack

func ifaceTest() Stack[int] {
	var stack Stack[int]
	stack = NewStack[int]()
	return stack
}

type Stack[T any] interface {
	Push(val T)
	Pop() T
	PopBack() T
	Top() T
	IsEmpty() bool
	GetSlice() []T
}

type StackImpl[T any] struct {
	stack []T
}

func NewStack[T any]() *StackImpl[T] {
	return &StackImpl[T]{
		stack: []T{},
	}
}

func (s *StackImpl[T]) Push(val T) {
	s.stack = append(s.stack, val)
}

func (s *StackImpl[T]) Pop() T {
	var frontVal T
	frontVal, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return frontVal
}

func (s *StackImpl[T]) PopBack() T {
	var backVal T
	backVal, s.stack = s.stack[0], s.stack[1:]
	return backVal
}

func (s StackImpl[T]) Top() T {
	return s.stack[0]
}

func (s StackImpl[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s StackImpl[T]) GetSlice() []T {
	return s.stack
}
