package main

type TreeNode struct {
	val   int
	op    string
	left  *TreeNode
	right *TreeNode
}

type CalcTree struct {
	head    *TreeNode
	current *TreeNode
}

func NewCalcTree(num int) *CalcTree {
	node := &TreeNode{val: num}
	return &CalcTree{
		head:    node,
		current: node,
	}
}

func (t *CalcTree) addOperation(_type string, num int) {
	switch _type {
	case MINUS, PLUS:
		newCurr := &TreeNode{val: num}
		newHead := &TreeNode{op: _type, left: t.head, right: newCurr}
		t.head = newHead
		t.current = newCurr
	case MULTIP, DIV:
		newCurr := &TreeNode{val: num}
		newLeft := &TreeNode{val: t.current.val}
		t.current.op = _type
		t.current.left = newLeft
		t.current.right = newCurr
		t.current = newCurr
	}
}

func (t *CalcTree) calc() int {
	
	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if node.left == nil {
			return node.val
		}
		switch node.op {
		case MULTIP:
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
