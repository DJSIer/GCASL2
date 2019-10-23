package token

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	REGISTER  = "REG"
	LABEL     = "LABEL"
	START     = "START"
	END       = "END"
	DS        = "DS"
	DC        = "DC"
	IN        = "IN"
	OUT       = "OUT"
	RPUSH     = "RPUSH"
	RPOP      = "RPOP"
	EOF       = "EOF"
	INT       = "INT"
	HEX       = "HEX"
	STRING    = "STRING"
	SHARP     = "#"
	COMMA     = ","
	SEMICOLON = ":"
	LD        = "LD"
	ST        = "ST"
	LAD       = "LAD"
	ADDA      = "ADDA"
	ADDL      = "ADDL"
	SUBA      = "SUBA"
	SUBL      = "SUBL"
	AND       = "AND"
	OR        = "OR"
	XOR       = "XOR"
	CPA       = "CPA"
	CPL       = "CPL"
	SLA       = "SLA"
	SRA       = "SRA"
	SLL       = "SLL"
	SRL       = "SRL"
	JPL       = "JPL"
	JMI       = "JMI"
	JNZ       = "JNZ"
	JZE       = "JZE"
	JOV       = "JOV"
	JUMP      = "JUMP"
	PUSH      = "PUSH"
	POP       = "POP"
	CALL      = "CALL"
	RET       = "RET"
	SVC       = "SVC"
	NOP       = "NOP"
	ADLI      = "DC"
)

type Token struct {
	Type    TokenType `json:"-"`
	Literal string
}

var caslLetter = map[byte]uint8{
	' ': 0x02, '!': 0x12, '"': 0x22, '#': 0x32,
	'$': 0x42, '%': 0x52, '&': 0x62, '\'': 0x72,
	'(': 0x82, ')': 0x92, '*': 0xA2, '+': 0xB2,
	',': 0xC2, '-': 0xD2, '.': 0xE2, '/': 0xF2,
	'0': 0x03, '1': 0x13, '2': 0x23, '3': 0x33,
	'4': 0x43, '5': 0x53, '6': 0x63, '7': 0x73,
	'8': 0x83, '9': 0x93, ':': 0xA3, ';': 0xB3,
	'<': 0xC3, '=': 0xD3, '>': 0xE3, '?': 0xF3,
	'@': 0x04, 'A': 0x14, 'B': 0x24, 'C': 0x34,
	'D': 0x44, 'E': 0x54, 'F': 0x64, 'G': 0x74,
	'H': 0x84, 'I': 0x94, 'J': 0xA4, 'K': 0xB4,
	'L': 0xC4, 'M': 0xD4, 'N': 0xE4, 'O': 0xF4,
	'P': 0x05, 'Q': 0x15, 'R': 0x25, 'S': 0x35,
	'T': 0x45, 'U': 0x55, 'V': 0x65, 'W': 0x75,
	'X': 0x85, 'Y': 0x95, 'Z': 0xA5, '[': 0xB5,
	'\\': 0xC5, ']': 0xD5, '^': 0xE5, '_': 0xF5,
	'`': 0x06, 'a': 0x16, 'b': 0x26, 'c': 0x36,
	'd': 0x46, 'e': 0x56, 'f': 0x66, 'g': 0x76,
	'h': 0x86, 'i': 0x96, 'j': 0xA6, 'k': 0xB6,
	'l': 0xC6, 'm': 0xD6, 'n': 0xE6, 'o': 0xF6,
	'p': 0x07, 'q': 0x17, 'r': 0x27, 's': 0x37,
	't': 0x47, 'u': 0x57, 'v': 0x67, 'w': 0x77,
	'x': 0x87, 'y': 0x97, 'z': 0xA7, '{': 0xB7,
	'|': 0xC7, '}': 0xD7, '~': 0xE7,
}
var keywords = map[string]TokenType{
	"GR0":   REGISTER,
	"GR1":   REGISTER,
	"GR2":   REGISTER,
	"GR3":   REGISTER,
	"GR4":   REGISTER,
	"GR5":   REGISTER,
	"GR6":   REGISTER,
	"GR7":   REGISTER,
	"START": START,
	"END":   END,
	"DS":    DS,
	"DC":    DC,
	"IN":    IN,
	"OUT":   OUT,
	"RPUSH": RPUSH,
	"RPOP":  RPOP,
	"LD":    LD,
	"ST":    ST,
	"LAD":   LAD,
	"ADDA":  ADDA,
	"ADDL":  ADDL,
	"SUBA":  SUBA,
	"SUBL":  SUBL,
	"AND":   AND,
	"OR":    OR,
	"XOR":   XOR,
	"CPA":   CPA,
	"CPL":   CPL,
	"SLA":   SLA,
	"SRA":   SRA,
	"SLL":   SLL,
	"SRL":   SRL,
	"JPL":   JPL,
	"JMI":   JMI,
	"JNZ":   JNZ,
	"JZE":   JZE,
	"JOV":   JOV,
	"JUMP":  JUMP,
	"PUSH":  PUSH,
	"POP":   POP,
	"CALL":  CALL,
	"RET":   RET,
	"SVC":   SVC,
	"NOP":   NOP,
}

func LookupLetter(ch byte) (uint8, bool) {
	c, ok := caslLetter[ch]
	return c, ok
}
func LookupInst(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return LABEL
}
