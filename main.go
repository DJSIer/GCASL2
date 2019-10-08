package main

import (
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/parser"
)

func main() {
	lex := lexer.New(`
	PRG START
		GO LAD	GR1,GO
		GO LAD	GR1,GO
		RET
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
