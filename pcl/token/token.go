package token

type TokenType int16

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}

const (
	// Special
	Error TokenType = -1
	EOF   TokenType = 0

	// Single character tokens
	Plus TokenType = iota
	Minus
	Star
	Slash
	Dot
	Comma
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	QuestionMark
	SemiColon
	Colon

	// Single or double character tokens
	Equal
	EqualEq
	Bang
	BangEq
	Greater
	GreaterEq
	Less
	LessEq

	// Multicharacter tokens
	String
	Number
	Identifier

	// Keywords
	Return
)
