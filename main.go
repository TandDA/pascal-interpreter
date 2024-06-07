package main

import "fmt"

func main() {
	program := `
	BEGIN
	    begin
	        number := 2;
	        a := number;
	        b := 10 * a + 10 * number div 4;
	        c := a - - b
	    END;
	    x := 11;
	 end.
		`
	lexer := NewLexer(program)
	parser := NewParser(lexer)
	intrpr := NewInterpreter(parser)
	fmt.Println(intrpr.interpret())
	fmt.Println(intrpr.globalScope)
}
