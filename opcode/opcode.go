package opcode

type Opcode struct {
	Code   uint16
	Addr   uint16
	Op     uint8
	Length int
}

func (op *Opcode) GetAddr() uint16 {
	return op.Code
}
