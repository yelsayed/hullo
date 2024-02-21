package lexer

import "hullo/token"

type Lexer struct {
	input        string
	position     int  // points to current char that we're processing
	readPosition int  // one after current char
	ch           byte // the value of the current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.peekNext('=') {
			tok.Type = token.EQUALS
			tok.Literal = "=="
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '!':
		if l.peekNext('=') {
			tok.Type = token.NOT_EQUALS
			tok.Literal = "!="
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupTokenType(tok.Literal)
			return tok
		} else if isInteger(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()

	return tok
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekNext(ch byte) bool {
	if l.readPosition >= len(l.input) {
		return false
	} else if ch == l.input[l.readPosition] {
		l.readChar() // Advance the pointer
		return true
	}

	return false
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isInteger(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}
