package main

type TreeNode interface {
	Accept(visitor Visitor) any
}

type Visitor interface {
	VisitBinOpNode(n *BinOpNode) any
	VisitNumNode(n *NumNode) any
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
