package Expressions

import "fmt"

type Expression interface {
	Eval() int
	Pretty() string
}

// ---------------------------- IntExp ----------------------------
type IntExp struct {
	value int
}

func NewIntExp(value int) IntExp {
	return IntExp{
		value: value,
	}
}

func (ie IntExp) Eval() int {
	return ie.value
}

func (ie IntExp) Pretty() string {
	return fmt.Sprintf("%d", ie.value)
}

// ---------------------------- PlusExp ----------------------------
type PlusExp struct {
	l Expression
	r Expression
}

func NewPlusExp(left Expression, right Expression) PlusExp {
	return PlusExp{
		l: left,
		r: right,
	}
}

func (pe PlusExp) Eval() int {
	return (pe.l.Eval() + pe.r.Eval())
}

func (pe PlusExp) Pretty() string {
	return fmt.Sprintf("(%s + %s)", pe.l.Pretty(), pe.r.Pretty())
}

// ---------------------------- MultExp ----------------------------
type MultExp struct {
	l Expression
	r Expression
}

func NewMultExp(left Expression, right Expression) MultExp {
	return MultExp{
		l: left,
		r: right,
	}
}

func (me MultExp) Eval() int {
	return (me.l.Eval() * me.r.Eval())
}

func (me MultExp) Pretty() string {
	return fmt.Sprintf("(%s * %s)", me.l.Pretty(), me.r.Pretty())
}
