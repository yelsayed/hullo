package lexer

import "hullo/token"

func newToken(tokType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokType, Literal: string(ch)}
}

// Returns whether a byte is a character that is valid for an identifier or a keyword.
// Notice how we don't consider dashes or spaces to be characters.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isInteger(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
