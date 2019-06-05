package parser

import (
	"testing"

	"github.com/DJSIer/GCASL2/lexer"
)

func TestSymbol(t *testing.T) {
	tests := []struct {
		input         string
		expectedlabel string
	}{
		{"RAMEN LAD GR1,0", "RAMEN"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode := p.ParseProgram()
		if len(p.Errors()) != 0 {
			t.Fatal(p.errors)
		}

		op := opcode[0]
		if op.Label.Label != tt.expectedlabel {
			t.Fatalf("Label : %s now :%s", tt.expectedlabel, op.Label.Label)
		}
	}
}
func TestLDAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"RAMEN LAD GR1,0", 0x12, 0x1210, 0x0000},
		{"LAD GR3,65535", 0x12, 0x1230, 0xFFFF},
		{"LAD GR1,0,GR3", 0x12, 0x1213, 0x0000},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode := p.ParseProgram()
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
	}
}

func TestLDStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"LD GR2,0", 0x10, 0x1020, 0x0000},
		{"LD GR2,GR3", 0x14, 0x1423, 0x0000},
		{"LD GR1,0,GR3", 0x10, 0x1013, 0x0000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode := p.ParseProgram()
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
	}
}
