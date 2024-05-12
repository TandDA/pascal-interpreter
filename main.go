package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	intpr := NewInterpreter(" 99 +      9 ")
	fmt.Println(intpr.Expr())
}

const (
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	EOF     = "EOF"
)

type Token struct {
	Type string
	Val  interface{}
}

type Interpreter struct {
	Text         string
	Pos          int
	CurrentToken Token
	CurrentChar  rune
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{Text: text, CurrentChar: rune(text[0])}
}

func (i *Interpreter) advance() {
	i.Pos += 1
	if i.Pos >= len(i.Text) {
		i.CurrentChar = 0
	} else {
		i.CurrentChar = rune(i.Text[i.Pos])
	}
}

func (i *Interpreter) skipWhitespace() {
	for i.CurrentChar == ' ' {
		i.advance()
	}
}

func (i *Interpreter) integer() int {
	num := ""
	for unicode.IsDigit(i.CurrentChar) {
		num += string(i.CurrentChar)
		i.advance()
	}
	ans, _ := strconv.Atoi(num)
	return ans
}

func (i *Interpreter) getNextToken() Token {
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
		panic("Invalid token")
	}
	return Token{EOF, nil}
}

func (i *Interpreter) eat(_type string) {
	if i.CurrentToken.Type == _type {
		i.CurrentToken = i.getNextToken()
	} else {
		panic("eat Type do not equal to current token type")
	}
}

func (i *Interpreter) Expr() int {
	i.CurrentToken = i.getNextToken()

	left := i.CurrentToken.Val.(int)
	i.eat(INTEGER)

	var plus bool
	if i.CurrentToken.Type == PLUS {
		plus = true
		i.eat(PLUS)
	} else {
		i.eat(MINUS)
	}

	right := i.CurrentToken.Val.(int)
	i.eat(INTEGER)

	if plus {
		return left + right
	} else {
		return left - right
	}
}
