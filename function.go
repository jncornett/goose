package goose

type Function interface {
	Call(s *State) (returnValues []interface{}, err error)
}

type GoFunc func(*State) (returnValues []interface{}, err error)

func (f GoFunc) Call(s *State) (returnValues []interface{}, err error) {
	return f(s)
}
