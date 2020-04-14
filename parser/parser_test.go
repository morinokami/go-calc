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

func TestInfixExpression(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		es := p.Parse()
		checkParserErrors(t, p)

		ie, ok := es.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("ie not *ast.InfixExpression. got=%T", es.Expression)
		}
		if !testIntegerLiteral(t, ie.Left, tt.leftValue) {
			return
		}
		if ie.Operator != tt.operator {
			t.Fatalf("ie.Operator is not %s. got=%s", tt.operator, ie.Operator)
		}
		if !testIntegerLiteral(t, ie.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-1 * 2",
			"((-1) * 2)",
		},
		{
			"--1",
			"(-(-1))",
		},
		{
			"1 + 2 + 3",
			"((1 + 2) + 3)",
		},
		{
			"1 + 2 - 3",
			"((1 + 2) - 3)",
		},
		{
			"1 * 2 * 3",
			"((1 * 2) * 3)",
		},
		{
			"1 * 2 / 3",
			"((1 * 2) / 3)",
		},
		{
			"1 + 2 / 3",
			"(1 + (2 / 3))",
		},
		{
			"1 + 2 * 3 + 4 / 5 - 6",
			"(((1 + (2 * 3)) + (4 / 5)) - 6)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		es := p.Parse()
		checkParserErrors(t, p)

		got := es.String()
		if got != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, got)
		}
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
