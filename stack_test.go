package goose_test

import (
	"fmt"
	"testing"

	"github.com/jncornett/goose"
)

func TestStackPos_AbsIndex(t *testing.T) {
	tests := []struct {
		Size             int
		Pos              goose.StackPos
		ExpectedAbsIndex uint
		ExpectError      bool
	}{
		{0, 0, 0, true},
		{0, 1, 0, true},
		{0, 2, 0, true},
		{0, -1, 0, true},
		{0, -2, 0, true},

		{1, 0, 0, true},
		{1, 1, 0, false},
		{1, 2, 0, true},
		{1, -1, 0, false},
		{1, -2, 0, true},

		{2, 0, 0, true},
		{2, 1, 0, false},
		{2, 2, 1, false},
		{2, -1, 1, false},
		{2, -2, 0, false},

		{2, -3, 0, true},
		{2, 3, 0, true},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.Size, "_", test.Pos), func(t *testing.T) {
			actual, err := test.Pos.AbsIndex(test.Size)
			haveError := err != nil
			if test.ExpectError != haveError {
				t.Errorf("expected error %v, got %v", test.ExpectError, haveError)
			}
			if test.ExpectedAbsIndex != actual {
				t.Errorf("expected %v, got %v", test.ExpectedAbsIndex, actual)
			}
		})
	}
}
