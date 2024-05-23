package main

import "fmt"

func main() {
	lexer := NewLexer("8 + 2 * 3  - 7 * 2 * 3          ")
	intr := NewInerpreter(lexer)
	fmt.Println(intr.expr())
}
