package parser

import (
	"fmt"
	"strconv"

	"github.com/DJSIer/GCASL2/lexer"
	"github.com/DJSIer/GCASL2/opcode"
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

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
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
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
func (p *Parser) ParseProgram() []opcode.Opcode {
	Excode := []opcode.Opcode{}
	code := &opcode.Opcode{}
	switch p.curToken.Type {
	case token.LAD:
		code = p.LDAStatment()
	case token.LD:
		code = p.LDStatment()
	}
	if code != nil {
		Excode = append(Excode, *code)
	}
	return Excode
}
func (p *Parser) LDStatment() *opcode.Opcode {
	code := &opcode.Opcode{}
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
		code.Op = 0x10
		code.Length = 2
		addr, err := strconv.ParseUint(p.curToken.Literal, 0, 16)
		if err != nil {
			msg := fmt.Sprintf("parse error %q as Addr", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		code.Addr = uint16(addr)
		p.nextToken()
		if !p.expectPeek(token.REGISTER) {
			code.Code |= uint16(code.Op) << 8
			return code
		}
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	case token.REGISTER:
		code.Op = 0x14
		code.Length = 1
		code.Code |= uint16(registerNumber[p.curToken.Literal])
	default:
		code.Op = 0xFF
	}
	code.Code |= uint16(code.Op) << 8

	return code
}
func (p *Parser) LDAStatment() *opcode.Opcode {
	code := &opcode.Opcode{Code: 0x1200, Op: 0x12, Length: 2}
	if !p.expectPeek(token.REGISTER) {
		return nil
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal]) << 4
	p.nextToken()
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
	p.nextToken()
	if !p.expectPeek(token.REGISTER) {
		return code
	}
	code.Code |= uint16(registerNumber[p.curToken.Literal])
	return code
}
