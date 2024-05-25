package main

import (
	"strconv"
	"unicode"
)

const (
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	DIV     = "DIV"
	MUL     = "MUL"
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
	INTEGER = "INTEGER"
	EOF     = "EOF"
)

type Token struct {
	_type string
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

func (l *Lexer) integer() int {
	temp := ""
	for unicode.IsDigit(l.currentChar) {
		temp+= string(l.currentChar)
		l.advance()
	}
	res, _ := strconv.Atoi(temp)
	return res
}

func (l *Lexer) getNextToken() Token {
	for l.currentChar != 0 {
		if l.currentChar == ' ' {
			for l.currentChar == ' ' {
				l.advance()
			}
			continue
		}
		if unicode.IsDigit(l.currentChar) {
			return Token{INTEGER, l.integer()}
		}
		switch l.currentChar {
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
		panic("Unrecognized token")
	}
	return Token{EOF, nil}
}