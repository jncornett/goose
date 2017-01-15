package goose

const (
	OpNoop OpCode = iota
	OpRawGetTable
	OpRawSetTable
	OpSetTable
	OpGetTable
	OpPop
	OpCall
	OpLoadValue
	OpLoadConstInt
	OpLoadConstBool
	OpLoadConstNil
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpBitwiseXor
	OpBitwiseOr
	OpBitwiseAnd
	OpBitwiseLeftshift
	OpBitwiseRightshift
	OpLessThan
	OpGreaterThan
	OpLessThanEqual
	OpGreaterThanEqual
	OpEqual
	OpNotEqual
	OpApproxEqual
	OpIdentical
	OpJump
	OpPanic
	OpMax // Must be the last identifier
)

func opRawGetTable(s *State)       {}
func opRawSetTable(s *State)       {}
func opSetTable(s *State)          {}
func opGetTable(s *State)          {}
func opPop(s *State)               {}
func opCall(s *State)              {}
func opLoadValue(s *State)         {}
func opLoadConstInt(s *State)      {}
func opLoadConstBool(s *State)     {}
func opLoadConstNil(s *State)      {}
func opAdd(s *State)               {}
func opSub(s *State)               {}
func opMul(s *State)               {}
func opDiv(s *State)               {}
func opBitwiseXor(s *State)        {}
func opBitwiseOr(s *State)         {}
func opBitwiseAnd(s *State)        {}
func opBitwiseLeftshift(s *State)  {}
func opBitwiseRightshift(s *State) {}
func opLessThan(s *State)          {}
func opGreaterThan(s *State)       {}
func opLessThanEqual(s *State)     {}
func opGreaterThanEqual(s *State)  {}
func opEqual(s *State)             {}
func opNotEqual(s *State)          {}
func opApproxEqual(s *State)       {}
func opIdentical(s *State)         {}
func opJump(s *State)              {}
func opPanic(s *State)             {}

type mapOpTable map[OpCode]Op

// NewMapOpTable creates a new map object that implements the OpTable interface.
// NOTE Until the op set becomes more stabilized, keeping the op table as a map
// makes sense. In the future, we may switch to a faster implementation.
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
