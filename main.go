package main

import (
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/parser"
)

func main() {
	lex := lexer.New(`
	PRG START
		ADDL GR1,GR2
		ADDL GR1,16
		ADDL GR1,16,GR2
		ADDL GR1,GO,GR2 
	`)
	p := parser.New(lex)
	fmt.Println(p.ParseProgram())
	fmt.Println(p.Errors())
	/*for {
		t := lex.NextToken()
		fmt.Println(t)
		if t.Type == token.EOF {
			break
		}
	}*/

}
