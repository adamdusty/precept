package token

type TokenType string

type Token struct {
	Type   TokenType
	Lexeme string
	Column int
	Line   int
}

const (
	// Single character tokens
	Plus         = "+"
	Minus        = "-"
	Star         = "*"
	Slash        = "/"
	Dot          = "."
	Comma        = ","
	LeftParen    = "("
	RightParen   = ")"
	LeftBrace    = "{"
	RightBrace   = "}"
	LeftBracket  = "["
	RightBracket = "]"
	QuestionMark = "?"

	// Single or double character tokens
	Equal     = "="
	EqualEq   = "=="
	Bang      = "!"
	BangEq    = "!="
	Greater   = ">"
	GreaterEq = ">="
	Less      = "<"
	LessEq    = "<="

	// Multicharacter tokens
	String     = "string"
	Number     = "number"
	Identifier = "identifier"

	// Keywords
)
