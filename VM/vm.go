package VM

import (
	. "github.com/Lucky31415/GoProjektarbeit/Code"
	. "github.com/Lucky31415/GoProjektarbeit/Expression"
	. "github.com/Lucky31415/GoProjektarbeit/Optional"
	. "github.com/Lucky31415/GoProjektarbeit/Parser"
	. "github.com/Lucky31415/GoProjektarbeit/Stack"
)

// -------------------------- VM Interface --------------------------
type VM interface {
	push(i int)
	plus()
	mult()
	Run() Optional[int]
}

// -------------------------- VMImpl --------------------------
type VMImpl struct {
	codes []Code
	stack Stack[int]
}

func NewVm(codes []Code) *VMImpl {
	return &VMImpl{
		codes: codes,
		stack: NewStack[int](),
	}
}

func NewVmFromExp(exp Expression) *VMImpl {
	stack := NewStack[Code]()
	//codesFromExpression(exp, &stack)
	exp.Compile(stack)
	codes := stack.GetSlice()

	return &VMImpl{
		codes: codes,
		stack: NewStack[int](),
	}
}

func (vm *VMImpl) push(i int) {
	vm.stack.Push(i)
}

func (vm *VMImpl) plus() {
	right := vm.stack.Pop()
	left := vm.stack.Pop()
	vm.stack.Push(left + right)
}

func (vm *VMImpl) mult() {
	right := vm.stack.Pop()
	left := vm.stack.Pop()
	vm.stack.Push(left * right)
}

func (vm *VMImpl) Run() Optional[int] {
	vm.stack = NewStack[int]()

	for _, opCode := range vm.codes {
		switch opCode.Kind {
		case PUSH:
			vm.push(opCode.Val)
		case MULT:
			vm.mult()
		case PLUS:
			vm.plus()
		}
	}

	if vm.stack.IsEmpty() {
		return Nothing[int]()
	}

	return Just(vm.stack.Top())
}

// Static function
//func codesFromExpression(e Expression, codeStack *Stack[Code]) {
//	switch e.(type) {
//	case IntExp:
//		codeStack.Push(NewPush(e.Eval()))
//	case PlusExp:
//		v := e.(PlusExp)
//		codesFromExpression(v.L, codeStack)
//		codesFromExpression(v.R, codeStack)
//		codeStack.Push(NewPlus())
//	case MultExp:
//		v := e.(MultExp)
//		codesFromExpression(v.L, codeStack)
//		codesFromExpression(v.R, codeStack)
//		codeStack.Push(NewMult())
//	}
//}

func RunOnVM(rawExp string) Optional[int] {
	expOpt := NewParser(rawExp).Parse()
	if expOpt.IsNothing() {
		return Nothing[int]()
	}
	exp := expOpt.FromJust()

	vm := NewVmFromExp(exp)
	return vm.Run()
}
