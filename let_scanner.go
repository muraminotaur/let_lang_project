package main

type Scanner struct{
	input string
	position int
	readPosition int
	ch byte
}

func New(input string) *Scanner{
	s := &Scanner{input: input}
	s.readChar()
	return s
}

// readChar()
// reads ahead and increments s.position and s.readPosition
func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input){
		s.ch = 0
	} else {
		s.ch = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition += 1
}

// peekChar()
// reads ahead but does not increment position
func (s *Scanner) peekChar() byte {
	if s.readPosition >= len(s.input){
		return 0
	} else {
		return s.input[s.readPosition]
	}
}

//NextToken()
func (s *Scanner) NextToken() token.Token { 
	var tok token.Token

	s.skipWhitespace()

	switch s.ch {
		case '=':
			tok = newToken(token.EQUAL, s.ch)
		case 'minus':
			tok = newToken(token.MINUS, s.ch)
		case ',':
			tok = newToken(token.COMMA, s.ch)
		case '(':
			tok = newToken(token.LPAREN, s.ch)
		case ')':
			tok = newToken(token.RPAREN, s.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(s.ch) {
				tok.Literal = s.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok
			} else if (isDigit(s.ch)){
				tok.Type = token.INT
				tok.Literal = s.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, s.ch)
			}
	}
	s.readChar()
	return tok
}

func (s *Scanner) readNumber() string {
	position := s.position
	for isDigit(s.ch){
		s.readChar()
	}
	return s.input[position:s.position]
}

func isDigit(ch byte) bool{
	return '0' <= ch && ch <= '9'
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (s *Scanner) readIdentifier() string {
	position := s.position
	for isLetter(s.ch){
		s.readChar()
	}
	return s.input[position:s.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

