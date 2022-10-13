package lexer_test

import (
	"testing"

	"github.com/adamdusty/precept/pcl/lexer"
	"github.com/adamdusty/precept/pcl/token"
)

const input = `
	a: i32;
`

// b: i32 = 42;
// c := 84;
// d: f32 = 42.5;
// e: f32 = .75;
// f := 3.14159
// g := .505

// add : (x: i32, y: i32) i32 = (x: i32, y: i32) i32 { return x + y; }

var tokens = []token.Token{
	{Type: token.Identifier, Lexeme: "a", Line: 1, Column: 0},
	{Type: token.Colon, Lexeme: ":", Line: 1, Column: 1},
	{Type: token.Identifier, Lexeme: "i32", Line: 1, Column: 3},
	{Type: token.Identifier, Lexeme: ";", Line: 1, Column: 6},
}

func TestLexer(t *testing.T) {
	lex := lexer.New("test")
	lex.Tokenize(input)
}
