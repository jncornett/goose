package goose_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jncornett/goose"
)

type AStruct struct {
	Key   string
	Value int
}

func TestPush(t *testing.T) {
	tests := []struct {
		Name  string
		Value interface{}
	}{
		{"nil", nil},
		{"int", 42},
		{"float", 3.14},
		{"string", "hello"},
		{"bool", true},
		{"AStruct", AStruct{"foo", 3}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			var s goose.SliceBasedStack
			err := s.Push(test.Value)
			if err != nil {
				t.Fatal(err)
			}
			if len(s.Slice) != 1 {
				t.Fatalf("expected size %v, got %v", 1, s.Size())
			}
			if !reflect.DeepEqual(test.Value, s.Slice[0]) {
				t.Fatalf("expected %v, got %v", test.Value, s.Slice[0])
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		OldSize   int
		PopAmount int
		NewSize   int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 0},
		{5, 2, 3},
		{5, 0, 5},
		{5, -3, 5},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.OldSize, "_", test.PopAmount), func(t *testing.T) {
			s := goose.SliceBasedStack{make([]interface{}, test.OldSize)}
			// pop
			s.Pop(test.PopAmount)
			if test.NewSize != len(s.Slice) {
				t.Errorf("expected size %v, got %v", test.NewSize, len(s.Slice))
			}
		})
	}
}

func TestGetAbsIndex(t *testing.T) {
	tests := []struct {
		Size          int
		RelativeIndex int
		ExpectError   bool
		AbsIndex      int
	}{
		{0, 0, true, 0},
		{0, 1, true, 0},
		{0, -1, true, 0},
		{5, -1, false, 4},
		{5, 0, false, 0},
		{5, 10, true, 0},
		{5, -5, false, 0},
		{5, -6, true, 0},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.Size, "_", test.RelativeIndex), func(t *testing.T) {
			s := goose.SliceBasedStack{make([]interface{}, test.Size)}
			index, err := s.GetAbsIndex(test.RelativeIndex)
			if test.ExpectError {
				if err == nil {
					t.Fatal("expected an error, got nothing")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if test.AbsIndex != index {
				t.Errorf("expected index %v, got %v", test.AbsIndex, index)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		StackSize     int
		PeekPos       int
		ExpectError   bool
		ExpectedValue int
	}{
		{0, 0, true, 0},
		{0, -1, true, 0},
		{1, 0, false, 1},
		{5, -1, false, 5},
		{5, 2, false, 3},
		{5, -5, false, 1},
		{5, -6, true, 0},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprint(test.StackSize, "_", test.PeekPos), func(t *testing.T) {
			s := goose.SliceBasedStack{makeIntRange(test.StackSize)}
			// peek
			val, err := s.Peek(test.PeekPos)
			if test.ExpectError {
				if err == nil {
					t.Fatal("expected an error, got nothing")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			v, ok := val.(int)
			if !ok {
				t.Fatalf("expected an int, got %v", val)
			}
			if test.ExpectedValue != v {
				t.Errorf("expected value %v, got %v", test.ExpectedValue, v)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	t.Run("OutOfRange", func(t *testing.T) {
		s := goose.SliceBasedStack{makeIntRange(5)}
		err := s.Copy(-10)
		if err == nil {
			t.Errorf("expected an error, got nothing")
		}
		checkIntRangeValid(t, s.Slice)
	})
	t.Run("InRange", func(t *testing.T) {
		s := goose.SliceBasedStack{makeIntRange(5)}
		err := s.Copy(-1)
		if err != nil {
			t.Error(err)
		}
		if 6 != len(s.Slice) {
			t.Fatalf("expected size %v, got %v", 6, len(s.Slice))
		}
		checkIntRangeValid(t, s.Slice, 5)
		if s.Slice[4] != s.Slice[5] {
			t.Errorf("expected %v, got %v", s.Slice[4], s.Slice[5])
		}
	})
}

func TestReplace(t *testing.T) {
	t.Run("OutOfRange", func(t *testing.T) {
		s := goose.SliceBasedStack{makeIntRange(5)}
		err := s.Replace(-10, 42)
		if err == nil {
			t.Errorf("expected an error, got nothing")
		}
		checkIntRangeValid(t, s.Slice)
	})
	t.Run("InRange", func(t *testing.T) {
		s := goose.SliceBasedStack{makeIntRange(5)}
		err := s.Replace(-1, 42)
		if err != nil {
			t.Error(err)
		}
		if 5 != len(s.Slice) {
			t.Fatalf("expected size %v, got %v", 5, len(s.Slice))
		}
		checkIntRangeValid(t, s.Slice, 4)
		if 42 != s.Slice[4] {
			t.Errorf("expected %v, got %v", 42, s.Slice[4])
		}
	})
}

func TestSwap(t *testing.T) {
	t.Run("OldOutOfRange", func(t *testing.T) {
		s := goose.SliceBasedStack{makeIntRange(5)}
		err := s.Swap(-10, 0)
		if err == nil {
			t.Errorf("expected an error, got nothing")
		}
		checkIntRangeValid(t, s.Slice)
	})
	t.Run("NewOutOfRange", func(t *testing.T) {
		s := goose.SliceBasedStack{makeIntRange(5)}
		err := s.Swap(0, -10)
		if err == nil {
			t.Errorf("expected an error, got nothing")
		}
		checkIntRangeValid(t, s.Slice)
	})
	t.Run("InRange", func(t *testing.T) {
		s := goose.SliceBasedStack{makeIntRange(5)}
		err := s.Swap(-1, 0)
		if err != nil {
			t.Error(err)
		}
		if 5 != len(s.Slice) {
			t.Fatalf("expected size %v, got %v", 5, len(s.Slice))
		}
		checkIntRangeValid(t, s.Slice, 0, 4)
		if 5 != s.Slice[0] {
			t.Errorf("expected %v, got %v", 5, s.Slice[0])
		}
		if 1 != s.Slice[4] {
			t.Errorf("expected %v, got %v", 1, s.Slice[4])
		}
	})
}

func TestSize(t *testing.T) {
	tests := []struct {
		Size  int
		Value []interface{}
	}{
		{0, nil},
		{0, makeIntRange(0)},
		{1, makeIntRange(1)},
		{5, makeIntRange(5)},
	}
	for _, test := range tests {
		test := test
		t.Run("", func(t *testing.T) {
			s := goose.SliceBasedStack{test.Value}
			if test.Size != s.Size() {
				t.Errorf("expected size %v, got %v", test.Size, s.Size())
			}
		})
	}
}

func makeIntRange(size int) []interface{} {
	slice := make([]interface{}, size)
	for i, _ := range slice {
		slice[i] = i + 1
	}
	return slice
}

func checkIntRangeValid(t *testing.T, slice []interface{}, ignores ...int) {
	ignoreSet := make(map[int]bool)
	for _, i := range ignores {
		ignoreSet[i] = true
	}
	for i, v := range slice {
		_, ok := ignoreSet[i]
		if ok {
			continue // in the ignore set, so ignore
		}
		if i+1 != v {
			t.Errorf("expected %v at index %v, got %v", i+1, i, v)
		}
	}
}
