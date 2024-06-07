package main

import "fmt"

func main() {
	program := `
	BEGIN
	    BEGIN
	        number := 2;
	        a := number;
	        b := 10 * a + 10 * number / 4;
	        c := a - - b
	    END;
	    x := 11;
	 END.
		`
	//lexer := NewLexer("7 + 3 * (10 / (12 / (3 + 1) - 1)) / (2 + 3) - 5 - 3 + (8)") // 10
	lexer := NewLexer(program)
	parser := NewParser(lexer)
	intrpr := NewInterpreter(parser)
	fmt.Println(intrpr.interpret())
	fmt.Println(intrpr.globalScope)
}
