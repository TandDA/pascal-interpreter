package main

type OldTreeNode struct {
	val   int
	op    string
	left  *OldTreeNode
	right *OldTreeNode
}

type AST struct {
	head *OldTreeNode
}

func NewAST(num int) *AST {
	node := &OldTreeNode{val: num}
	return &AST{head: node}
}

func (t *AST) addOperation(_type string, num int) {
	switch _type {
	case MINUS, PLUS:
		newRight := &OldTreeNode{val: num}
		newHead := &OldTreeNode{op: _type, left: t.head, right: newRight}
		t.head = newHead
	case MUL, DIV:
		newRightNum := &OldTreeNode{val: num}
		newRightOp := &OldTreeNode{op: _type, left: t.head.right, right: newRightNum}
		t.head.right = newRightOp
	}
}

func (t *AST) calc() int {

	var f func(node *OldTreeNode) int
	f = func(node *OldTreeNode) int {
		if node.left == nil {
			return node.val
		}
		switch node.op {
		case MUL:
			return f(node.left) * f(node.right)
		case DIV:
			return f(node.left) / f(node.right)
		case PLUS:
			return f(node.left) + f(node.right)
		case MINUS:
			return f(node.left) - f(node.right)
		}
		return node.val
	}
	return f(t.head)
}
