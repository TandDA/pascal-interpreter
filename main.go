package main

import "fmt"

func main() {
	lexer := NewLexer("7 + 3 * (10 / (12 / (3 + 1) - 1)) / (2 + 3) - 5 - 3 + (8)")
	intr := NewInterpreter(lexer)
	fmt.Println(intr.expr())
}
