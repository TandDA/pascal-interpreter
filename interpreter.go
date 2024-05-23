package main

type Interpreter struct {
	currentToken Token
	lexer        *Lexer
}

func NewInerpreter(lxr *Lexer) *Interpreter {
	return &Interpreter{lexer: lxr, currentToken: lxr.getNextToken()}
}

func (i *Interpreter) eat(_type string) {
	if i.currentToken._type == _type {
		i.currentToken = i.lexer.getNextToken()
		return
	}
	panic("Syntax error")
}

func (i *Interpreter) factor() int {
	val := i.currentToken.val
	i.eat(INTEGER)
	return val.(int)
}

func (i *Interpreter) term() int {
	result := i.factor()

	for i.currentToken._type == MUL || i.currentToken._type == DIV {
		if i.currentToken._type == MUL {
			i.eat(MUL)
			result *= i.factor()
		} else {
			i.eat(DIV)
			result /= i.factor()
		}
	}
	return result
}

func (i *Interpreter) expr() int {
	result := i.term()

	for i.currentToken._type == PLUS || i.currentToken._type == MINUS {
		if i.currentToken._type == PLUS {
			i.eat(PLUS)
			result += i.term()
		} else {
			i.eat(MINUS)
			result -= i.term()
		}
	}
	return result
}
