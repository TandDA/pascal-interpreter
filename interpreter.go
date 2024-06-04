package main

type Interpreter struct {
	parser *Parser
}

func NewInterpreter(parser *Parser) *Interpreter {
	return &Interpreter{parser: parser}
}

func (i *Interpreter) interpret() int {
	visitor := &NodeVisitor{}
	return i.parser.expr().Accept(visitor).(int)
}

type NodeVisitor struct{}

func (nv *NodeVisitor) VisitBinOpNode(n *BinOpNode) any {
	l := n.left.Accept(nv).(int)
	r := n.right.Accept(nv).(int)
	switch n.token._type {
	case MUL:
		return l * r
	case DIV:
		return l / r
	case PLUS:
		return l + r
	case MINUS:
		return l - r
	}
	return nil
}

func (nv *NodeVisitor) VisitNumNode(n *NumNode) any {
	return n.token.val.(int)
}

func (nv *NodeVisitor) VisitUnaryOp(n *UnaryOpNode) any {
	op := n.token._type
	num := n.expr.Accept(nv).(int)
	if op == MINUS {
		return -num
	} else {
		return num
	}
}
