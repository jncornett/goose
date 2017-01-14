package goose

import (
	"bytes"
	"io"
)

type State struct {
	stack   Stack
	globals Table
}

func NewState() *State {
	return &State{
		&SliceBasedStack{},
		NewDefaultTable(),
	}
}

func (s *State) Load(b []byte) error {
	return s.LoadFromReader(bytes.NewReader(b))
}

func (s *State) LoadFromReader(r io.Reader) error {
	return nil
}

func (s *State) Call(nargs, nret int) error {
	return nil
}

func (s *State) Push(v interface{}) error {
	// FIXME do we want to do type-checking/conversion at this level?
	// or do we want to push that responsibility to the stack (pun intended)
	// +the stack should be as dumb as possible
	return s.stack.Push(v)
}

func (s *State) PushCopy(pos int) error {
	return s.stack.Copy(pos)
}

func (s *State) Peek(pos int) (interface{}, error) {
	return s.stack.Peek(pos)
}

func (s *State) Pop(n int) {
	s.stack.Pop(n)
}

func (s *State) StackSize() int {
	return s.stack.Size()
}

func (s *State) RawGetTable(tablePos int) error {
	obj, err := s.stack.Peek(tablePos)
	if err != nil {
		return err
	}
	t, ok := obj.(Table)
	if !ok {
		// FIXME we may want to include the expected type
		return ErrWrongType
	}
	key, err := s.stack.Peek(-1)
	if err != nil {
		return err
	}
	val, err := t.Get(key)
	if err != nil {
		return err
	}
	return s.stack.Replace(-1, val)
}

func (s *State) RawSetTable(tablePos int) error {
	obj, err := s.stack.Peek(tablePos)
	if err != nil {
		return err
	}
	t, ok := obj.(Table)
	if !ok {
		// FIXME we may want to include the expected type
		return ErrWrongType
	}
	key, err := s.stack.Peek(-2)
	if err != nil {
		return err
	}
	val, err := s.stack.Peek(-1)
	if err != nil {
		return err
	}
	err = t.Set(key, val)
	if err != nil {
		return err
	}
	s.stack.Pop(2)
	return nil
}

func (s *State) GetTable(tablePos, keyPos, valuePos int) error {
	return nil // TODO implement
}

func (s *State) SetTable(tablePos, keyPos, valuePos int) error {
	return nil // TODO implement
}

func (s *State) NewTable() error {
	return s.stack.Push(NewDefaultTable())
}

func (s *State) Close() error {
	// any cleanup code goes here
	return nil
}
