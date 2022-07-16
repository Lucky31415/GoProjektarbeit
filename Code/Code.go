package Code

type OpCode_t string

const (
	PUSH OpCode_t = "PUSH"
	PLUS OpCode_t = "PLUS"
	MULT OpCode_t = "MULT"
)

// -------------------------- Code --------------------------
type Code struct {
	Kind OpCode_t
	Val  int
}

func newCode(o OpCode_t) Code {
	return Code{
		Kind: o,
		Val:  0,
	}
}

func newCode2(o OpCode_t, val int) Code {
	return Code{
		Kind: o,
		Val:  val,
	}
}

func NewPush(i int) Code {
	return newCode2(PUSH, i)
}

func NewPlus() Code {
	return newCode(PLUS)
}

func NewMult() Code {
	return newCode(MULT)
}
