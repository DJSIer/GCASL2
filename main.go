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
	OVERFLOW START
			LAD     GR0, 32765
			LAD     GR1, 1
	LOOP0   ADDA    GR0, GR1
			JOV     OUTLOOP0
			JUMP    LOOP0
	OUTLOOP0
			LAD     GR0, 32767
			LAD     GR1, 1
	LOOP1   SUBA    GR0, GR1
			JOV     OUTLOOP1
			JUMP    LOOP1
	OUTLOOP1
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
