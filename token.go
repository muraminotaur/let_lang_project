package main

type TokenType string

type Token struct {
	Type string
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//5
	IDENT = "identifier"
	INT = "integer" // 1

	//operators 2
	MINUS = "minus"
	EQUAL = "="

	//delimiting
	COMMA = ","
	LPAREN = "("
	RPAREN = ")"
	

	//conditionals 4
	IF = "if"
	THEN = "then"
	ELSE = "else"
	ISZERO = "iszero" //3

	// keywords
	LET = "let" //6
	IN = "in"
	
)

var keywords = map[string]TokenType{
	"let": LET,
	"if": IF,
	"then": THEN,
	"else": ELSE,
	"iszero": ISZERO,
	"in": IN,
    "minus" : MINUS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
