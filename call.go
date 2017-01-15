package goose

type Symbol string

type Frame interface {
	Step() bool
	Locals() Table
}

type CallStack struct {
	Frames []Frame
}

func (s CallStack) Lookup(sym Symbol) interface{} {
	return nil
}

func (s CallStack) Assign(sym Symbol, val interface{}) {
}

func (s CallStack) Step() bool {
	return false
}

func (s CallStack) Push(code Code, locals Table) {
}
