package goose

// Op is a function that is triggered as the result of stepping through
// instruction table.
type Op func(s *State)

// OpTable represents a mapping of OpCodes to Ops
type OpTable interface {
	Get(OpCode) Op
}

type mapOpTable map[OpCode]Op

// NewMapOpTable creates a new map object that implements the OpTable interface.
func NewMapOpTable() OpTable {
	return &mapOpTable{
		OpRawGetTable:       opRawGetTable,
		OpRawSetTable:       opRawSetTable,
		OpSetTable:          opSetTable,
		OpGetTable:          opGetTable,
		OpPop:               opPop,
		OpCall:              opCall,
		OpLoadValue:         opLoadValue,
		OpLoadConstInt:      opLoadConstInt,
		OpLoadConstBool:     opLoadConstBool,
		OpLoadConstNil:      opLoadConstNil,
		OpAdd:               opAdd,
		OpSub:               opSub,
		OpMul:               opMul,
		OpDiv:               opDiv,
		OpBitwiseXor:        opBitwiseXor,
		OpBitwiseOr:         opBitwiseOr,
		OpBitwiseAnd:        opBitwiseAnd,
		OpBitwiseLeftshift:  opBitwiseLeftshift,
		OpBitwiseRightshift: opBitwiseRightshift,
		OpLessThan:          opLessThan,
		OpGreaterThan:       opGreaterThan,
		OpLessThanEqual:     opLessThanEqual,
		OpGreaterThanEqual:  opGreaterThanEqual,
		OpEqual:             opEqual,
		OpNotEqual:          opNotEqual,
		OpApproxEqual:       opApproxEqual,
		OpIdentical:         opIdentical,
		OpJump:              opJump,
		OpPanic:             opPanic,
	}
}

func (m mapOpTable) Get(code OpCode) Op {
	return m[code]
}
