package goose

import "fmt"

// SliceBasedStack implements the Stack interface
type SliceBasedStack struct {
	Slice []interface{}
}

// Push places v at the top of the stack
func (s *SliceBasedStack) Push(v interface{}) error {
	s.Slice = append(s.Slice, v)
	return nil
}

// Pop pops n values off of the top of the stack
// If n is greater than the size of the stack, then the stack will be emptied.
// If n is less than zero, then Pop is a no-op
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

func (s SliceBasedStack) Peek(pos int) (interface{}, error) {
	index, err := s.GetAbsIndex(pos)
	if err != nil {
		return nil, err
	}
	val := s.Slice[index]
	return val, nil
}

func (s *SliceBasedStack) Copy(pos int) error {
	index, err := s.GetAbsIndex(pos)
	if err != nil {
		return err
	}
	s.Slice = append(s.Slice, s.Slice[index])
	return nil
}

func (s SliceBasedStack) Replace(pos int, v interface{}) error {
	index, err := s.GetAbsIndex(pos)
	if err != nil {
		return err
	}
	s.Slice[index] = v
	return nil
}

func (s SliceBasedStack) Swap(oldPos, newPos int) error {
	oldIndex, err := s.GetAbsIndex(oldPos)
	if err != nil {
		return err
	}
	newIndex, err := s.GetAbsIndex(newPos)
	if err != nil {
		return err
	}
	tmp := s.Slice[oldIndex]
	s.Slice[oldIndex] = s.Slice[newIndex]
	s.Slice[newIndex] = tmp
	return nil
}

func (s SliceBasedStack) Size() int {
	return len(s.Slice)
}

func (s SliceBasedStack) GetAbsIndex(pos int) (int, error) {
	var index int
	if pos < 0 {
		index = len(s.Slice) + pos
	} else {
		index = pos
	}
	if index < 0 || index >= len(s.Slice) {
		return 0, errOutOfRange(len(s.Slice), pos)
	}
	return index, nil
}

func errOutOfRange(size, index int) error {
	return fmt.Errorf("Index out of range (max %v): %v", size, index)
}
