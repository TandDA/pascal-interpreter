package main

import (
	"strconv"
	"unicode"
)

type Lexer struct {
	Text        string
	Pos         int
	CurrentChar rune
}

func NewLexer(text string) *Lexer {
	return &Lexer{Text: text, CurrentChar: rune(text[0])}
}

func (i *Lexer) advance() {
	i.Pos += 1
	if i.Pos >= len(i.Text) {
		i.CurrentChar = 0
	} else {
		i.CurrentChar = rune(i.Text[i.Pos])
	}
}

func (i *Lexer) skipWhitespace() {
	for i.CurrentChar == ' ' {
		i.advance()
	}
}

func (i *Lexer) integer() int {
	num := ""
	for unicode.IsDigit(i.CurrentChar) {
		num += string(i.CurrentChar)
		i.advance()
	}
	ans, _ := strconv.Atoi(num)
	return ans
}

func (i *Lexer) getNextToken() Token {
	for i.CurrentChar != 0 {
		if i.CurrentChar == ' ' {
			i.skipWhitespace()
			continue
		}
		if unicode.IsDigit(i.CurrentChar) {
			return Token{INTEGER, i.integer()}
		}
		if i.CurrentChar == '+' {
			i.advance()
			return Token{PLUS, '+'}
		}
		if i.CurrentChar == '-' {
			i.advance()
			return Token{MINUS, '-'}
		}
		if i.CurrentChar == '*' {
			i.advance()
			return Token{MULTIP, '*'}
		}
		panic("Invalid token")
	}
	return Token{EOF, nil}
}
