package main

import "fmt"

func main() {
	lexer := NewLexer("2 + 3 * 5")
	//lexer := NewLexer("7 + 3 * (10 / (12 / (3 + 1) - 1)) / (2 + 3) - 5 - 3 + (8)")
	parser := NewParser(lexer)
	//intrpr := NewInterpreter(parser)
	//fmt.Println(intrpr.interpret())
	fmt.Println(LISP(parser.expr()))
}
