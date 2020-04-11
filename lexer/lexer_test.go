package lexer

import (
	"testing"

	"github.com/morinokami/go.calc/token"
)

func TestNextToken(t *testing.T) {
	input := `(1 - 5) * ((2 + 3) / 4)`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.MINUS, "-"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.TIMES, "*"},
		{token.LPAREN, "("},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.DIVIDE, "/"},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
