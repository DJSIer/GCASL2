package opcode

import "github.com/DJSIer/GCASL2/symbol"

type Opcode struct {
	Code   uint16
	Addr   uint16
	Op     uint8
	Length int
	Label  symbol.Symbol
}

func (op *Opcode) GetAddr() uint16 {
	return op.Code
}
