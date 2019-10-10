package main

import (
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/parser"
)

func main() {
	lex := lexer.New(`
	PRG START
	GO LAD GR1,0,GR2
	LAD GR2,B
	RET
	A DS 1
	END
	`)
	p := parser.New(lex)
	code := p.ParseProgram()
	fmt.Println(code)
	fmt.Println(p.LabelToAddress(code))
	fmt.Println(p.Errors())

	/*for {
		t := lex.NextToken()
		fmt.Println(t)
		if t.Type == token.EOF {
			break
		}
	}*/

}
