package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + literals
	IDENT   = "IDENT" // add, foobar, x, y, ...
	INT     = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT 		 = "<"
	GT 		 = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	EQ	   = "=="
	NOT_EQ = "!="

	// Keywords
	FUNCTION = "FUNCTION"
	LET		 = "LET"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
	IF	     = "IF"
	ELSE	 = "ELSE"
	RETURN	 = "RETURN"
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

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENT
}
