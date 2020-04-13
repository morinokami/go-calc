package parser

import (
	"fmt"
	"strconv"

	"github.com/morinokami/go.calc/ast"
	"github.com/morinokami/go.calc/lexer"
	"github.com/morinokami/go.calc/token"
)

const (
	_ int = iota
	LOWEST
	PREFIX
)

type (
	prefixParseFn func() ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Parse() *ast.ExpressionStatement {
	es := p.parseStatement()
	p.nextToken()

	return es
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) noPrefixParserFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseStatement() *ast.ExpressionStatement {
	return p.parseExpressionStatement()
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	es := &ast.ExpressionStatement{Token: p.curToken}

	es.Expression = p.parseExpression(LOWEST)

	return es
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParserFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
