package lexer

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/adamdusty/precept/pcl/token"
)

var (
	id_start = []*unicode.RangeTable{
		unicode.L,
		unicode.Nl,
		unicode.Other_ID_Start,
	}
	id_continue = []*unicode.RangeTable{
		unicode.L,
		unicode.Nl,
		unicode.Other_ID_Start,
		unicode.Mn,
		unicode.Mc,
		unicode.Nd,
		unicode.Pc,
		unicode.Other_ID_Continue,
	}
	id_exclude = []*unicode.RangeTable{
		unicode.Pattern_Syntax,
		unicode.Pattern_White_Space,
	}
)

type lexerState struct {
	Source   []rune
	Length   int
	Line     int
	Column   int
	Position int
	Error    string
}

func (l *lexerState) advance(n int) *lexerState {
	next := new(lexerState)
	next.Source = l.Source
	next.Length = l.Length
	next.Line = l.Line
	next.Column = l.Column + n
	next.Position = l.Position + n
	next.Error = l.Error
	return next
}

func (lex *lexerState) match(c rune) bool {
	if !(lex.Position < lex.Length-1) {
		return false
	}

	next := lex.Source[lex.Position+1]

	return next == c
}

func (lex *lexerState) peek() rune {
	if !(lex.Position < lex.Length-1) {
		return utf8.RuneError
	}
	return lex.Source[lex.Position+1]
}

func isIdentStartChar(c rune) bool {
	return unicode.IsOneOf(id_start, c)
}

func isIdentContinueChar(c rune) bool {
	return unicode.IsOneOf(id_continue, c) && !unicode.IsOneOf(id_exclude, c)
}

func Tokenize(source string) ([]token.Token, error) {
	state := new(lexerState)
	state.Source = []rune(source)
	state.Line = 1
	state.Column = 0
	state.Position = 0
	state.Length = utf8.RuneCountInString(source)
	state.Error = ""
	tokens := []token.Token{}

	var tok token.Token

	for {
		tok, state = scanNextToken(state)
		tokens = append(tokens, tok)

		if tok.Type == token.Error {
			state.Error = tok.Lexeme
		}

		if state.Position > state.Length-1 {
			break
		}
	}

	if state.Error != "" {
		return tokens, errors.New(state.Error)
	}

	return tokens, nil
}

func scanNextToken(state *lexerState) (token.Token, *lexerState) {
	current := state.Source[state.Position]

	if unicode.IsSpace(current) {
		state = state.advance(1)
	}

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
		return readStringLiteral(state)
	default:
		if unicode.IsDigit(current) {
			return readNumberLiteral(state)
		}

		if isIdentStartChar(current) {
			return readIdentifier(state)
		}
		return token.Token{Type: token.Error, Lexeme: "Unimplemented."}, state.advance(1)
	}
}

func readStringLiteral(state *lexerState) (token.Token, *lexerState) {
	builder := strings.Builder{}

	for i := state.Position + 1; i <= state.Length-1; i += 1 {
		c := state.Source[i]
		if c == '"' {
			break
		}

		builder.WriteRune(c)
	}

	tok := token.Token{Type: token.String, Line: state.Line, Column: state.Column, Lexeme: builder.String()}

	// Advance state by 2 to account for the first '"' we didn't consume on entry
	return tok, state.advance(utf8.RuneCountInString(tok.Lexeme) + 2)
}

func readNumberLiteral(state *lexerState) (token.Token, *lexerState) {
	pos := state.Position
	builder := strings.Builder{}

	for pos < state.Length {
		if !unicode.IsDigit(state.Source[pos]) {
			break
		}

		builder.WriteRune(state.Source[pos])

		pos += 1
	}

	if pos >= state.Length {
		tok := token.Token{Type: token.Number, Line: state.Line, Column: state.Column, Lexeme: builder.String()}
		return tok, state.advance(utf8.RuneCountInString(tok.Lexeme) + 1)
	}

	if state.Source[pos] == '.' {
		// Write decimal
		builder.WriteRune(state.Source[pos])
		pos += 1

		for pos < state.Length {
			if !unicode.IsDigit(state.Source[pos]) {
				break
			}
			builder.WriteRune(state.Source[pos])

			pos += 1
		}
	}

	tok := token.Token{Type: token.Number, Line: state.Line, Column: state.Column, Lexeme: builder.String()}
	return tok, state.advance(pos)
}

func readIdentifier(state *lexerState) (token.Token, *lexerState) {

}
