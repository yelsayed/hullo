package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENT      = "IDENT" // variables
	INT        = "INT"   // integers
	ASSIGN     = "="
	PLUS       = "+"
	COMMA      = ","
	SEMICOLON  = ";"
	LPAREN     = "("
	RPAREN     = ")"
	LBRACE     = "{"
	RBRACE     = "}"
	FUNCTION   = "FUNCTION"
	LET        = "LET"
	BANG       = "!"
	MINUS      = "-"
	DIVIDE     = "/"
	MULTIPLY   = "*"
	GT         = ">"
	LT         = "<"
	IF         = "if"
	ELSE       = "else"
	RETURN     = "return"
	TRUE       = "true"
	FALSE      = "false"
	EQUALS     = "=="
	NOT_EQUALS = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupTokenType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	// If it's not in the keywords that we allow
	return IDENT
}
