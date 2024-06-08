package main

type Interpreter struct {
	parser      *Parser
	globalScope map[string]float64
}

func NewInterpreter(parser *Parser) *Interpreter {
	return &Interpreter{parser: parser, globalScope: make(map[string]float64)}
}

func (i *Interpreter) interpret() any {
	tree := i.parser.parse()
	if tree == nil {
		return 0
	}
	return tree.Accept(i)
}

func (nv *Interpreter) VisitBinOpNode(n *BinOpNode) any {
	l := n.left.Accept(nv).(float64)
	r := n.right.Accept(nv).(float64)
	switch n.token._type {
	case MUL:
		return l * r
	case INTEGER_DIV:
		return float64(int(l / r))
	case FLOAT_DIV:
		return l / r
	case PLUS:
		return l + r
	case MINUS:
		return l - r
	}
	return nil
}

func (nv *Interpreter) VisitNumNode(n *NumNode) any {
	return n.token.val.(float64)
}

func (nv *Interpreter) VisitUnaryOp(n *UnaryOpNode) any {
	op := n.token._type
	num := n.expr.Accept(nv).(float64)
	if op == MINUS {
		return -num
	} else {
		return num
	}
}

func (nv *Interpreter) VisitCompoundNode(n *CompoundNode) any {
	for _, child := range n.children {
		child.Accept(nv)
	}
	return nil
}

func (nv *Interpreter) VisitAssignNode(n *AssignNode) any {
	varName := n.left.(*VarNode).value
	nv.globalScope[varName] = n.right.Accept(nv).(float64)
	return nil
}

func (nv *Interpreter) VisitVarNode(n *VarNode) any {
	val, ok := nv.globalScope[n.value]
	if ok {
		return val
	} else {
		panic("Var " + n.value + " undefiend")
	}
}

func (nv *Interpreter) VisitNoOpNode(n *NoOpNode) any {
	return NoOpNode{}
}

func (nv *Interpreter) VisitProgramNode(n *ProgramNode) any {
	n.block.Accept(nv)
	return nil
}
func (nv *Interpreter) VisitBlockNode(n *BlockNode) any {
	for _, declaration := range n.declarations {
		declaration.Accept(nv)
	}
	n.compoundStatement.Accept(nv)
	return nil
}
func (nv *Interpreter) VisitVarDeclNode(n *VarDeclNode) any {
	return nil
}

func (nv *Interpreter) VisitTypeNode(n *TypeNode) any {
	return nil
}
