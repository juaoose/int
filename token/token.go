package token

// These are the tokens the lexer will output
// after going through the source code

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// We're using string for simplicity and convenience
// using int or byte could be more performant
const (
	ILLEGAL = "ILLEGAL" // A token we don't know about
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDEN"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
