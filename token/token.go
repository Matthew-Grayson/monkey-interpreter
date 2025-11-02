package token

// TokenType is a string for readability
// TODO: consider a different type to improve performance
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//	Identifiers & literals
	IDENT = "IDENT"
	INT   = "INT"

	//	Operators
	ASSIGN = "="
	PLUS   = "+"

	//	Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//	Token Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
