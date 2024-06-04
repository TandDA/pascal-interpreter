package main

import "fmt"

// reverse polish notation
func RPN(node *OldTreeNode) string {
	var f func(node *OldTreeNode) string
	f = func(node *OldTreeNode) string {
		if node.t._type == INTEGER {
			return fmt.Sprint(node.t.val.(int))
		}
		return f(node.left) + " " + f(node.right) + " " + node.t.val.(string) + " "
	}
	return f(node)
}

// LISP notation
func LISP(node *OldTreeNode) string {
	var f func(node *OldTreeNode) string
	f = func(node *OldTreeNode) string {
		if node.t._type == INTEGER {
			return fmt.Sprint(node.t.val.(int))
		}
		return fmt.Sprintf("(%s %s %s)", node.t.val, f(node.left), f(node.right))
	}
	return f(node)
}
