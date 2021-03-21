package lexer

import "geomys/token"

type Lexer struct {
	input		  string
	position	  int // current position in input (points to current chart)
	readPosition  int // current reading position in input (after current chart)
	currentChar	  byte // current char under examination
}

func New(input string) *Lexer {
	lexPtr := &Lexer{input: input}
	lexPtr.readChar()
	return lexPtr
}

func (lexPtr *Lexer) readChar() {
	if lexPtr.readPosition >= len(lexPtr.input) {
		lexPtr.currentChar = 0
	} else {
		lexPtr.currentChar = lexPtr.input[lexPtr.readPosition]
	}
	lexPtr.position = lexPtr.readPosition
	lexPtr.readPosition += 1
}

func (lexPtr *Lexer) NextToken() token.Token {
	var nextTok token.Token

	lexPtr.skipWhitespace()

	switch lexPtr.currentChar {
	case '=':
		if lexPtr.peekChar() == '=' {
			currentChar := lexPtr.currentChar
			lexPtr.readChar()
			nextTok = token.Token{Type: token.EQ, Literal: string(currentChar) + string(lexPtr.currentChar)}
		} else {
			nextTok = newToken(token.ASSIGN, lexPtr.currentChar)
		}
	case '+':
		nextTok = newToken(token.PLUS, lexPtr.currentChar)
	case '-':
		nextTok = newToken(token.MINUS, lexPtr.currentChar)
	case '!':
		if lexPtr.peekChar() == '=' {
			currentChar := lexPtr.currentChar
			lexPtr.readChar()
			nextTok = token.Token{Type: token.NOT_EQ, Literal: string(currentChar) + string(lexPtr.currentChar)}
		} else {
			nextTok = newToken(token.BANG, lexPtr.currentChar)
		}
	case '/':
		nextTok = newToken(token.SLASH, lexPtr.currentChar)
	case '*':
		nextTok = newToken(token.ASTERISK, lexPtr.currentChar)
	case '<':
		nextTok = newToken(token.LT, lexPtr.currentChar)
	case '>':
		nextTok = newToken(token.GT, lexPtr.currentChar)
	case ';':
		nextTok = newToken(token.SEMICOLON, lexPtr.currentChar)
	case ',':
		nextTok = newToken(token.COMMA, lexPtr.currentChar)
	case '(':
		nextTok = newToken(token.LPAREN, lexPtr.currentChar)
	case ')':
		nextTok = newToken(token.RPAREN, lexPtr.currentChar)
	case '{':
		nextTok = newToken(token.LBRACE, lexPtr.currentChar)
	case '}':
		nextTok = newToken(token.RBRACE, lexPtr.currentChar)
	case 0:
		nextTok.Literal = ""
		nextTok.Type = token.EOF
	default:
		if isLetter(lexPtr.currentChar) {
			nextTok.Literal = lexPtr.readIdentifier()
			nextTok.Type = token.LookupIdentifier(nextTok.Literal)
			return nextTok
		} else if isDigit(lexPtr.currentChar) {
			nextTok.Type = token.INT
			nextTok.Literal = lexPtr.readNumber()
			return nextTok
		} else {
			nextTok = newToken(token.ILLEGAL, lexPtr.currentChar)
		}
	}
	lexPtr.readChar()
	return nextTok
}

func (lexPtr *Lexer) readIdentifier() string {
	position := lexPtr.position
	for isLetter(lexPtr.currentChar) {
		lexPtr.readChar()
	}
	return lexPtr.input[position:lexPtr.position]
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func newToken(tokenType token.TokenType, currentChar byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(currentChar)}
}

func (lexPtr *Lexer) skipWhitespace() {
	for lexPtr.currentChar == ' ' || lexPtr.currentChar == '\t' || lexPtr.currentChar == '\n' || lexPtr.currentChar == '\r' {
		lexPtr.readChar()
	}
}

func (lexPtr *Lexer) readNumber() string {
	position := lexPtr.position
	for isDigit(lexPtr.currentChar) {
		lexPtr.readChar()
	}
	return lexPtr.input[position:lexPtr.position]
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexPtr *Lexer) peekChar() byte {
	if lexPtr.readPosition >= len(lexPtr.input) {
		return 0
	} else {
		return lexPtr.input[lexPtr.readPosition]
	}
}
