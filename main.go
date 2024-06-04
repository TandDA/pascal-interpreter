package main

import "fmt"

func main() {

	//lexer := NewLexer("7 + 3 * (10 / (12 / (3 + 1) - 1)) / (2 + 3) - 5 - 3 + (8)") // 10
	lexer := NewLexer("7---3")
	parser := NewParser(lexer)
	intrpr := NewInterpreter(parser)
	fmt.Println(intrpr.interpret())

}
