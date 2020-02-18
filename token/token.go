package token

// TokenType is string
type TokenType string

// Token consists of Type and Literal
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL means unknown
	ILLEGAL = "ILLEGAL"
	// EOF means end of file
	EOF = "EOF"

	// IDENT is identifier
	IDENT = "IDENT"
	// INT is integer
	INT = "INT"

	// ASSIGN is =
	ASSIGN = "="
	// PLUS is +
	PLUS = "+"

	// COMMA is ,
	COMMA = ","
	// SEMICOLON is ;
	SEMICOLON = ";"

	// LPAREN is (
	LPAREN = "("
	// RPAREN is )
	RPAREN = ")"
	// LBRACE is {
	LBRACE = "{"
	// RBRACE is }
	RBRACE = "}"

	// FUNCTION is func
	FUNCTION = "FUNCTION"
	// LET is LET
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
