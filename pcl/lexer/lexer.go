package lexer

import (
	"unicode"

	"github.com/adamdusty/precept/pcl/token"
)

type Lexer struct {
	source string
	name   string // Lexer name for debugging/errors
	column int    // Current column
	line   int    // Current line
	index  int    // Current index in string
	start  int    // Start of current token
}

func New(name string) *Lexer {
	return &Lexer{
		name: name,
	}
}

func (lex *Lexer) Tokenize(input string) []token.Token {
	lex.source = input
}

// func Tokenize(source io.Reader) {
// 	lex := lexer{
// 		source: *bufio.NewReader(source),
// 		tokens: make(chan token.Token, 2),
// 	}

// 	for c, _, err := lex.source.ReadRune(); err == nil; {

// 		if c == '\n' {
// 			lex.line += 1
// 			lex.column = 0
// 			continue
// 		}

// 		switch c {
// 		case '"':
// 			readStringLiteral(lex)
// 		case '+':
// 			lex.emit(token.Token{Type: token.Plus})
// 		case '-':
// 			lex.emit(token.Token{Type: token.Minus})
// 		case '*':
// 			lex.emit(token.Token{Type: token.Star})
// 		case '/':
// 			lex.emit(token.Token{Type: token.Slash})
// 		case '.':
// 			lex.emit(token.Token{Type: token.Dot})
// 		case ',':
// 			lex.emit(token.Token{Type: token.Comma})
// 		case '(':
// 			lex.emit(token.Token{Type: token.LeftParen})
// 		case ')':
// 			lex.emit(token.Token{Type: token.RightParen})
// 		case '{':
// 			lex.emit(token.Token{Type: token.LeftBrace})
// 		case '}':
// 			lex.emit(token.Token{Type: token.RightBrace})
// 		case '[':
// 			lex.emit(token.Token{Type: token.LeftBracket})
// 		case ']':
// 			lex.emit(token.Token{Type: token.RightBracket})
// 		case ';':
// 			lex.emit(token.Token{Type: token.SemiColon})
// 		case ':':
// 			lex.emit(token.Token{Type: token.Colon})
// 		case '=':
// 			if match(lex, '=') {
// 				lex.emit(token.Token{Type: token.EqualEq})
// 			} else {
// 				lex.emit(token.Token{Type: token.Equal})
// 			}
// 		case '!':
// 			if match(lex, '=') {
// 				lex.emit(token.Token{Type: token.BangEq})
// 			} else {
// 				lex.emit(token.Token{Type: token.Bang})
// 			}
// 		case '>':
// 			if match(lex, '=') {
// 				lex.emit(token.Token{Type: token.GreaterEq})
// 			} else {
// 				lex.emit(token.Token{Type: token.Greater})
// 			}
// 		case '<':
// 			if match(lex, '=') {
// 				lex.emit(token.Token{Type: token.LessEq})
// 			} else {
// 				lex.emit(token.Token{Type: token.Less})
// 			}
// 		default:
// 			if unicode.IsDigit(c) {
// 				readNumberLiteral(lex)
// 			} else if isIdentChar(c, true) {
// 				readIdentifer(lex)
// 			}
// 		}
// 	}
// }

// func match(lex *lexer, expected rune) bool {
// 	c, _, err := lex.source.ReadRune()
// 	if err == nil && c == expected {
// 		lex.column += 1
// 		return true
// 	}

// 	lex.source.UnreadRune()
// 	return false
// }

func (lex *lexer) emit(t token.Token) {
	lex.tokens <- t
}

func readNumberLiteral(lex lexer) token.Token { return token.Token{} }

func readStringLiteral(lex lexer) token.Token { return token.Token{} }

func readIdentifer(lex lexer) token.Token {

}

func isIdentChar(c rune, start bool) bool {
	if unicode.IsLetter(c) || c == '_' {
		return true
	}

	if !start && unicode.IsDigit(c) {
		return true
	}

	return false
}
