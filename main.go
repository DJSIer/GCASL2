package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/parser"
)

func main() {
	lex := lexer.New(`
	PRG START
	A LD GR1,GR2
	LAD GR1,A,GR3
	ADDA GR1,A
	A DS 1
	  DC 1010,16
	`)
	p := parser.New(lex)
	code, err := p.ParseProgram()
	if err != nil {
		var buf bytes.Buffer
		b, _ := json.Marshal(p.Errors())
		buf.Write(b)
		fmt.Println("{\"Result\":\"NG\",\"error\" :" + buf.String() + "}")
		return
	}
	code, err = p.LabelToAddress(code)
	if err != nil {
		var buf bytes.Buffer
		b, _ := json.Marshal(p.Errors())
		buf.Write(b)
		fmt.Println("{\"Result\":\"NG\",\"error\" :" + buf.String() + "}")
		return
	}
	var buf bytes.Buffer
	b, _ := json.Marshal(code)
	buf.Write(b)
	fmt.Println("{\"Result\":\"OK\",\"code\" :" + buf.String() + "}")
}
