package lexer

import (
	"testing"

	"hourglass.com/hourglass/src/token"
)

func TestNextToken(t *testing.T) {
	input := `var numOne = 7;
	var numTwo = 6;
	
	var add = func(x, y) {
		x + y;
	}

	var res = add(numOne, numTwo);

	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	42 == 42;
	21 != 7;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.ID, "numOne"},
		{token.ASSIGN, "="},
		{token.NUM, "7"},
		{token.SEMICOLON, ";"},

		{token.VAR, "var"},
		{token.ID, "numTwo"},
		{token.ASSIGN, "="},
		{token.NUM, "6"},
		{token.SEMICOLON, ";"},

		{token.VAR, "var"},
		{token.ID, "add"},
		{token.ASSIGN, "="},
		{token.FUNC, "func"},
		{token.LPAREN, "("},
		{token.ID, "x"},
		{token.COMMA, ","},
		{token.ID, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.ID, "x"},
		{token.PLUS, "+"},
		{token.ID, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.VAR, "var"},
		{token.ID, "res"},
		{token.ASSIGN, "="},
		{token.ID, "add"},
		{token.LPAREN, "("},
		{token.ID, "numOne"},
		{token.COMMA, ","},
		{token.ID, "numTwo"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.DIV, "/"},
		{token.MUL, "*"},
		{token.NUM, "5"},
		{token.SEMICOLON, ";"},

		{token.NUM, "5"},
		{token.LT, "<"},
		{token.NUM, "10"},
		{token.GT, ">"},
		{token.NUM, "5"},
		{token.SEMICOLON, ";"},
		
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUM, "5"},
		{token.LT, "<"},
		{token.NUM, "10"},
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

		{token.NUM, "42"},
		{token.EQ, "=="},
		{token.NUM, "42"},
		{token.SEMICOLON, ";"},

		{token.NUM, "21"},
		{token.NOT_EQ, "!="},
		{token.NUM, "7"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected = %q, got = %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
