package main

type Interpreter struct {
	parser *Parser
}

func NewInterpreter(parser *Parser) *Interpreter {
	return &Interpreter{parser: parser}
}

func (i *Interpreter) interpret() int {
	var visit func(node *TreeNode) int
	visit = func(node *TreeNode) int {
		if node.left == nil {
			return node.t.val.(int)
		}
		switch node.t._type {
		case MUL:
			return visit(node.left) * visit(node.right)
		case DIV:
			return visit(node.left) / visit(node.right)
		case PLUS:
			return visit(node.left) + visit(node.right)
		case MINUS:
			return visit(node.left) - visit(node.right)
		}
		return node.t.val.(int)
	}

	node := i.parser.expr()
	return visit(node)
}
