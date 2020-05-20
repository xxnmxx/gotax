package lexer

import "gotax/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition > len(l.input) {
		l.ch = 0
	}
	l.position = l.readPosition
	l.ch = l.input[l.readPosition]
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '.':
		tok = newToken(token.DOT, l.ch)
	case '(':
		tok = newToken(token.RPAREN, l.ch)
	case ')':
		tok = newToken(token.LPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LSBLCK, l.ch)
	case ']':
		tok = newToken(token.RSBLCK, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '0':
		tok = newToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}
	readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Token: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position = l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' <= ch || 'A' <= ch && 'Z' <= ch || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '/r' {
		l.readChar()
	}
}
