package Expression

import (
	"fmt"

	. "github.com/Lucky31415/GoProjektarbeit/Code"
	. "github.com/Lucky31415/GoProjektarbeit/Stack"
)

type Expression interface {
	Eval() int
	Pretty() string
	Compile(codeStack Stack[Code])
}

// ---------------------------- IntExp ----------------------------
type IntExp struct {
	Value int
}

func NewIntExp(value int) IntExp {
	return IntExp{
		Value: value,
	}
}

func (ie IntExp) Eval() int {
	return ie.Value
}

func (ie IntExp) Pretty() string {
	return fmt.Sprintf("%d", ie.Value)
}

func (ie IntExp) Compile(codeStack Stack[Code]) {
	codeStack.Push(NewPush(ie.Eval()))
}

// ---------------------------- PlusExp ----------------------------
type PlusExp struct {
	L Expression
	R Expression
}

func NewPlusExp(left Expression, right Expression) PlusExp {
	return PlusExp{
		L: left,
		R: right,
	}
}

func (pe PlusExp) Eval() int {
	return (pe.L.Eval() + pe.R.Eval())
}

func (pe PlusExp) Pretty() string {
	return fmt.Sprintf("(%s + %s)", pe.L.Pretty(), pe.R.Pretty())
}

func (pe PlusExp) Compile(codeStack Stack[Code]) {
	pe.L.Compile(codeStack)
	pe.R.Compile(codeStack)
	codeStack.Push(NewPlus())
}

// ---------------------------- MultExp ----------------------------
type MultExp struct {
	L Expression
	R Expression
}

func NewMultExp(left Expression, right Expression) MultExp {
	return MultExp{
		L: left,
		R: right,
	}
}

func (me MultExp) Eval() int {
	return (me.L.Eval() * me.R.Eval())
}

func (me MultExp) Pretty() string {
	return fmt.Sprintf("(%s * %s)", me.L.Pretty(), me.R.Pretty())
}

func (me MultExp) Compile(codeStack Stack[Code]) {
	me.L.Compile(codeStack)
	me.R.Compile(codeStack)
	codeStack.Push(NewMult())
}
