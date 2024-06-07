package main

import (
	"strconv"
	"unicode"
)

const (
	_ = iota
	PLUS
	MINUS
	DIV
	MUL
	LPAREN
	RPAREN
	INTEGER
	EOF
	BEGIN
	END
	DOT
	ASSIGN
	SEMI
	ID
)

var RESERVED_KEYWORDS = map[string]Token{
	"BEGIN": {BEGIN, "BEGIN"},
	"END":   {END, "END"},
}

type Token struct {
	_type int
	val   any
}

type Lexer struct {
	text        string
	pos         int
	currentChar rune
}

func NewLexer(text string) *Lexer {
	return &Lexer{text: text, currentChar: rune(text[0])}
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos >= len(l.text) {
		l.currentChar = 0
		return
	}
	l.currentChar = rune(l.text[l.pos])
}

func (l *Lexer) peek() rune {
	peek_pos := l.pos + 1
	if peek_pos >= len(l.text) {
		return 0
	}
	return rune(l.text[peek_pos])
}

func (l *Lexer) _id() Token {
	result := ""
	for unicode.IsDigit(l.currentChar) || unicode.IsLetter(l.currentChar) {
		result += string(l.currentChar)
		l.advance()
	}
	token, ok := RESERVED_KEYWORDS[result]
	if !ok {
		token = Token{ID, result}
	}
	return token
}

func (l *Lexer) integer() int {
	temp := ""
	for unicode.IsDigit(l.currentChar) {
		temp += string(l.currentChar)
		l.advance()
	}
	res, _ := strconv.Atoi(temp)
	return res
}

func (l *Lexer) getNextToken() Token {
	for l.currentChar != 0 {
		if l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\t' {
			for l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\t' {
				l.advance()
			}
			continue
		}
		if unicode.IsDigit(l.currentChar) {
			return Token{INTEGER, l.integer()}
		}
		if unicode.IsLetter(l.currentChar) {
			return l._id()
		}
		if l.currentChar == ':' && l.peek() == '=' {
			l.advance()
			l.advance()
			return Token{ASSIGN, ":="}
		}
		switch l.currentChar {
		case ';':
			l.advance()
			return Token{SEMI, "-"}
		case '.':
			l.advance()
			return Token{DOT, "."}
		case '-':
			l.advance()
			return Token{MINUS, "-"}
		case '+':
			l.advance()
			return Token{PLUS, "+"}
		case '*':
			l.advance()
			return Token{MUL, "*"}
		case '/':
			l.advance()
			return Token{DIV, "/"}
		case '(':
			l.advance()
			return Token{LPAREN, "("}
		case ')':
			l.advance()
			return Token{RPAREN, ")"}
		}
		panic("Unrecognized token" + string(l.currentChar) + "aas")
	}
	return Token{EOF, nil}
}
