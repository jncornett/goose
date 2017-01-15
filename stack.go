package goose

import "errors"

// StackPos denotes a relative index into a stack.
// The Stack interface is indexed by 1.
// When StackPos is positive it denotes the 1-based index into the Stack.
// When StackPos is negative it denotes the 1-based distance from the top of the Stack.
// When StackPos is zero, it denotes the invalid Stack index.
type StackPos int

func (p StackPos) AbsIndex(size int) (uint, error) {
	if size < 0 {
		return 0, errors.New("Invalid stack size")
	}
	var index int
	if p < 0 {
		index = int(size) + int(p)
	} else {
		index = int(p) - 1
	}
	if index < 0 || index >= size {
		return 0, errors.New("Invalid index")
	}
	return uint(index), nil
}

// NOTE implementers should ensure that all implemented operations are atomic
type Stack interface {
	Push(v interface{}) error
	Pop(n int) // NOTE pop must not fail
	Peek(pos StackPos) (interface{}, error)
	Replace(pos StackPos, v interface{}) error
	Copy(pos StackPos) error
	Swap(oldPos, newPos StackPos) error
	Size() int
}
