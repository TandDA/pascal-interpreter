package main

type Interpreter struct {
	l            *Lexer
	currentToken Token
}

func NewInterpreter(lxr *Lexer) *Interpreter {
	return &Interpreter{l: lxr, currentToken: lxr.getNextToken()}
}

func (i *Interpreter) eat(_type string) {
	if _type == i.currentToken._type {
		i.currentToken = i.l.getNextToken()
		return
	}
	panic("Syntax error")
}

func (i *Interpreter) factor() int {
	if i.currentToken._type == INTEGER {
		val := i.currentToken.val
		i.eat(INTEGER)
		return val.(int)
	}
	if i.currentToken._type == LPAREN {
		i.eat(LPAREN)
		result := i.expr()
		i.eat(RPAREN)
		return result
	}
	panic("Syntax error")
}

func (i *Interpreter) term() int {
	result := i.factor()

	for i.currentToken._type == MUL || i.currentToken._type == DIV {
		if i.currentToken._type == MUL {
			i.eat(MUL)
			result *= i.factor()
		}
		if i.currentToken._type == DIV {
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
		}
		if i.currentToken._type == MINUS {
			i.eat(MINUS)
			result -= i.term()
		}
	}
	return result
}
