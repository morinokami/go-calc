package parser

import (
	"testing"

	"github.com/morinokami/go.calc/ast"
	"github.com/morinokami/go.calc/lexer"
)

func TestIntegerExpression(t *testing.T) {
	input := "8"

	l := lexer.New(input)
	p := New(l)
	es := p.Parse()
	checkParserErrors(t, p)

	exp, ok := es.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", es.Expression)
	}
	if exp.Value != 8 {
		t.Errorf("exp.Value not %d. got=%d", 8, exp.Value)
	}
	if exp.TokenLiteral() != "8" {
		t.Errorf("exp.TokenLiteral not %s. got=%s", "8", exp.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
