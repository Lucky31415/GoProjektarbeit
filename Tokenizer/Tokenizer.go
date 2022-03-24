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

	t := Tokenizer{
		tokenize: tokenize,
		Token:    tokenize.next(),
	}

	return t
}

func (t *Tokenizer) NextToken() {
	t.Token = t.tokenize.next()
}

func (t *tokenize) next() Token_t {
	if len(t.s) <= t.pos {
		return EOS
	}

	for true {

		if len(t.s) <= t.pos {
			return EOS
		}

		switch string(t.s[t.pos]) {
		case "0":
			t.pos++
			return ZERO
		case "1":
			t.pos++
			return ONE
		case "2":
			t.pos++
			return TWO
		case "(":
			t.pos++
			return OPEN
		case ")":
			t.pos++
			return CLOSE
		case "+":
			t.pos++
			return PLUS
		case "*":
			t.pos++
			return MULT
		default:
			t.pos++
			break
		}
	}

	return EOS
}

func (t *tokenize) scan() []Token_t {
	v := []Token_t{}
	token := ONE

	for token != EOS {
		token = t.next()
		v = append(v, token)
	}

	return v
}

func (t *tokenize) show() string {
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
