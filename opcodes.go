package goose

// OpCode represents an index into an op table.
type OpCode int

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
)
