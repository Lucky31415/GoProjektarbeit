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
func (p Parser) Parse() Optional[Expression] {
	return p.parseE()
}

func (p *Parser) parseE() Optional[Expression] {
	t := p.parseT()
	if t.IsNothing() {
		return t
	}

	return p.parseE2(t.FromJust())
}

func (p *Parser) parseE2(left Expression) Optional[Expression] {
	if p.tokenizer.Token == PLUS {
		p.tokenizer.NextToken()

		right := p.parseT()
		if right.IsNothing() {
			return right
		}

		return p.parseE2(NewPlusExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p *Parser) parseT() Optional[Expression] {
	f := p.parseF()
	if f.IsNothing() {
		return f
	}

	return p.parseT2(f.FromJust())
}

func (p *Parser) parseT2(left Expression) Optional[Expression] {
	if p.tokenizer.Token == MULT {
		p.tokenizer.NextToken()
		right := p.parseF()
		if right.IsNothing() {
			return right
		}
		return p.parseT2(NewMultExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p *Parser) parseF() Optional[Expression] {
	switch p.tokenizer.Token {
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

		if p.tokenizer.Token != CLOSE {
			return Nothing[Expression]()
		}

		p.tokenizer.NextToken()
		return e
	default:
		return Nothing[Expression]()
	}

}
