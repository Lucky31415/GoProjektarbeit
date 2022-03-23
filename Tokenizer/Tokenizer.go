package main

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

type Tokenize struct {
	s   string
	pos int
}

type Tokenizer struct {
	tokenize Tokenize
	token    Token_t
}

func NewTokenizer(s string) Tokenizer {
	tokenize := Tokenize{
		s:   s,
		pos: 0,
	}

	return Tokenizer{
		tokenize: tokenize,
		token:    tokenize.Next(),
	}
}

func (t Tokenizer) NextToken() {
	t.token = t.tokenize.Next()
}

func (t Tokenize) Next() Token_t {
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

func (t Tokenize) Scan() []Token_t {
	v := []Token_t{}
	token := ONE

	for token != EOS {
		token = t.Next()
		v = append(v, token)
	}

	return v
}

func (t Tokenize) Show() string {
	v := t.Scan()
	s := ""

	for i := 0; i < len(v); i++ {
		s += string(v[i])

		if (i + 1) < len(v) {
			s += ";"
		}
	}

	return s
}
