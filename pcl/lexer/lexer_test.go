package lexer_test

import (
	"testing"

	"github.com/adamdusty/precept/pcl/lexer"
	"github.com/adamdusty/precept/pcl/token"
)

// a: i32;
// b: i32 = 42;
// c := 84;
// d: f32 = 42.5;
// e: f32 = .75;
// f := 3.14159
// g := .505

// add : (x: i32, y: i32) i32 = (x: i32, y: i32) i32 { return x + y; }

func TestSingleCharacter(t *testing.T) {
	input := `({[]}),.:;?/*+-`
	expectedTokens := []token.Token{
		{Type: token.LeftParen, Line: 1, Column: 0},
		{Type: token.LeftBrace, Line: 1, Column: 1},
		{Type: token.LeftBracket, Line: 1, Column: 2},
		{Type: token.RightBracket, Line: 1, Column: 3},
		{Type: token.RightBrace, Line: 1, Column: 4},
		{Type: token.RightParen, Line: 1, Column: 5},
		{Type: token.Comma, Line: 1, Column: 6},
		{Type: token.Dot, Line: 1, Column: 7},
		{Type: token.Colon, Line: 1, Column: 8},
		{Type: token.SemiColon, Line: 1, Column: 9},
		{Type: token.QuestionMark, Line: 1, Column: 10},
		{Type: token.Slash, Line: 1, Column: 11},
		{Type: token.Star, Line: 1, Column: 12},
		{Type: token.Plus, Line: 1, Column: 13},
		{Type: token.Minus, Line: 1, Column: 14},
	}

	actualTokens, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actualTokens) != len(expectedTokens) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actualTokens))
	}

	for i := 0; i < len(actualTokens); i++ {
		if actualTokens[i] != expectedTokens[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actualTokens[i], expectedTokens[i])
		}
	}
}

func TestEqual(t *testing.T) {
	input := `=`
	expected := []token.Token{{Type: token.Equal, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestEqualEq(t *testing.T) {
	input := `==`
	expected := []token.Token{{Type: token.EqualEq, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestBang(t *testing.T) {
	input := `!`
	expected := []token.Token{{Type: token.Bang, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestBangEq(t *testing.T) {
	input := `!=`
	expected := []token.Token{{Type: token.BangEq, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestGreater(t *testing.T) {
	input := `>`
	expected := []token.Token{{Type: token.Greater, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestGreaterEq(t *testing.T) {
	input := `>=`
	expected := []token.Token{{Type: token.GreaterEq, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}
func TestLess(t *testing.T) {
	input := `<`
	expected := []token.Token{{Type: token.Less, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestLessEq(t *testing.T) {
	input := `<=`
	expected := []token.Token{{Type: token.LessEq, Line: 1, Column: 0}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello world"`
	expected := []token.Token{{Type: token.String, Line: 1, Column: 0, Lexeme: "Hello world"}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestIntegerNumberLiteral(t *testing.T) {
	input := "1234"
	expected := []token.Token{{Type: token.Number, Line: 1, Column: 0, Lexeme: "1234"}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}

func TestDecimalNumberLiteral(t *testing.T) {
	input := "1234.5678"
	expected := []token.Token{{Type: token.Number, Line: 1, Column: 0, Lexeme: "1234.5678"}}

	actual, err := lexer.Tokenize(input)

	if err != nil {
		t.Errorf("Error lexing: %s", err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length of actual tokens not equal to expected: %d", len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Actual token not equal to expected token: %v != %v", actual[i], expected[i])
		}
	}
}
