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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	// If not a keyword, then its an identifier
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
