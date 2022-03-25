package Expression

import "fmt"

type Expression interface {
	Eval() int
	Pretty() string
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
