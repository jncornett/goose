package goose

// NOTE implementers should ensure that all implemented operations are atomic
type Stack interface {
	Push(v interface{}) error
	Pop(n int) // NOTE pop must not fail
	Peek(pos int) (interface{}, error)
	Replace(pos int, v interface{}) error
	Copy(pos int) error
	Swap(oldPos, newPos int) error
	Size() int
}
