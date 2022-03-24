package VM

import (
	. "github.com/Lucky31415/GoProjektarbeit/Optional"
	. "github.com/Lucky31415/GoProjektarbeit/Stack"
)

type OpCode_t string

const (
	PUSH OpCode_t = "PUSH"
	PLUS OpCode_t = "PLUS"
	MULT OpCode_t = "MULT"
)

// -------------------------- Code --------------------------
type Code struct {
	Kind OpCode_t
	Val  int
}

func newCode(o OpCode_t) Code {
	return Code{
		Kind: o,
		Val:  0,
	}
}

func newCode2(o OpCode_t, val int) Code {
	return Code{
		Kind: o,
		Val:  val,
	}
}

func NewPush(i int) Code {
	return newCode2(PUSH, i)
}

func NewPlus() Code {
	return newCode(PLUS)
}

func NewMult() Code {
	return newCode(MULT)
}

// -------------------------- VM --------------------------
type VM struct {
	codes []Code
	stack Stack[int]
}

func NewVm(codes []Code) VM {
	return VM{
		codes: codes,
		stack: NewStack[int](),
	}
}

func (vm *VM) Run() Optional[int] {
	stack := NewStack[int]()

	for _, opCode := range vm.codes {
		switch opCode.Kind {
		case PUSH:
			stack.Push(opCode.Val)
		case MULT:
			right := stack.Pop()
			left := stack.Pop()
			stack.Push(left * right)
		case PLUS:
			right := stack.Pop()
			left := stack.Pop()
			stack.Push(left + right)
		}
	}

	if stack.IsEmpty() {
		return Nothing[int]()
	}

	return Just(stack.Top())
}
