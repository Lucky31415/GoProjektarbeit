package Parser

//import "github.com/Lucky31415/GoProjektarbeit/Tokenizer"

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
func (p Parser) ParseE() Optional[Expression] {
	t := p.ParseT()
	if t.IsNothing() {
		return t
	}

	return p.ParseE2(t.FromJust())
}

func (p Parser) ParseE2(left Expression) Optional[Expression] {
	tokenizer := p.tokenizer

	if tokenizer.token == PLUS {
		tokenizer.NextToken()

		right := p.ParseT()
		if right.IsNothing() {
			return right
		}

		return p.ParseE2(NewPlusExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p Parser) ParseT() Optional[Expression] {
	f := p.ParseF()
	if f.IsNothing() {
		return f
	}

	return p.ParseT2(f.FromJust())
}

func (p Parser) ParseT2(left Expression) Optional[Expression] {
	tokenizer := p.tokenizer
	if tokenizer.token == MULT {
		tokenizer.NextToken()
		right := p.ParseF()
		if right.IsNothing() {
			return right
		}
		return p.ParseT2(NewMultExp(left, right.FromJust()))
	}

	return Just(left)
}

func (p Parser) ParseF() Optional[Expression] {
	tokenizer := p.tokenizer

	switch tokenizer.token {
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
		e := p.ParseE()

		if e.IsNothing() {
			return e
		}

		if tokenizer.token != CLOSE {
			return Nothing[Expression]()
		}

		tokenizer.NextToken()
		return e
	default:
		return Nothing[Expression]()
	}

}
