package goose

// OpCode represents an index into an op table.
type OpCode int

// Op is a function that is triggered as the result of stepping through
// instruction table.
type Op func(s *State)

// OpTable represents a mapping of OpCodes to Ops
type OpTable interface {
	Get(OpCode) Op
}
