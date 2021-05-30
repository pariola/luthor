package token

/* Grammar
SIGN = + | - | / | *
<number> SIGN <number>
variable = <number> SIGN <number>
*/

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	// SIGNS
	EQ  // =
	ADD // +
	SUB // -
	MUL // *
	QUO // /

	NUMBER // 124.0, 145.90

	VARIABLE // a, b, interest
)

var tokenNames = map[Token]string{
	WS:       "WS",
	EOF:      "EOF",
	EQ:       "EQ",
	ADD:      "ADD",
	SUB:      "SUB",
	MUL:      "MUL",
	QUO:      "QUO",
	NUMBER:   "NUMBER",
	VARIABLE: "VARIABLE",
	ILLEGAL:  "ILLEGAL",
}

func (t Token) String() string {
	return tokenNames[t]
}
