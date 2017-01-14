package goose

type Function interface {
	Call(s *State) error
}

type GoFunc func(*State) error

func (f GoFunc) Call(s *State) error {
	return f(s)
}
