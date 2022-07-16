package Parser

import (
	. "github.com/Lucky31415/GoProjektarbeit/Expression"
	. "github.com/Lucky31415/GoProjektarbeit/Tokenizer"

	. "github.com/Lucky31415/GoProjektarbeit/Optional"
)

type Parser interface {
	Parse() Optional[Expression]
	parseE() Optional[Expression]
	parseE2(left Expression) Optional[Expression]
	parseT() Optional[Expression]
	parseT2(left Expression) Optional[Expression]
	parseF() Optional[Expression]
}

type ParserImpl struct {
	tokenizer Tokenizer
}

// Constructor
func NewParser(s string) *ParserImpl {
	return &ParserImpl{
		tokenizer: NewTokenizer(s),
	}
}

// Parsing Functions
func (p *ParserImpl) Parse() Optional[Expression] {
	return p.parseE()
}

func (p *ParserImpl) parseE() Optional[Expression] {
	t := p.parseT()
	if t.IsNothing() {
		return t
	}

	return p.parseE2(t.FromJust())
}

func (p *ParserImpl) parseE2(left Expression) Optional[Expression] {
	if p.tokenizer.GetToken() == PLUS {
		p.tokenizer.NextToken()

		right := p.parseT()
		if right.IsNothing() {
			return right
		}

		return p.parseE2(NewPlusExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p *ParserImpl) parseT() Optional[Expression] {
	f := p.parseF()
	if f.IsNothing() {
		return f
	}

	return p.parseT2(f.FromJust())
}

func (p *ParserImpl) parseT2(left Expression) Optional[Expression] {
	if p.tokenizer.GetToken() == MULT {
		p.tokenizer.NextToken()
		right := p.parseF()
		if right.IsNothing() {
			return right
		}
		return p.parseT2(NewMultExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p *ParserImpl) parseF() Optional[Expression] {
	switch p.tokenizer.GetToken() {
	case ZERO:
		p.tokenizer.NextToken()
		return Just[Expression](NewIntExp(0))
	case ONE:
		p.tokenizer.NextToken()
		return Just[Expression](NewIntExp(1))
	case TWO:
		p.tokenizer.NextToken()
		return Just[Expression](NewIntExp(2))
	case OPEN:
		p.tokenizer.NextToken()
		e := p.parseE()

		if e.IsNothing() {
			return e
		}

		if p.tokenizer.GetToken() != CLOSE {
			return Nothing[Expression]()
		}

		p.tokenizer.NextToken()
		return e
	default:
		return Nothing[Expression]()
	}

}
