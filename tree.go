package main

type TreeNode struct {
	val   int
	op    string
	left  *TreeNode
	right *TreeNode
}

type CalcTree struct {
	head *TreeNode
}

func NewCalcTree(num int) *CalcTree {
	node := &TreeNode{val: num}
	return &CalcTree{head: node}
}

func (t *CalcTree) addOperation(_type string, num int) {
	switch _type {
	case MINUS, PLUS:
		newRight := &TreeNode{val: num}
		newHead := &TreeNode{op: _type, left: t.head, right: newRight}
		t.head = newHead
	case MUL, DIV:
		newRightNum := &TreeNode{val: num}
		newRightOp := &TreeNode{op: _type, left: t.head.right, right: newRightNum}
		t.head.right = newRightOp
	}
}

func (t *CalcTree) calc() int {

	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
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
