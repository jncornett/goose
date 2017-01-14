package goose

import (
	"bytes"
	"io"
)

type State struct {
}

func NewState() *State {
	return nil
}

func (s *State) Load(b []byte) error {
	return s.LoadFromReader(bytes.NewReader(b))
}

func (s *State) LoadFromReader(r io.Reader) error {
	return nil
}
