package main

import (
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/token"
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

	LAD     GR5, 200
	LAD     GR6, 100
	SUBA    GR5, GR6
	SUBL    GR6, =300
	RET
	END`)

	for {
		t := lex.NextToken()
		fmt.Println(t)
		if t.Type == token.EOF {
			break
		}
	}

}
