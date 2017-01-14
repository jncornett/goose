package goose_test

import (
	"log"

	"github.com/jncornett/goose"
)

func ExampleState() {
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
