package main

import "fmt"

func main() {
	program := `
PROGRAM Part10AST;
VAR
   a, b : INTEGER;
   y    : REAL;

BEGIN {Part10AST}
   a := 2;
   b := 10 * a + 10 * a DIV 4;
   y := 20 / 7 + 3.14;
END.  {Part10AST}
		`
	lexer := NewLexer(program)
	parser := NewParser(lexer)
	intrpr := NewInterpreter(parser)
	fmt.Println(intrpr.interpret())
	fmt.Println(intrpr.globalScope)
}
