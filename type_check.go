package main

import (
	"reflect"
)

type SymbolTable map[string]reflect.Type

type SymbolTableVisitor struct {
	m SymbolTable
}

func NewSymbolTableVisitor() *SymbolTableVisitor {
	s := SymbolTableVisitor{m: SymbolTable{
		"INTEGER": reflect.TypeOf(1),
		"REAL":    reflect.TypeOf(float32(1)),
	}}
	return &s
}

func (s SymbolTableVisitor) VisitBinOpNode(n *BinOpNode) any {
	n.left.Accept(s)
	n.right.Accept(s)
	return nil
}

func (s SymbolTableVisitor) VisitNumNode(n *NumNode) any {
	return nil
}

func (s SymbolTableVisitor) VisitUnaryOp(n *UnaryOpNode) any {
	n.expr.Accept(s)
	return nil
}

func (s SymbolTableVisitor) VisitCompoundNode(n *CompoundNode) any {
	for _, child := range n.children {
		child.Accept(s)
	}
	return nil
}

func (s SymbolTableVisitor) VisitAssignNode(n *AssignNode) any {
	varName := n.left.(*VarNode).token.val.(string)
	if _, ok := s.m[varName]; !ok {
		panic(varName + " undefiend")
	}
	return nil
}

func (s SymbolTableVisitor) VisitVarNode(n *VarNode) any {
	varName := n.token.val.(string)
	if _, ok := s.m[varName]; !ok {
		panic(varName + " undefiend")
	}
	return nil
}

func (s SymbolTableVisitor) VisitNoOpNode(n *NoOpNode) any {
	return nil
}

func (s SymbolTableVisitor) VisitTypeNode(n *TypeNode) any {
	return nil
}

func (s SymbolTableVisitor) VisitVarDeclNode(n *VarDeclNode) any {
	typeName := n.typeNode.(*TypeNode).token.val.(string)
	typeType := s.m[typeName]
	varName := n.varNode.(*VarNode).token.val.(string)
	s.m[varName] = typeType
	return nil
}

func (s SymbolTableVisitor) VisitBlockNode(n *BlockNode) any {
	for _, decl := range n.declarations {
		decl.Accept(s)
	}
	n.compoundStatement.Accept(s)
	return nil
}

func (s SymbolTableVisitor) VisitProgramNode(n *ProgramNode) any {
	n.block.Accept(s)
	return nil
}
