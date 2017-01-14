package goose

type Array interface {
	Append(interface{})
}

type defaultArray struct{}

func NewDefaultArray() Array {
	return &defaultArray{}
}

func (a *defaultArray) Append(v interface{}) {}
