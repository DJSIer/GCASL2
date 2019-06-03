package parser

import (
	"testing"

	"github.com/DJSIer/GCASL2/lexer"
)

func TestLDAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"LAD GR1,0", 0x12, 0x1210, 0x0000},
		{"LAD GR3,65535", 0x12, 0x1230, 0xFFFF},
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
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Addr)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedCode, op.Addr)
		}
	}
}
