package lexer

import (
	"hullo/token"
	"testing"
)

type NextTokenTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func runNextTokenTests(t *testing.T, input string, tests []NextTokenTest) {
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSpecialCharacters(t *testing.T) {
	input := `=+(){},;`
	tests := []NextTokenTest{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runNextTokenTests(t, input, tests)
}

func TestAssignment(t *testing.T) {
	input := `let x = 9;`
	tests := []NextTokenTest{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runNextTokenTests(t, input, tests)
}

func TestFunctionAdd(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
		let add = fn(x, y) {
			x + y;
		};
		let res = add(five, ten);
	`

	tests := []NextTokenTest{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "res"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runNextTokenTests(t, input, tests)
}

func TestAssignmentLargeNumber(t *testing.T) {
	input := `let x = 9999991111122222;`
	tests := []NextTokenTest{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "9999991111122222"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runNextTokenTests(t, input, tests)
}

func TestSymbols(t *testing.T) {
	input := `!-/*5<>`
	tests := []NextTokenTest{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLY, "*"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EOF, ""},
	}

	runNextTokenTests(t, input, tests)
}

func TestConditionals(t *testing.T) {
	input := `
		if (5 < 10) {
			return true;
		} else {
			return false;
		}
	`
	tests := []NextTokenTest{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	runNextTokenTests(t, input, tests)
}

func TestDoubleCharSymbols(t *testing.T) {
	input := `10 == 10; 20 != 10`
	tests := []NextTokenTest{
		{token.INT, "10"},
		{token.EQUALS, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "20"},
		{token.NOT_EQUALS, "!="},
		{token.INT, "10"},
		{token.EOF, ""},
	}

	runNextTokenTests(t, input, tests)
}
