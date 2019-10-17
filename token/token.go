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

func LookupInst(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return LABEL
}
