package main

import (
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/parser"
)

func main() {
	lex := lexer.New(`
	PRG START
		JPL 0000,GR1
		JPL 1234
		JPL PRG
		JOV 0000,GR1
		JOV 1234
		JOV PRG
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
