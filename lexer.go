package main

import (
	"strconv"
	"unicode"
)

const (
	MINUS   = "MINUS"
	PLUS    = "PLUS"
	DIV     = "DIV"
	MUL     = "MUL"
	INTEGER = "INTEGER"
	EOF     = "EOF"
)

type Token struct {
	_type string
	val   any
}

type Lexer struct {
	pos         int
	currentChar rune
	text        string
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
		temp += string(l.currentChar)
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
		if l.currentChar == '-' {
			l.advance()
			return Token{MINUS, '-'}
		}
		if l.currentChar == '+' {
			l.advance()
			return Token{PLUS, '+'}
		}
		if l.currentChar == '*' {
			l.advance()
			return Token{MUL, '*'}
		}
		if l.currentChar == '/' {
			l.advance()
			return Token{DIV, '/'}
		}
		panic("Token recognize error")
	}
	return Token{EOF, nil}
}
