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
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Label.Label != tt.expectedlabel {
			t.Fatalf("Label : %s now :%s", tt.expectedlabel, op.Label.Label)
		}
	}
}
func TestLDCtatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"DC 0", 0x00, 0x0000, 0x0000},
		{"DC 10", 0x00, 0x0000, 0x000A},
		{"DC #1000", 0x00, 0x0000, 0x1000},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestLDAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"RAMEN LAD GR1,0", 0x12, 0x1210, 0x0000},
		{"LAD GR3,10", 0x12, 0x1230, 0x000A},
		{"LAD GR1,0,GR3", 0x12, 0x1213, 0x0000},
		{"LAD GR1,#0010,GR3", 0x12, 0x1213, 0x010},
		{"LAD GR1,GO,GR3", 0x12, 0x1213, 0x0000},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
		{"LD GR2,GO", 0x10, 0x1020, 0x0000},
		{"LD GR2,1234", 0x10, 0x1020, 0x04d2},
		{"LD GR1,#0010,GR3", 0x10, 0x1013, 0x010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestSTStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"ST GR2,0", 0x11, 0x1120, 0x0000},
		{"ST GR1,0,GR3", 0x11, 0x1113, 0x0000},
		{"ST GR7,0,GR7", 0x11, 0x1177, 0x0000},
		{"ST GR1,#0010,GR3", 0x11, 0x1113, 0x010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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

func TestADDAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"ADDA GR2,0", 0x20, 0x2020, 0x0000},
		{"ADDA GR1,0,GR3", 0x20, 0x2013, 0x0000},
		{"ADDA GR7,GR7", 0x24, 0x2477, 0x0000},
		{"ADDA GR1,#0010,GR3", 0x20, 0x2013, 0x010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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

func TestSUBAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"SUBA GR2,0", 0x21, 0x2120, 0x0000},
		{"SUBA GR1,0,GR3", 0x21, 0x2113, 0x0000},
		{"SUBA GR7,GR7", 0x25, 0x2577, 0x0000},
		{"SUBA GR1,#0010,GR3", 0x21, 0x2113, 0x010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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

func TestADDLStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"ADDL GR2,0", 0x22, 0x2220, 0x0000},
		{"ADDL GR1,0,GR3", 0x22, 0x2213, 0x0000},
		{"ADDL GR7,GR7", 0x26, 0x2677, 0x0000},
		{"ADDL GR1,#0010,GR3", 0x22, 0x2213, 0x010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestSUBLStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"SUBL GR2,0", 0x23, 0x2320, 0x0000},
		{"SUBL GR1,0,GR3", 0x23, 0x2313, 0x0000},
		{"SUBL GR7,GR7", 0x27, 0x2777, 0x0000},
		{"SUBL GR1,#0010,GR3", 0x23, 0x2313, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestANDStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"AND GR2,0", 0x30, 0x3020, 0x0000},
		{"AND GR1,0,GR3", 0x30, 0x3013, 0x0000},
		{"AND GR7,GR7", 0x34, 0x3477, 0x0000},
		{"AND GR1,#0010,GR3", 0x30, 0x3013, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestORStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"OR GR2,0", 0x31, 0x3120, 0x0000},
		{"OR GR1,0,GR3", 0x31, 0x3113, 0x0000},
		{"OR GR7,GR7", 0x35, 0x3577, 0x0000},
		{"OR GR1,#0010,GR3", 0x31, 0x3113, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestXORStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"XOR GR2,0", 0x32, 0x3220, 0x0000},
		{"XOR GR1,0,GR3", 0x32, 0x3213, 0x0000},
		{"XOR GR7,GR7", 0x36, 0x3677, 0x0000},
		{"XOR GR1,#0010,GR3", 0x32, 0x3213, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestCPAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"CPA GR2,0", 0x40, 0x4020, 0x0000},
		{"CPA GR1,10,GR3", 0x40, 0x4013, 0x000A},
		{"CPA GR7,GR7", 0x44, 0x4477, 0x0000},
		{"CPA GR1,#0010,GR3", 0x40, 0x4013, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestCPLStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"CPL GR2,0", 0x41, 0x4120, 0x0000},
		{"CPL GR1,A,GR3", 0x41, 0x4113, 0x0000},
		{"CPL GR7,GR7", 0x45, 0x4577, 0x0000},
		{"CPL GR1,#0010,GR3", 0x41, 0x4113, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestSLAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"SLA GR2,0", 0x50, 0x5020, 0x0000},
		{"SLA GR1,0,GR3", 0x50, 0x5013, 0x0000},
		{"SLA GR1,#0010,GR3", 0x50, 0x5013, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestSRAStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"SRA GR2,1000", 0x51, 0x5120, 0x03e8},
		{"SRA GR1,0,GR3", 0x51, 0x5113, 0x0000},
		{"SRA GR1,#0010,GR3", 0x51, 0x5113, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestSLLStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"SLL GR2,1000", 0x52, 0x5220, 0x03e8},
		{"SLL GR1,0,GR3", 0x52, 0x5213, 0x0000},
		{"SLL GR1,#0010,GR3", 0x52, 0x5213, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestSRLStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"SRL GR2,1000", 0x53, 0x5320, 0x03e8},
		{"SRL GR1,0,GR3", 0x53, 0x5313, 0x0000},
		{"SRL GR1,#0010,GR3", 0x53, 0x5313, 0x0010},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
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
func TestJMIStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"JMI 1000", 0x61, 0x6100, 0x03e8},
		{"JMI #1000, GR3", 0x61, 0x6103, 0x1000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
func TestJNZStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"JNZ 1000", 0x62, 0x6200, 0x03e8},
		{"JNZ #1000, GR3", 0x62, 0x6203, 0x1000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
func TestJZEStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"JZE 1000", 0x63, 0x6300, 0x03e8},
		{"JZE #1000, GR3", 0x63, 0x6303, 0x1000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
func TestJUMPStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"JUMP 1000", 0x64, 0x6400, 0x03e8},
		{"JUMP #1000, GR3", 0x64, 0x6403, 0x1000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
func TestJPLStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"JPL 1000", 0x65, 0x6500, 0x03e8},
		{"JPL #1000, GR3", 0x65, 0x6503, 0x1000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(p.Errors())
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
func TestJOVStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"JOV 1000", 0x66, 0x6600, 0x03e8},
		{"JOV #1000, GR3", 0x66, 0x6603, 0x1000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(p.Errors())
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
func TestPUSHStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"PUSH 1000", 0x70, 0x7000, 0x03e8},
		{"PUSH #1000, GR3", 0x70, 0x7003, 0x1000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(p.Errors())
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
func TestPOPStatment(t *testing.T) {
	tests := []struct {
		input        string
		expectedOp   uint8
		expectedCode uint16
		expectedAddr uint16
	}{
		{"POP GR1", 0x71, 0x7110, 0x0000},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		opcode, err := p.ParseProgram()
		if err != nil {
			t.Fatal(p.Errors())
			t.Fatal(err)
		}
		op := opcode[0]
		if op.Op != tt.expectedOp {
			t.Fatalf("Opcode : 0x%02x Now : 0x%02x", tt.expectedOp, op.Op)
		}
		if op.Addr != tt.expectedAddr {
			t.Fatalf("Addr : 0x%04x Now : 0x%04x", tt.expectedAddr, op.Addr)
		}
		if op.Code != tt.expectedCode {
			t.Fatalf("code : 0x%04x Now : 0x%04x", tt.expectedCode, op.Code)
		}
	}
}
