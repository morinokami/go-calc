package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Numbers
	INT = "INT"

	// Operators
	PLUS  = "+"
	MINUS = "-"
	STAR  = "*"
	SLASH = "/"

	LPAREN = "("
	RPAREN = ")"
)

type Token struct {
	Type    TokenType
	Literal string
}
