package main

import (
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/parser"
)

func main() {
	lex := lexer.New(`
	PRG START
		LD GR1,GR2
		LAD GR1,0,GR2
		LD	GR1,GO
	GO	RET
	`)
	p := parser.New(lex)
	fmt.Println(p.ParseProgram())
	/*for {
		t := lex.NextToken()
		fmt.Println(t)
		if t.Type == token.EOF {
			break
		}
	}*/

}
