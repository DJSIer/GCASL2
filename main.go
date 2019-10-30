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
	ADDSUB  START
	LAD     GR1, 1024
	LAD     GR2, 2048
	ADDA    GR1, GR2

	LAD     GR3, #1000
	LAD     GR4, #2000
	ADDL    GR3, GR4

	LAD     GR5200
	LAD     GR6, 100
	SUBA    GR5, GR6
	SUBL    GR6, =300
	RET
	END

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
	var buf bytes.Buffer
	b, _ := json.Marshal(code)
	buf.Write(b)
	fmt.Println("{\"result\":\"OK\",\"code\" :" + buf.String() + "}")
}
