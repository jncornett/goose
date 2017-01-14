package goose_test

import (
	"errors"
	"log"
	"strings"

	"github.com/jncornett/goose"
)

// This is example creates a new VM with default settings, loads a chunk,
// executes the chunk, and then cleans up (via the deferred Close())
func ExampleOne() {
	g := goose.NewState()
	defer g.Close()
	err := g.Load([]byte(`fmt.Println("hello, world")`))
	if err != nil {
		log.Fatal(err)
	}
	err = g.Call(0, 0)
	if err != nil {
		log.Fatal(err)
	}
}

// This is an example of how to create a Go function that can be called from
// inside of the VM.
// FIXME GoFuncs should be easier to implement, we should pass in a modified
// state that gets rid of much of the boilerplate
// FIXME Need to revisit use of []interface{} to pass return values
// FIXME Need to revisit calling methods on data types directly (e.g. Array.Append)
func ExampleTwo() {
	var _ goose.GoFunc = func(s *goose.State) ([]interface{}, error) {
		arg, err := s.Peek(0)
		if err != nil {
			return nil, err
		}
		str, ok := arg.(string)
		if !ok {
			return nil, errors.New("Wrong type, blah")
		}
		arg, err = s.Peek(1)
		if err != nil {
			return nil, err
		}
		sep, ok := arg.(string)
		if !ok {
			return nil, errors.New("Wrong type, blah")
		}
		var array goose.Array
		for _, v := range strings.Split(str, sep) {
			array.Append(v)
		}
		return []interface{}{array}, nil
	}
}
