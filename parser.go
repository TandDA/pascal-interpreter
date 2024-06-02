package main

type TreeNode struct {
	t     Token
	left  *TreeNode
	right *TreeNode
}

type Parser struct {
	l            *Lexer
	currentToken Token
}

func NewParser(lxr *Lexer) *Parser {
	return &Parser{l: lxr, currentToken: lxr.getNextToken()}
}

func (i *Parser) eat(_type string) {
	if _type == i.currentToken._type {
		i.currentToken = i.l.getNextToken()
		return
	}
	panic("Syntax error")
}

func (i *Parser) factor() *TreeNode {
	if i.currentToken._type == PLUS {
		tok := i.currentToken
		i.eat(PLUS)
		return &TreeNode{t: tok, left: i.factor()}
	}
	if i.currentToken._type == MINUS {
		tok := i.currentToken
		i.eat(MINUS)
		return &TreeNode{t: tok, left: i.factor()}
	}
	if i.currentToken._type == INTEGER {
		tok := i.currentToken
		i.eat(INTEGER)
		return &TreeNode{t: tok}
	}
	if i.currentToken._type == LPAREN {
		i.eat(LPAREN)
		node := i.expr()
		i.eat(RPAREN)
		return node
	}
	panic("Syntax error")
}

func (i *Parser) term() *TreeNode {
	result := i.factor()

	for i.currentToken._type == MUL || i.currentToken._type == DIV {
		op := i.currentToken
		if i.currentToken._type == MUL {
			i.eat(MUL)
		}
		if i.currentToken._type == DIV {
			i.eat(DIV)
		}
		result = &TreeNode{left: result, t: op, right: i.factor()}
	}
	return result
}

func (i *Parser) expr() *TreeNode {
	result := i.term()

	for i.currentToken._type == PLUS || i.currentToken._type == MINUS {
		op := i.currentToken
		if i.currentToken._type == PLUS {
			i.eat(PLUS)
		}
		if i.currentToken._type == MINUS {
			i.eat(MINUS)
		}
		result = &TreeNode{left: result, t: op, right: i.term()}
	}
	return result
}
