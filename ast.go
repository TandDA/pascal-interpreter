package main

type TreeNode interface {
	Accept(visitor Visitor) any
}

type Visitor interface {
	VisitBinOpNode(n *BinOpNode) any
	VisitNumNode(n *NumNode) any
	VisitUnaryOp(n *UnaryOpNode) any
	VisitCompoundNode(n *CompoundNode) any
	VisitAssignNode(n *AssignNode) any
	VisitVarNode(n *VarNode) any
	VisitNoOpNode(n *NoOpNode) any
}

type (
	BinOpNode struct {
		left  TreeNode
		right TreeNode
		token Token
	}
	NumNode struct {
		token Token
	}
	UnaryOpNode struct {
		token Token
		expr  TreeNode
	}
	CompoundNode struct {
		children []TreeNode
	}
	AssignNode struct {
		left  TreeNode
		right TreeNode
		token Token
	}
	VarNode struct {
		token Token
		value string
	}
	NoOpNode struct {

	}
)

func (n *BinOpNode) Accept(visitor Visitor) any {
	return visitor.VisitBinOpNode(n)
}
func (n *NumNode) Accept(visitor Visitor) any {
	return visitor.VisitNumNode(n)
}
func (n *UnaryOpNode) Accept(visitor Visitor) any {
	return visitor.VisitUnaryOp(n)
}
func (n *CompoundNode) Accept(visitor Visitor) any {
	return visitor.VisitCompoundNode(n)
}
func (n *AssignNode) Accept(visitor Visitor) any {
	return visitor.VisitAssignNode(n)
}
func (n *VarNode) Accept(visitor Visitor) any {
	return visitor.VisitVarNode(n)
}
func (n *NoOpNode) Accept(visitor Visitor) any {
	return visitor.VisitNoOpNode(n)
}

