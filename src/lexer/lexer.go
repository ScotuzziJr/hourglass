package lexer

import (
	"hourglass.com/hourglass/src/token"
)

type Lexer struct {
	input        string // source code
	position     int  	// current position within input
	nextPosition int  	// next position in input (this is important to implement look ahead)
	char         byte 	// current char we're looking at in input
}

// create a lexer struct saving the input stream and pointing to the first char of it
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.char = 0 // 0 represents EOF
	} else {
		l.char = l.input[l.nextPosition]
	}

	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.char {
	case '=':
		if l.lookAhead() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(char) + string(l.char)}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '*':
		tok = newToken(token.MUL, l.char)
	case '/':
		tok = newToken(token.DIV, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case '!':
		if l.lookAhead() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(char) + string(l.char)}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUM
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
			l.readChar()
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) lookAhead() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}
