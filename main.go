package main

import (
	"fmt"
)

func main() {
	lexer := NewLexer("1 + 10 / 5 * 2")
	intpr := NewInterpreter(lexer)
	fmt.Println(intpr.Expr())
}

const (
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	MULTIP  = "MULTIP"
	EOF     = "EOF"
	DIV     = "DIV"
)

type Token struct {
	Type string
	Val  interface{}
}

type Interpreter struct {
	Lexer        *Lexer
	CurrentToken Token
}

func NewInterpreter(lexer *Lexer) *Interpreter {
	return &Interpreter{Lexer: lexer, CurrentToken: lexer.getNextToken()}
}

func (i *Interpreter) eat(_type string) {
	if i.CurrentToken.Type == _type {
		i.CurrentToken = i.Lexer.getNextToken()
	} else {
		panic("eat Type do not equal to current token type")
	}
}

func (i *Interpreter) factor() int {
	token := i.CurrentToken
	i.eat(INTEGER)
	return token.Val.(int)
}

func (i *Interpreter) Expr() int {

	result := NewCalcTree(i.factor())

	for i.CurrentToken.Type != EOF {
		if i.CurrentToken.Type == PLUS {
			i.eat(PLUS)
			result.addOperation(PLUS, i.factor())
		} else if i.CurrentToken.Type == MULTIP {
			i.eat(MULTIP)
			result.addOperation(MULTIP, i.factor())
		} else if i.CurrentToken.Type == MINUS {
			i.eat(MINUS)
			result.addOperation(MINUS, i.factor())
		} else if i.CurrentToken.Type == DIV {
			i.eat(DIV)
			result.addOperation(DIV, i.factor())
		}
	}
	return result.calc()
}
