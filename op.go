package goose

// OpCode represents an index into an op table.
type OpIndex int

// Op is a function that is triggered as the result of stepping through
// instruction table.
type OpFunc func(s *State)

// OpTable represents a mapping of OpCodes to Ops
type OpTable interface {
	Get(OpIndex) OpFunc
}

type Op OpIndex

type Code []Op
