// Token definition (not ipsis litteris)
// Louden: meaningful units collected from input stream

// My own definition: by reading an input stream (source code), one identify lexemes (sequence of characters)
//					and classify them into pre-defined categories, assigning the category itself and a number
//					to identify the category. The combination of category and value (number) is called 'token'

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	ID  = "ID"  // foo, bar, first_name...
	NUM = "NUM" // 42, 3.14, ...

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MUL    = "*"
	DIV    = "/"
	BANG   = "!"

	LT     = "<"
	GT     = ">"
	EQ 	   = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNC   = "FUNC"
	VAR    = "VAR"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	RETURN = "RETURN"
	IF 	   = "IF"
	ELSE   = "ELSE"
)

var Keywords = map[string]TokenType{
	"func": FUNC,
	"var":  VAR,
	"return": RETURN,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
}

// check weather an identifier is a keyword or not
// this is important to make a distinction between variable/function names and keywords,
//	considering they are all strings 
func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}

	return ID
}
