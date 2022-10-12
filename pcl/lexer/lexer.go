package lexer

import (
	"bufio"
	"io"
	"unicode"

	"github.com/adamdusty/precept/pcl/token"
)

type lexer struct {
	Source bufio.Reader
	Column int
	Line   int
}

func Tokenize(source io.Reader) []token.Token {
	lex := lexer{Source: *bufio.NewReader(source)}
	tokens := make([]token.Token, 16)

	var tok token.Token

	for c, _, err := lex.Source.ReadRune(); err == nil; {
		lex.Column += 1

		if c == '\n' {
			lex.Line += 1
			lex.Column = 0
			continue
		}

		switch c {
		case '+':
			tok = token.Token{Type: token.Plus, Lexeme: "+"}
		case '-':
			tok = token.Token{Type: token.Minus, Lexeme: "-"}
		case '*':
			tok = token.Token{Type: token.Star, Lexeme: "*"}
		case '/':
			tok = token.Token{Type: token.Slash, Lexeme: "/"}
		case '.':
			tok = token.Token{Type: token.Dot, Lexeme: "."}
		case ',':
			tok = token.Token{Type: token.Comma, Lexeme: ","}
		case '(':
			tok = token.Token{Type: token.LeftParen, Lexeme: "("}
		case ')':
			tok = token.Token{Type: token.RightParen, Lexeme: ")"}
		case '{':
			tok = token.Token{Type: token.LeftBrace, Lexeme: "{"}
		case '}':
			tok = token.Token{Type: token.RightBrace, Lexeme: "}"}
		case '[':
			tok = token.Token{Type: token.LeftBracket, Lexeme: "["}
		case ']':
			tok = token.Token{Type: token.RightBracket, Lexeme: "]"}
		case '=':
			if match(lex, '=') {
				tok = token.Token{Type: token.EqualEq, Lexeme: "=="}
			} else {
				tok = token.Token{Type: token.Equal, Lexeme: "="}
			}
		case '!':
			if match(lex, '=') {
				tok = token.Token{Type: token.BangEq, Lexeme: "!="}
			} else {
				tok = token.Token{Type: token.Bang, Lexeme: "!"}
			}
		case '>':
			if match(lex, '=') {
				tok = token.Token{Type: token.GreaterEq, Lexeme: "=="}
			} else {
				tok = token.Token{Type: token.Greater, Lexeme: "="}
			}
		case '<':
			if match(lex, '=') {
				tok = token.Token{Type: token.LessEq, Lexeme: "=="}
			} else {
				tok = token.Token{Type: token.Less, Lexeme: "="}
			}
		case '"':
			tok = readStringLiteral(lex)
		default:
			if unicode.IsDigit(c) {
				tok = readNumberLiteral(lex)
			} else if unicode.Is(unicode.Other_ID_Start, c) {
				tok = readIdentifer(lex)
			}
		}

		tok.Line = lex.Line
		tok.Column = lex.Column
		tokens = append(tokens, tok)
	}

	return tokens
}

func match(lex lexer, expected rune) bool {
	c, _, err := lex.Source.ReadRune()
	if err == nil && c == expected {
		lex.Column += 1
		return true
	}

	lex.Source.UnreadRune()
	return false
}

func readNumberLiteral(lex lexer) token.Token { return token.Token{} }

func readStringLiteral(lex lexer) token.Token { return token.Token{} }

func readIdentifer(lex lexer) token.Token { return token.Token{} }
