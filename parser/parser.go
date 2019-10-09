package parser

import (
	"fmt"
	"strconv"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/opcode"
	"github.com/DJSIer/GCASL2/symbol"
	"github.com/DJSIer/GCASL2/token"
)

var registerNumber = map[string]uint8{
	"GR0": 0x00,
	"GR1": 0x01,
	"GR2": 0x02,
	"GR3": 0x03,
	"GR4": 0x04,
	"GR5": 0x05,
	"GR6": 0x06,
	"GR7": 0x07,
}

type functype func(*opcode.Opcode) *opcode.Opcode

// Parser CASL2 Assembly Parser Struct
type Parser struct {
	l           *lexer.Lexer
	curToken    token.Token
	peekToken   token.Token
	byteAdress  uint16
	symbolTable *symbol.SymbolTable
	errors      []string
	warnings    []string
	instSet     map[token.TokenType]functype
}

// New Parser New
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:        l,
		errors:   []string{},
		warnings: []string{},
	}
	p.instSet = map[token.TokenType]functype{
		token.LD:    p.LDStatment,
		token.LAD:   p.LADStatment,
		token.ST:    p.STStatment,
		token.ADDA:  p.ADDAStatment,
		token.SUBA:  p.SUBAStatment,
		token.ADDL:  p.ADDLStatment,
		token.SUBL:  p.SUBLStatment,
		token.AND:   p.ANDStatment,
		token.OR:    p.ORStatment,
		token.XOR:   p.XORStatment,
		token.CPA:   p.CPAStatment,
		token.CPL:   p.CPLStatment,
		token.SLA:   p.SLAStatment,
		token.SRA:   p.SRAStatment,
		token.SLL:   p.SLLStatment,
		token.SRL:   p.SRLStatment,
		token.JMI:   p.JMIStatment,
		token.START: p.STARTStatment,
		token.RET:   p.RETStatment,
	}
	p.symbolTable = symbol.NewSymbolTable()
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// Errors CASL2 Parse Error message
func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// ParseProgram CASL2 Parse
func (p *Parser) ParseProgram() []opcode.Opcode {
	Excode := []opcode.Opcode{}
	for !p.curTokenIs(token.EOF) {
		code := &opcode.Opcode{}
		//Label
		if p.curTokenIs(token.LABEL) {
			sy, flag := p.symbolTable.Define(p.curToken.Literal, p.byteAdress)
			if flag {
				code.Label = sy
			} else {
				msg := fmt.Sprintf("重複定義エラー Label:%s\n", p.curToken.Literal)
				p.errors = append(p.errors, msg)
			}
			p.nextToken()
		}
		switch p.curToken.Type {
		case token.LAD:
			code = p.instSet[p.curToken.Type](code)
		case token.LD:
			code = p.instSet[p.curToken.Type](code)
		case token.ST:
			code = p.instSet[p.curToken.Type](code)
		case token.ADDA:
			code = p.instSet[p.curToken.Type](code)
		case token.SUBA:
			code = p.instSet[p.curToken.Type](code)
		case token.ADDL:
			code = p.instSet[p.curToken.Type](code)
		case token.SUBL:
			code = p.instSet[p.curToken.Type](code)
		case token.AND:
			code = p.instSet[p.curToken.Type](code)
		case token.OR:
			code = p.instSet[p.curToken.Type](code)
		case token.XOR:
			code = p.instSet[p.curToken.Type](code)
		case token.CPA:
			code = p.instSet[p.curToken.Type](code)
		case token.CPL:
			code = p.instSet[p.curToken.Type](code)
		case token.SLA:
			code = p.instSet[p.curToken.Type](code)
		case token.SRA:
			code = p.instSet[p.curToken.Type](code)
		case token.SLL:
			code = p.instSet[p.curToken.Type](code)
		case token.SRL:
			code = p.instSet[p.curToken.Type](code)
		case token.JMI:
			code = p.instSet[p.curToken.Type](code)
		case token.START:
			code = p.instSet[p.curToken.Type](code)
		case token.RET:
			code = p.instSet[p.curToken.Type](code)
		default:
			code = nil
		}
		if code != nil {
			Excode = append(Excode, *code)
			p.byteAdress += uint16(code.Length)
		}
		p.nextToken()
	}
	return Excode
}

// STARTStatment `Label START` - [実行番地]
// START プログラムの実行番地を定義
func (p *Parser) STARTStatment(code *opcode.Opcode) *opcode.Opcode {
	sy, ok := p.symbolTable.Resolve(code.Label.Label)
	if !ok {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	code = &opcode.Opcode{Op: 0x00, Code: 0x0000, Length: 1, Label: sy}
	return code
}

// RETStatment Return from subroutine Parser
// RET ;PR ← ((SP)),
//	   ;SP ← (SP) + 1
func (p *Parser) RETStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x81, Code: 0x8100, Length: 1, Label: code.Label}
	return code
}

// LDStatment Load Parser
// LD r1, r2 		;r1 ← (r2)
// LD r, adr [,x] 	;r  ← (実行アドレス)
func (p *Parser) LDStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		msg := "no ,"
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x10
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x14
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.LABEL:
		code.Op = 0x10
		code.AddrLabel = p.curToken.Literal
		code.Length = 2
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// LADStatment Load Address Parser
// LAD r,adr [,x] ; r ← 実行アドレス
func (p *Parser) LADStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x12, Code: 0x1200, Length: 2, Label: code.Label}

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	if !p.expectPeek(token.COMMA) {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			msg := fmt.Sprintf("65535以上です。対象数値:%q ", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}

		code.Addr = uint16(addr)
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
	}
	if !p.peekTokenIs(token.COMMA) {
		return code
	}
	p.nextToken()
	if !p.peekTokenIs(token.REGISTER) {
		msg := fmt.Sprintf("レジスタではありません。parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])

	return code
}

// STStatment Store Parser
// ST r,adr [,x];実行アドレス ← (r)
func (p *Parser) STStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Code: 0x1100, Op: 0x11, Length: 2, Label: code.Label}
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	if !p.expectPeek(token.COMMA) {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
	}
	if !p.peekTokenIs(token.COMMA) {
		return code
	}
	p.nextToken()
	if !p.peekTokenIs(token.REGISTER) {
		msg := fmt.Sprintf("レジスタではありません。parse error %q", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// ADDAStatment ADD Arithmetic Parser
// ADDA r1, r2		; r1  ← (r1) + (r2)
// ADDA r, adr [,x]	; r   ← (r)  + 実行アドレス
func (p *Parser) ADDAStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		msg := "no ,"
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x20
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x24
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.LABEL:
		code.Op = 0x20
		code.AddrLabel = p.curToken.Literal
		code.Length = 2
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

func (p *Parser) SUBAStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) {
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x21
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x25
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// ADDLStatment ADD logical Parser
// ADDL r1, r2		; r1  ← (r1) + (r2)
// ADDL r, adr [,x]	; r   ← (r)  + 実行アドレス
func (p *Parser) ADDLStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		msg := "no ,"
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	p.nextToken()

	code.Op = 0x22
	code.Length = 2

	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x26
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}
func (p *Parser) SUBLStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) {
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x23
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x27
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

func (p *Parser) ANDStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) {
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x30
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x34
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}
func (p *Parser) ORStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) {
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x31
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x35
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}
func (p *Parser) XORStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) {
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x32
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x36
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

//compare arithmetic
func (p *Parser) CPAStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) {
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x40
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x44
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

//compare logical
func (p *Parser) CPLStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) {
		return nil
	}
	p.nextToken()

	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x41
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x45
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

//shift left arithmetic
func (p *Parser) SLAStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) {
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x50
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8
	return code
}

//shift right arithmetic
func (p *Parser) SRAStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) {
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x51
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8
	return code
}

//shift left logical
func (p *Parser) SLLStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) {
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x52
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8
	return code
}

//shift right logical
func (p *Parser) SRLStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
	if !p.peekTokenIs(token.INT) {
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		code.Op = 0x53
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)

		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8
	return code
}

// JMIStatment JMI
func (p *Parser) JMIStatment(code *opcode.Opcode) *opcode.Opcode {
	code.Op = 0x61
	code.Length = 2
	if !p.expectPeek(token.INT) {
		return nil
	}
	addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
	if err != nil {
		msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	code.Addr = uint16(addr)
	if !p.peekTokenIs(token.COMMA) {
		return nil
	}
	p.nextToken()
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	code.Code |= uint16(code.Op) << 8
	return code
}
