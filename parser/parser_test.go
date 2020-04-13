package parser

import (
	"fmt"
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

	testIntegerLiteral(t, es.Expression, 8)
}

func TestPrefixExpression(t *testing.T) {
	input := "-23"

	l := lexer.New(input)
	p := New(l)
	es := p.Parse()
	checkParserErrors(t, p)

	pe, ok := es.Expression.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("pe not *ast.PrefixExpression. got=%T", es.Expression)
	}
	if pe.Operator != "-" {
		t.Fatalf("pe.Operator is not %s. got=%s", "-", pe.Operator)
	}
	if !testIntegerLiteral(t, pe.Right, 23) {
		return
	}
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	il, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", exp)
		return false
	}

	if il.Value != value {
		t.Errorf("il.Value not %d. got=%d", value, il.Value)
		return false
	}

	if il.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("il.TokenLiteral not %d. got=%s", value, il.TokenLiteral())
		return false
	}

	return true
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
