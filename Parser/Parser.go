package Parser

import (
	. "github.com/Lucky31415/GoProjektarbeit/Expressions"
	. "github.com/Lucky31415/GoProjektarbeit/Tokenizer"

	. "github.com/Lucky31415/GoProjektarbeit/Optional"
)

type Parser struct {
	tokenizer Tokenizer
}

// Constructor
func NewParser(s string) Parser {
	return Parser{
		tokenizer: NewTokenizer(s),
	}
}

// Parsing Functions
func (p Parser) parseE() Optional[Expression] {
	t := p.parseT()
	if t.IsNothing() {
		return t
	}

	return p.parseE2(t.FromJust())
}

func (p Parser) parseE2(left Expression) Optional[Expression] {
	tokenizer := p.tokenizer

	if tokenizer.Token == PLUS {
		tokenizer.NextToken()

		right := p.parseT()
		if right.IsNothing() {
			return right
		}

		return p.parseE2(NewPlusExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p Parser) parseT() Optional[Expression] {
	f := p.parseF()
	if f.IsNothing() {
		return f
	}

	return p.parseT2(f.FromJust())
}

func (p Parser) parseT2(left Expression) Optional[Expression] {
	tokenizer := p.tokenizer
	if tokenizer.Token == MULT {
		tokenizer.NextToken()
		right := p.parseF()
		if right.IsNothing() {
			return right
		}
		return p.parseT2(NewMultExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p Parser) parseF() Optional[Expression] {
	tokenizer := p.tokenizer

	switch tokenizer.Token {
	case ZERO:
		tokenizer.NextToken()
		return Just[Expression](NewIntExp(0))
	case ONE:
		tokenizer.NextToken()
		return Just[Expression](NewIntExp(1))
	case TWO:
		tokenizer.NextToken()
		return Just[Expression](NewIntExp(2))
	case OPEN:
		tokenizer.NextToken()
		e := p.parseE()

		if e.IsNothing() {
			return e
		}

		if tokenizer.Token != CLOSE {
			return Nothing[Expression]()
		}

		tokenizer.NextToken()
		return e
	default:
		return Nothing[Expression]()
	}

}
