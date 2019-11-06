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
	ST GR1,0
	`)

	p := parser.New(lex)
	code, err := p.ParseProgram()
	if err != nil {
		var buf bytes.Buffer
		b, _ := json.Marshal(p.Errors())
		buf.Write(b)
		fmt.Println("{\"result\":\"NG\",\"error\" :" + buf.String() + "}")
		return
	}
	code, err = p.LiteralToMemory(code)
	code, err = p.LabelToAddress(code)
	if err != nil {
		var buf bytes.Buffer
		b, _ := json.Marshal(p.Errors())
		buf.Write(b)
		fmt.Println("{\"result\":\"NG\",\"error\" :" + buf.String() + "}")
		return
	}
	var buf, waningbuf bytes.Buffer
	b, _ := json.Marshal(code)
	buf.Write(b)
	warning, _ := json.Marshal(p.Warnings())
	waningbuf.Write(warning)
	fmt.Println("{\"result\":\"OK\",\"code\" :" + buf.String() + ",\"warning\" :" + waningbuf.String() + "}")
}
