package Tokenizer

type Token_t string

const (
	EOS   Token_t = "EOS"
	ZERO  Token_t = "ZERO"
	ONE   Token_t = "ONE"
	TWO   Token_t = "TWO"
	OPEN  Token_t = "OPEN"
	CLOSE Token_t = "CLOSE"
	PLUS  Token_t = "PLUS"
	MULT  Token_t = "MULT"
)

type tokenize struct {
	s   string
	pos int
}

type Tokenizer struct {
	tokenize tokenize
	Token    Token_t
}

func NewTokenizer(s string) Tokenizer {
	tokenize := tokenize{
		s:   s,
		pos: 0,
	}

	return Tokenizer{
		tokenize: tokenize,
		Token:    tokenize.next(),
	}
}

func (t Tokenizer) NextToken() {
	t.Token = t.tokenize.next()
}

func (t tokenize) next() Token_t {
	pos := t.pos
	s := t.s

	if len(s) <= pos {
		return EOS
	}

	for true {

		if len(s) <= pos {
			return EOS
		}

		switch s[pos] {
		case '0':
			pos++
			return ZERO
		case '1':
			pos++
			return ONE
		case '2':
			pos++
			return TWO
		case '(':
			pos++
			return OPEN
		case ')':
			pos++
			return CLOSE
		case '+':
			pos++
			return PLUS
		case '*':
			pos++
			return MULT
		default:
			pos++
			break
		}
	}

	return EOS
}

func (t tokenize) scan() []Token_t {
	v := []Token_t{}
	token := ONE

	for token != EOS {
		token = t.next()
		v = append(v, token)
	}

	return v
}

func (t tokenize) show() string {
	v := t.scan()
	s := ""

	for i := 0; i < len(v); i++ {
		s += string(v[i])

		if (i + 1) < len(v) {
			s += ";"
		}
	}

	return s
}
