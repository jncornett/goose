package goose

// SliceBasedStack implements the Stack interface
type SliceBasedStack struct {
	Slice []interface{}
}

// Push places v at the top of the stack.
// This implementation of Push never fails.
func (s *SliceBasedStack) Push(v interface{}) error {
	s.Slice = append(s.Slice, v)
	return nil
}

// Pop pops n values off of the top of the stack.
// If n is greater than the size of the stack, then the stack will be emptied.
// If n is less than zero, then Pop is a no-op.
func (s *SliceBasedStack) Pop(n int) {
	if n < 0 {
		n = 0
	}
	newSize := len(s.Slice) - n
	if newSize < 0 {
		newSize = 0
	}
	s.Slice = s.Slice[:newSize]
}

// Peek returns the value at pos in the stack.
// If pos is invalid, Peek returns an error.
func (s SliceBasedStack) Peek(pos StackPos) (interface{}, error) {
	index, err := pos.AbsIndex(len(s.Slice))
	if err != nil {
		return nil, err
	}
	val := s.Slice[index]
	return val, nil
}

// Copy copies the value at pos and pushes it onto the top of the stack>
// If pos is invalid, Copy returns an error.
func (s *SliceBasedStack) Copy(pos StackPos) error {
	index, err := pos.AbsIndex(len(s.Slice))
	if err != nil {
		return err
	}
	s.Slice = append(s.Slice, s.Slice[index])
	return nil
}

// Replace replaces the value at pos with v.
// If pos is invalid, Replace returns an error.
func (s SliceBasedStack) Replace(pos StackPos, v interface{}) error {
	index, err := pos.AbsIndex(len(s.Slice))
	if err != nil {
		return err
	}
	s.Slice[index] = v
	return nil
}

// Swap swaps the values at oldPos and newPos.
// If oldPos or newPos is invalid, Swap returns an error.
func (s SliceBasedStack) Swap(oldPos, newPos StackPos) error {
	oldIndex, err := oldPos.AbsIndex(len(s.Slice))
	if err != nil {
		return err
	}
	newIndex, err := newPos.AbsIndex(len(s.Slice))
	if err != nil {
		return err
	}
	tmp := s.Slice[oldIndex]
	s.Slice[oldIndex] = s.Slice[newIndex]
	s.Slice[newIndex] = tmp
	return nil
}

// Size returns the size of the stack
func (s SliceBasedStack) Size() int {
	return len(s.Slice)
}
