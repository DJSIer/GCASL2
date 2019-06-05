package main

import (
	"fmt"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/token"
)

func main() {
	lex := lexer.New(`
	RAMEN LD GR2,0 
	LAD GR2,0`)

	for {
		t := lex.NextToken()
		fmt.Println(t)
		if t.Type == token.EOF {
			break
		}
	}

}
