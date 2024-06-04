package main

type TreeNode interface {
	Accept(visitor Visitor) any
}

type Visitor interface {
	VisitBinOpNode(n *BinOpNode) any
	VisitNumNode(n *NumNode) any
	VisitUnaryOp(n *UnaryOpNode) any
}

type BinOpNode struct {
	left  TreeNode
	right TreeNode
	token Token
}

func (n *BinOpNode) Accept(visitor Visitor) any {
	return visitor.VisitBinOpNode(n)
}

type NumNode struct {
	token Token
}

func (n *NumNode) Accept(visitor Visitor) any {
	return visitor.VisitNumNode(n)
}

type UnaryOpNode struct {
	token Token
	expr  TreeNode
}

func (n *UnaryOpNode) Accept(visitor Visitor) any {
	return visitor.VisitUnaryOp(n)
}
