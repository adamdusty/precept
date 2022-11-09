package lexer

import (
	"errors"
	"unicode/utf8"

	"github.com/adamdusty/precept/pcl/token"
)

type lexerState struct {
	Source   []rune
	Length   int
	Line     int
	Column   int
	Position int
	Error    string
}

func (l *lexerState) advance(n int) lexerState {
	source := l.Source
	length := l.Length
	line := l.Line
	col := l.Column + n
	pos := l.Position + n
	err := l.Error
	return lexerState{Source: source, Line: line, Column: col, Position: pos, Length: length, Error: err}
}

func Tokenize(source string) ([]token.Token, error) {
	state := lexerState{Source: []rune(source), Line: 1, Column: 0, Position: 0, Length: utf8.RuneCountInString(source) - 1, Error: ""}
	tokens := []token.Token{}

	var tok token.Token

	for {
		tok, state = scanNextToken(&state)
		tokens = append(tokens, tok)

		if tok.Type == token.Error {
			state.Error = tok.Lexeme
		}

		if state.Position > state.Length {
			break
		}
	}

	if state.Error != "" {
		return tokens, errors.New(state.Error)
	}

	return tokens, nil
}

func scanNextToken(state *lexerState) (token.Token, lexerState) {
	current := state.Source[state.Position]

	switch current {
	case '(':
		return token.Token{Type: token.LeftParen, Line: state.Line, Column: state.Column}, state.advance(1)
	case ')':
		return token.Token{Type: token.RightParen, Line: state.Line, Column: state.Column}, state.advance(1)
	case '{':
		return token.Token{Type: token.LeftBrace, Line: state.Line, Column: state.Column}, state.advance(1)
	case '}':
		return token.Token{Type: token.RightBrace, Line: state.Line, Column: state.Column}, state.advance(1)
	case '[':
		return token.Token{Type: token.LeftBracket, Line: state.Line, Column: state.Column}, state.advance(1)
	case ']':
		return token.Token{Type: token.RightBracket, Line: state.Line, Column: state.Column}, state.advance(1)
	case '.':
		return token.Token{Type: token.Dot, Line: state.Line, Column: state.Column}, state.advance(1)
	case ',':
		return token.Token{Type: token.Comma, Line: state.Line, Column: state.Column}, state.advance(1)
	case ':':
		return token.Token{Type: token.Colon, Line: state.Line, Column: state.Column}, state.advance(1)
	case ';':
		return token.Token{Type: token.SemiColon, Line: state.Line, Column: state.Column}, state.advance(1)
	case '?':
		return token.Token{Type: token.QuestionMark, Line: state.Line, Column: state.Column}, state.advance(1)
	case '/':
		return token.Token{Type: token.Slash, Line: state.Line, Column: state.Column}, state.advance(1)
	case '*':
		return token.Token{Type: token.Star, Line: state.Line, Column: state.Column}, state.advance(1)
	case '+':
		return token.Token{Type: token.Plus, Line: state.Line, Column: state.Column}, state.advance(1)
	case '-':
		return token.Token{Type: token.Minus, Line: state.Line, Column: state.Column}, state.advance(1)
	case '=':
		if state.match('=') {
			return token.Token{Type: token.EqualEq, Line: state.Line, Column: state.Column}, state.advance(2)
		} else {
			return token.Token{Type: token.Equal, Line: state.Line, Column: state.Column}, state.advance(1)
		}
	case '!':
		if state.match('=') {
			return token.Token{Type: token.BangEq, Line: state.Line, Column: state.Column}, state.advance(2)
		} else {
			return token.Token{Type: token.Bang, Line: state.Line, Column: state.Column}, state.advance(1)
		}
	case '>':
		if state.match('=') {
			return token.Token{Type: token.GreaterEq, Line: state.Line, Column: state.Column}, state.advance(2)
		} else {
			return token.Token{Type: token.Greater, Line: state.Line, Column: state.Column}, state.advance(1)
		}
	case '<':
		if state.match('=') {
			return token.Token{Type: token.LessEq, Line: state.Line, Column: state.Column}, state.advance(2)
		} else {
			return token.Token{Type: token.Less, Line: state.Line, Column: state.Column}, state.advance(1)
		}
	case '"':
		return token.Token{Type: token.Error, Lexeme: "Unimplemented."}, state.advance(1)
	default:
		return token.Token{Type: token.Error, Lexeme: "Unimplemented."}, state.advance(1)
	}
}

func (lex *lexerState) match(c rune) bool {
	if !(lex.Position < lex.Length) {
		return false
	}

	next := lex.Source[lex.Position+1]

	return next == c
}

func (lex *lexerState) peek() rune {
	if !(lex.Position < lex.Length) {
		return utf8.RuneError
	}
	return lex.Source[lex.Position+1]
}

// func Tokenize(rd io.Reader) ([]token.Token, error) {
// 	lex := lexer{source: *bufio.NewReader(rd), line: 1, column: 0}

// 	var tokens []token.Token

// 	for {
// 		tok := scanNextToken(&lex)
// 		if tok.Type == token.EOF {
// 			break
// 		}

// 		if tok.Type == token.Error {
// 			log.Printf("ERROR: %s", tok.Lexeme)
// 			return []token.Token{}, errors.New(tok.Lexeme)
// 		}

// 		tokens = append(tokens, tok)
// 	}

// 	return tokens, nil
// }

// type lexer struct {
// 	source bufio.Reader
// 	line   int
// 	column int
// }

// func scanNextToken(lex *lexer) token.Token {
// 	for c, _, err := lex.source.ReadRune(); err != io.EOF; {
// 		if c == '\n' {
// 			lex.line += 1
// 			continue
// 		}

// 		if unicode.IsSpace(c) {
// 			continue
// 		}

// 		switch c {
// 		case '(':
// 			tok := token.Token{Type: token.LeftParen, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case ')':
// 			tok := token.Token{Type: token.RightParen, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '{':
// 			tok := token.Token{Type: token.LeftBrace, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '}':
// 			tok := token.Token{Type: token.RightBrace, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '[':
// 			tok := token.Token{Type: token.LeftBracket, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case ']':
// 			tok := token.Token{Type: token.RightBracket, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '.':
// 			tok := token.Token{Type: token.Dot, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case ',':
// 			tok := token.Token{Type: token.Comma, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case ':':
// 			tok := token.Token{Type: token.Colon, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case ';':
// 			tok := token.Token{Type: token.SemiColon, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '?':
// 			tok := token.Token{Type: token.QuestionMark, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '/':
// 			tok := token.Token{Type: token.Slash, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '*':
// 			tok := token.Token{Type: token.Star, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '+':
// 			tok := token.Token{Type: token.Plus, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '-':
// 			tok := token.Token{Type: token.Minus, Line: lex.line, Column: lex.column}
// 			lex.column += 1
// 			return tok
// 		case '=':
// 			if lex.match('=') {
// 				tok := token.Token{Type: token.EqualEq, Line: lex.line, Column: lex.column}
// 				lex.column += 2
// 				return tok
// 			} else {
// 				tok := token.Token{Type: token.Equal, Line: lex.line, Column: lex.column}
// 				lex.column += 1
// 				return tok
// 			}
// 		case '!':
// 			if lex.match('=') {
// 				tok := token.Token{Type: token.BangEq, Line: lex.line, Column: lex.column}
// 				lex.column += 2
// 				return tok
// 			} else {
// 				tok := token.Token{Type: token.Bang, Line: lex.line, Column: lex.column}
// 				lex.column += 1
// 				return tok
// 			}
// 		case '>':
// 			if lex.match('=') {
// 				tok := token.Token{Type: token.GreaterEq, Line: lex.line, Column: lex.column}
// 				lex.column += 2
// 				return tok
// 			} else {
// 				tok := token.Token{Type: token.Greater, Line: lex.line, Column: lex.column}
// 				lex.column += 1
// 				return tok
// 			}
// 		case '<':
// 			if lex.match('=') {
// 				tok := token.Token{Type: token.LessEq, Line: lex.line, Column: lex.column}
// 				lex.column += 2
// 				return tok
// 			} else {
// 				tok := token.Token{Type: token.Less, Line: lex.line, Column: lex.column}
// 				lex.column += 1
// 				return tok
// 			}
// 		case '"':
// 			tok := lex.readStringLiteral()
// 			lex.column += utf8.RuneCountInString(tok.Lexeme)
// 			return tok
// 		default:
// 			if unicode.IsDigit(c) {
// 				tok := lex.readNumberLiteral()
// 				lex.column += utf8.RuneCountInString(tok.Lexeme)
// 				return tok
// 			}
// 			return token.Token{Type: token.Error, Lexeme: "Unrecognized character."}
// 		}
// 	}

// 	return token.Token{Type: token.EOF}
// }

// func (lex *lexer) match(c rune) bool {
// 	next, _, err := lex.source.ReadRune()
// 	if err != nil {
// 		return false
// 	}

// 	if next == c {
// 		return true
// 	}

// 	lex.source.UnreadRune()
// 	return false
// }

// func (lex *lexer) peek() rune {
// 	c, _, err := lex.source.ReadRune()

// 	if err != nil {
// 		return unicode.ReplacementChar
// 	}

// 	lex.source.UnreadRune()
// 	return c
// }

// func (lex *lexer) advance() rune {
// 	c, _, err := lex.source.ReadRune()

// 	if err != nil {
// 		return unicode.ReplacementChar
// 	}

// 	return c
// }

// func (lex *lexer) readStringLiteral() token.Token {
// 	var builder strings.Builder

// 	for {
// 		c, _, err := lex.source.ReadRune()
// 		if err != nil {
// 			if err == io.EOF {
// 				return token.Token{Type: token.Error, Lexeme: "Expected closing \""}
// 			}

// 			return token.Token{Type: token.Error, Lexeme: "Error while scanning string literal"}
// 		}

// 		if c == '"' {
// 			break
// 		}

// 		builder.WriteRune(c)
// 	}

// 	tok := token.Token{Type: token.String, Line: lex.line, Column: lex.column, Lexeme: builder.String()}

// 	return tok
// }

// func (lex *lexer) readNumberLiteral() token.Token {
// 	lex.source.UnreadRune()
// 	var builder strings.Builder

// 	for {
// 		c, _, err := lex.source.ReadRune()
// 		if err != nil {
// 			if err == io.EOF {
// 				return token.Token{Type: token.Number, Line: lex.line, Column: lex.column, Lexeme: builder.String()}
// 			}

// 			return token.Token{Type: token.Error, Lexeme: "Error while scanning string literal", Line: lex.line, Column: lex.column}
// 		}

// 		if unicode.IsDigit(c) {
// 			builder.WriteRune(c)
// 		} else {
// 			lex.source.UnreadRune()
// 			break
// 		}
// 	}

// 	if c := lex.peek(); c == '.' {
// 		builder.WriteRune(lex.advance())

// 		for {
// 			c, _, err := lex.source.ReadRune()
// 			if err != nil {
// 				if err == io.EOF {
// 					return token.Token{Type: token.Number, Line: lex.line, Column: lex.column, Lexeme: builder.String()}
// 				}

// 				return token.Token{Type: token.Error, Lexeme: "Error while scanning string literal", Line: lex.line, Column: lex.column}
// 			}

// 			if unicode.IsDigit(c) {
// 				builder.WriteRune(c)
// 			} else {
// 				lex.source.UnreadRune()
// 				break
// 			}
// 		}
// 	}

// 	return token.Token{Type: token.Number, Line: lex.line, Column: lex.column, Lexeme: builder.String()}
// }
