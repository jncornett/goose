package goose

import "errors"

var (
	ErrWrongType = errors.New("WrongType") // FIXME error should probably include expected value
)
