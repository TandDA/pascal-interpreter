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
	i.eat(PROGRAM)
	programName := i.variable().(*VarNode).value
	i.eat(SEMI)
	blockNode := i.block()
	i.eat(DOT)
	return &ProgramNode{name: programName, block: blockNode}
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
	node := &VarNode{token: i.currentToken, value: i.currentToken.val.(string)}
	i.eat(ID)
	return node
}
func (i *Parser) empty() TreeNode {
	return &NoOpNode{}
}

func (i *Parser) block() TreeNode {
	declarationNodes := i.declarations()
	compoundStatementNodes := i.compoundStatement()
	return &BlockNode{declarations: declarationNodes, compoundStatement: compoundStatementNodes}
}

func (i *Parser) declarations() []TreeNode {
	declarations := []TreeNode{}
	if i.currentToken._type == VAR {
		i.eat(VAR)
		for i.currentToken._type == ID {
			varDecl := i.variableDeclaration()
			declarations = append(declarations, varDecl...)
			i.eat(SEMI)
		}
	}

	for i.currentToken._type == PROCEDURE {
		i.eat(PROCEDURE)
		procName := i.currentToken.val.(string)
		i.eat(ID)
		i.eat(SEMI)
		blockNode := i.block()
		procDel := ProcedureDeclNode{procName: procName, blockNode: blockNode}
		declarations = append(declarations, &procDel)
		i.eat(SEMI)
	}
	return declarations
}

func (i *Parser) variableDeclaration() []TreeNode {
	varNodes := []VarNode{VarNode{token: i.currentToken}}
	i.eat(ID)

	for i.currentToken._type == COMMA {
		i.eat(COMMA)
		varNodes = append(varNodes, VarNode{token: i.currentToken})
		i.eat(ID)
	}

	i.eat(COLON)
	typeNode := i.typeSpec()
	varDeclarations := []TreeNode{}

	for i, _ := range varNodes {
		varDeclarations = append(varDeclarations, &VarDeclNode{varNode: &varNodes[i], typeNode: typeNode})
	}
	return varDeclarations

}
func (i *Parser) typeSpec() TreeNode {
	token := i.currentToken
	if token._type == INTEGER {
		i.eat(INTEGER)
	} else {
		i.eat(REAL)
	}
	return &TypeNode{token: token}
}
func (i *Parser) factor() TreeNode {
	if i.currentToken._type == PLUS {
		tok := i.currentToken
		i.eat(PLUS)
		return &UnaryOpNode{token: tok, expr: i.factor()}
	} else if i.currentToken._type == MINUS {
		tok := i.currentToken
		i.eat(MINUS)
		return &UnaryOpNode{token: tok, expr: i.factor()}
	} else if i.currentToken._type == INTEGER_CONST {
		tok := i.currentToken
		i.eat(INTEGER_CONST)
		return &NumNode{token: tok}
	} else if i.currentToken._type == REAL_CONST {
		tok := i.currentToken
		i.eat(REAL_CONST)
		return &NumNode{token: tok}
	} else if i.currentToken._type == LPAREN {
		i.eat(LPAREN)
		node := i.expr()
		i.eat(RPAREN)
		return node
	} else if i.currentToken._type == ID {
		return i.variable()
	} else {
		return i.variable()
	}
}

func (i *Parser) term() TreeNode {
	result := i.factor()

	for i.currentToken._type == MUL ||
		i.currentToken._type == INTEGER_DIV ||
		i.currentToken._type == FLOAT_DIV {
		op := i.currentToken
		i.eat(op._type)
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
