package parser

import (
	"fmt"
	"strconv"
	"strings"

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
	errors      []ParserError
	warnings    []string
	instSet     map[token.TokenType]functype
	Excode      []opcode.Opcode
	line        int //line number
}

// ParserError Parse Error Message struct
type ParserError struct {
	Line    int    //line number
	Message string //ErrorMessage
}

// New Parser New
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:        l,
		errors:   []ParserError{},
		warnings: []string{},
		line:     1,
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
		token.JNZ:   p.JNZStatment,
		token.JZE:   p.JZEStatment,
		token.JUMP:  p.JUMPStatment,
		token.JPL:   p.JPLStatment,
		token.JOV:   p.JOVStatment,
		token.PUSH:  p.PUSHStatment,
		token.POP:   p.POPStatment,
		token.START: p.STARTStatment,
		token.RET:   p.RETStatment,
		token.DS:    p.DSStatment,
		token.DC:    p.DCStatment,
		token.END:   p.ENDStatment,
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
func (p *Parser) Errors() []ParserError {
	return p.errors
}
func (p *Parser) peekError(t token.TokenType) {
	e := &ParserError{Line: p.line, Message: fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)}
	p.errors = append(p.errors, *e)
}
func (p *Parser) parserError(line int, msg string) {
	e := &ParserError{Line: line, Message: msg}
	p.errors = append(p.errors, *e)
}
func (p *Parser) parserWarning(msg string) {
	p.warnings = append(p.warnings, msg)
}

// ParseProgram CASL2 Parse
func (p *Parser) ParseProgram() ([]opcode.Opcode, error) {
	p.Excode = []opcode.Opcode{}
	for !p.curTokenIs(token.EOF) {
		code := &opcode.Opcode{Length: 1}
		//Label
		if p.curTokenIs(token.LABEL) {
			sy, flag := p.symbolTable.Define(p.curToken.Literal, p.byteAdress)
			if flag {
				code.Label = &sy
			} else {
				p.parserError(p.line, fmt.Sprintf("重複定義エラー Label : %q", p.curToken.Literal))
				return nil, fmt.Errorf("LABEL Error")
			}
			p.nextToken()
		}
		code.Token = p.curToken

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
		case token.JNZ:
			code = p.instSet[p.curToken.Type](code)
		case token.JZE:
			code = p.instSet[p.curToken.Type](code)
		case token.JUMP:
			code = p.instSet[p.curToken.Type](code)
		case token.START:
			code = p.instSet[p.curToken.Type](code)
		case token.RET:
			code = p.instSet[p.curToken.Type](code)
		case token.JPL:
			code = p.instSet[p.curToken.Type](code)
		case token.JOV:
			code = p.instSet[p.curToken.Type](code)
		case token.PUSH:
			code = p.instSet[p.curToken.Type](code)
		case token.POP:
			code = p.instSet[p.curToken.Type](code)
		case token.DS:
			code = p.instSet[p.curToken.Type](code)
		case token.DC:
			code = p.instSet[p.curToken.Type](code)
		case token.END:
			code = p.instSet[p.curToken.Type](code)
		default:
			//p.parserError(p.line, fmt.Sprintf("%q : 解決できません\n", p.curToken.Literal))
			code = nil
		}
		if code == nil {
			return p.Excode, fmt.Errorf("%q : コンパイルエラー", p.curToken)
		}

		p.Excode = append(p.Excode, *code)
		p.byteAdress += uint16(code.Length)

		p.nextToken()
		//program line number add
		p.line++
	}
	return p.Excode, nil
}

//LabelToAddress ラベルアドレスの解決
func (p *Parser) LabelToAddress(code []opcode.Opcode) ([]opcode.Opcode, error) {
	for i, op := range code {
		if len(op.AddrLabel) != 0 {
			addr, ok := p.symbolTable.Resolve(op.AddrLabel)
			if !ok {
				p.parserError(p.line, fmt.Sprintf("%qは解決できません", op.AddrLabel))
				return nil, nil
			}
			code[i].Addr = addr.Address
		}
	}
	return code, nil
}

//DCStatment 定数定義
func (p *Parser) DCStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x00, Code: 0x0000, Length: 1, Label: code.Label, Token: code.Token}
	if !p.peekTokenIs(token.INT) {
		p.parserError(p.line, fmt.Sprintf("数値でなければいけません。対象 : %q", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	num, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil
	}
	code.Addr = uint16(num)
	for {
		if !p.peekTokenIs(token.COMMA) {
			return code
		}
		p.nextToken()
		if !p.peekTokenIs(token.INT) {
			p.parserError(p.line, fmt.Sprintf("数値でなければいけません。対象 : %q", p.peekToken.Literal))
			return nil
		}
		p.nextToken()
		num, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			return nil
		}
		p.Excode = append(p.Excode, *code)
		p.byteAdress++
		code = &opcode.Opcode{Op: 0x00, Code: 0x0000, Length: 1, Token: code.Token}
		code.Addr = uint16(num)
	}
}

// DSStatment 領域確保
// [LABEL] DS NUM
func (p *Parser) DSStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x00, Code: 0x0000, Length: 1, Label: code.Label, Token: code.Token}
	if !p.peekTokenIs(token.INT) {
		p.parserError(p.line, fmt.Sprintf("数値でなければいけません。対象 : %q", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	Length, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil
	}
	code.Length = int(Length)
	return code
}

// STARTStatment `Label START` - [実行番地]
// START プログラムの実行番地を定義
func (p *Parser) STARTStatment(code *opcode.Opcode) *opcode.Opcode {
	sy, ok := p.symbolTable.Resolve(code.Label.Label)
	if !ok {
		p.parserError(p.line, fmt.Sprintf("STARTにラベルがありません。対象 : %q", p.peekToken.Literal))
		return nil
	}
	code = &opcode.Opcode{Op: 0x00, Code: 0x0000, Length: 1, Label: &sy, Token: code.Token}
	return code
}

// ENDStatment `END`
func (p *Parser) ENDStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x00, Code: 0x0000, Length: 1, Label: code.Label, Token: code.Token}
	return code
}

// RETStatment Return from subroutine Parser
// RET ;PR ← ((SP)),
//	   ;SP ← (SP) + 1
func (p *Parser) RETStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x81, Code: 0x8100, Length: 1, Label: code.Label, Token: code.Token}
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
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x10
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}

	case token.REGISTER:
		code.Op = 0x14
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

// LADStatment Load Address Parser
// LAD r,adr [,x] ; r ← 実行アドレス
func (p *Parser) LADStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x12, Code: 0x1200, Length: 2, Label: code.Label, Token: code.Token}

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	if !p.expectPeek(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}

	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
	}
	if !p.peekTokenIs(token.COMMA) {
		return code
	}
	p.nextToken()
	if !p.peekTokenIs(token.REGISTER) {
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])

	return code
}

// STStatment Store Parser
// ST r,adr [,x];実行アドレス ← (r)
func (p *Parser) STStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Code: 0x1100, Op: 0x11, Length: 2, Label: code.Label, Token: code.Token}
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	if !p.expectPeek(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
	}
	if !p.peekTokenIs(token.COMMA) {
		return code
	}
	p.nextToken()
	if !p.peekTokenIs(token.REGISTER) {
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// ADDAStatment ADD Arithmetic Parser
// ADDA r1, r2		; r1  ← (r1) + (r2)
// ADDA r, adr [,x]	; r   ← (r)  + (実行アドレス)
func (p *Parser) ADDAStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x20
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 64)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.REGISTER:
		code.Op = 0x24
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		if !p.peekTokenIs(token.REGISTER) {
			p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q", p.peekToken.Literal))
			return nil
		}
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// SUBAStatment  subtract arithmetic Parser
// SUBA r1, r2		; r1  ← (r1) - (r2)
// SUBA r, adr [,x]	; r   ← (r)  - (実行アドレス)
func (p *Parser) SUBAStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x21
	code.Length = 2

	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		if !p.peekTokenIs(token.REGISTER) {
			p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q", p.peekToken.Literal))
			return nil
		}
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		if !p.peekTokenIs(token.REGISTER) {
			p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q", p.peekToken.Literal))
			return nil
		}
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// ADDLStatment ADD logical Parser
// ADDL r1, r2		; r1  ← (r1) + (r2)
// ADDL r, adr [,x]	; r   ← (r)  + (実行アドレス)
func (p *Parser) ADDLStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()

	code.Op = 0x22
	code.Length = 2

	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		if !p.peekTokenIs(token.REGISTER) {
			p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q", p.peekToken.Literal))
			return nil
		}
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		if !p.peekTokenIs(token.REGISTER) {
			p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q", p.peekToken.Literal))
			return nil
		}
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// SUBLStatment  subtract logical Parser
// SUBL r1, r2		; r1  ← (r1) - (r2)
// SUBL r, adr [,x]	; r   ← (r)  - (実行アドレス)
func (p *Parser) SUBLStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()

	code.Op = 0x23
	code.Length = 2

	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.REGISTER:
		code.Op = 0x27
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		if !p.peekTokenIs(token.REGISTER) {
			p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q", p.peekToken.Literal))
			return nil
		}
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// ANDStatment  AND Parser
// AND r1, r2		; r1  ← (r1) AND (r2)
// AND r, adr [,x]	; r   ← (r)  AND (実行アドレス)
func (p *Parser) ANDStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x30
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.REGISTER:
		code.Op = 0x34
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
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// ORStatment  OR Parser
// OR r1, r2		; r1  ← (r1) OR (r2)
// OR r, adr [,x]	; r   ← (r)  OR (実行アドレス)
func (p *Parser) ORStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()

	code.Op = 0x31
	code.Length = 2

	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.REGISTER:
		code.Op = 0x35
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
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// XORStatment  XOR Parser
// XOR r1, r2		; r1  ← (r1) XOR (r2)
// XOR r, adr [,x]	; r   ← (r)  XOR (実行アドレス)
func (p *Parser) XORStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x32
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.REGISTER:
		code.Op = 0x36
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
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// CPAStatment  CPA Parser
// CPA r1, r2		;
// CPA r, adr [,x]	;
func (p *Parser) CPAStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x40
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.REGISTER:
		code.Op = 0x44
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
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// CPLStatment  CPL Parser
// CPL r1, r2
// CPL r, adr [,x]
func (p *Parser) CPLStatment(code *opcode.Opcode) *opcode.Opcode {

	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.REGISTER) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・レジスタ・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()

	code.Op = 0x41
	code.Length = 2

	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.REGISTER:
		code.Op = 0x45
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
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}

// SLAStatment  SLA Parser
// SLA r, adr [,x]	;
func (p *Parser) SLAStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) && !p.peekTokenIs(token.HEX) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x50
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
			return nil
		}
		code.Addr = uint16(addr)
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	case token.LABEL:
		code.AddrLabel = p.curToken.Literal
		if !p.peekTokenIs(token.COMMA) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		p.nextToken()
		p.nextToken()
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.HEX:
		addr, err := p.hexToAddress(p.curToken.Literal)
		if err != nil {
			return nil
		}
		code.Addr = addr
		code, err = p.indexRegisterParse(code)
		if err != nil {
			return nil
		}
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8
	return code
}

// SRAStatment  SRA Parser
// SRA r, adr [,x]	;
func (p *Parser) SRAStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x51
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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

// SLLStatment  SLL Parser
// SLL r, adr [,x]	;
func (p *Parser) SLLStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x52
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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

// SRLStatment  SRL Parser
// SRL r, adr [,x]	;
func (p *Parser) SRLStatment(code *opcode.Opcode) *opcode.Opcode {
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	// Next Token is ','
	if !p.peekTokenIs(token.COMMA) {
		p.parserError(p.line, fmt.Sprintf("%qがありません。対象 : %q\n", ",", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	// Next Token is 'INT' or register or Label
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Op = 0x53
	code.Length = 2
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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

// JMIStatment Jump on Minus
// JMI adr, [,x];
func (p *Parser) JMIStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x61, Code: 0x6100, Length: 2, Label: code.Label}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// JNZStatment Jump on non Zero
// JNZ adr, [,x];
func (p *Parser) JNZStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x62, Code: 0x6200, Length: 2, Label: code.Label}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// JZEStatment Jump on Zero
// JZE adr, [,x];
func (p *Parser) JZEStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x63, Code: 0x6300, Length: 2, Label: code.Label}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// JUMPStatment Unconditional jump
// JUMP adr, [,x];
func (p *Parser) JUMPStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x64, Code: 0x6400, Length: 2, Label: code.Label}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// JPLStatment Jump on plus
// JPL adr, [,x];
func (p *Parser) JPLStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x65, Code: 0x6500, Length: 2, Label: code.Label}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// JOVStatment Jump on Overflow
// JOV adr, [,x];
func (p *Parser) JOVStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x66, Code: 0x6600, Length: 2, Label: code.Label}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// PUSHStatment PUSH
// PUSH adr, [,x];
func (p *Parser) PUSHStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x70, Code: 0x7000, Length: 2, Label: code.Label}
	if !p.peekTokenIs(token.INT) && !p.peekTokenIs(token.LABEL) {
		p.parserError(p.line, fmt.Sprintf("数値・ラベルではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	switch p.curToken.Type {
	case token.INT:
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			p.parserError(p.line, fmt.Sprintf("数値が適正ではありません。対象 : %q\n", p.curToken.Literal))
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
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}

// POPStatment PUSH
// POP adr, [,x];
func (p *Parser) POPStatment(code *opcode.Opcode) *opcode.Opcode {
	code = &opcode.Opcode{Op: 0x71, Code: 0x7100, Length: 1, Label: code.Label}
	if !p.peekTokenIs(token.REGISTER) {
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q\n", p.peekToken.Literal))
		return nil
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	return code
}

// hexToAddress #1000 → 4096(10)
func (p *Parser) hexToAddress(hex string) (uint16, error) {
	address, err := strconv.ParseUint(strings.Replace(hex, "#", "", 1), 16, 16)
	if err != nil {
		p.parserError(p.line, fmt.Sprintf("16進数数値が適正ではありません。\n#0000~#FFFFまで使用できます対象 : %q", hex))
		return 0, err
	}
	return uint16(address), nil
}

// indexRegisterParse
func (p *Parser) indexRegisterParse(code *opcode.Opcode) (*opcode.Opcode, error) {
	if !p.peekTokenIs(token.COMMA) {
		code.Code |= uint16(code.Op) << 8
		return code, nil
	}
	p.nextToken()
	if !p.peekTokenIs(token.REGISTER) {
		p.parserError(p.line, fmt.Sprintf("レジスタではありません。対象 : %q", p.peekToken.Literal))
		return nil, fmt.Errorf("Register Error")
	}
	p.nextToken()
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code, nil
}
