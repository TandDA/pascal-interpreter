package main

type OldTreeNode struct {
	t     Token
	left  *OldTreeNode
	right *OldTreeNode
}

type Parser struct {
	l            *Lexer
	currentToken Token
}

func NewParser(lxr *Lexer) *Parser {
	return &Parser{l: lxr, currentToken: lxr.getNextToken()}
}

func (i *Parser) eat(_type int) {
	if _type == i.currentToken._type {
		i.currentToken = i.l.getNextToken()
		return
	}
	panic("Syntax error")
}

func (i *Parser) program() TreeNode {
	node := i.compoundStatement()
	i.eat(DOT)
	return node
}

func (i *Parser) compoundStatement() TreeNode {
	i.eat(BEGIN)
	nodes := i.statementList()
	i.eat(END)

	root := &CompoundNode{}
	for _, node := range nodes {
		root.children = append(root.children, node)
	}
	return root
}

func (i *Parser) statementList() []TreeNode {
	node := i.statement()

	result := []TreeNode{node}
	for i.currentToken._type == SEMI {
		i.eat(SEMI)
		result = append(result, i.statement())
	}
	return result
}

func (i *Parser) statement() TreeNode {
	var node TreeNode
	if i.currentToken._type == BEGIN {
		node = i.compoundStatement()
	} else if i.currentToken._type == ID {
		node = i.assignmentStatement()
	} else {
		node = i.empty()
	}
	return node
}

func (i *Parser) assignmentStatement() TreeNode {
	left := i.variable()
	token := i.currentToken
	i.eat(ASSIGN)

	right := i.expr()
	return &AssignNode{left, right, token}
}
func (i *Parser) variable() TreeNode {
	node := &VarNode{token: i.currentToken}
	i.eat(ID)
	return node
}
func (i *Parser) empty() TreeNode {
	return &NoOpNode{}
}
func (i *Parser) factor() TreeNode {
	if i.currentToken._type == PLUS {
		tok := i.currentToken
		i.eat(PLUS)
		return &UnaryOpNode{token: tok, expr: i.factor()}
	}
	if i.currentToken._type == MINUS {
		tok := i.currentToken
		i.eat(MINUS)
		return &UnaryOpNode{token: tok, expr: i.factor()}
	}
	if i.currentToken._type == INTEGER {
		tok := i.currentToken
		i.eat(INTEGER)
		return &NumNode{token: tok}
	}
	if i.currentToken._type == LPAREN {
		i.eat(LPAREN)
		node := i.expr()
		i.eat(RPAREN)
		return node
	}
	if i.currentToken._type == ID {
		return i.variable()
	}
	panic("Syntax error")
}

func (i *Parser) term() TreeNode {
	result := i.factor()

	for i.currentToken._type == MUL || i.currentToken._type == DIV {
		op := i.currentToken
		if i.currentToken._type == MUL {
			i.eat(MUL)
		}
		if i.currentToken._type == DIV {
			i.eat(DIV)
		}
		result = &BinOpNode{left: result, token: op, right: i.factor()}
	}
	return result
}

func (i *Parser) expr() TreeNode {
	result := i.term()

	for i.currentToken._type == PLUS || i.currentToken._type == MINUS {
		op := i.currentToken
		if i.currentToken._type == PLUS {
			i.eat(PLUS)
		}
		if i.currentToken._type == MINUS {
			i.eat(MINUS)
		}
		result = &BinOpNode{left: result, token: op, right: i.term()}
	}
	return result
}

func (i *Parser) parse() TreeNode {
	node := i.program()
	if i.currentToken._type != EOF {
		panic("EOF error")
	}
	return node
}
