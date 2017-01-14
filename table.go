package goose

type Table interface {
	Get(key interface{}) (interface{}, error)
	Set(key, val interface{}) error
	Size() int
}

type defaultTable struct{}

func NewDefaultTable() Table {
	return &defaultTable{}
}

func (t *defaultTable) Get(key interface{}) (interface{}, error) {
	return nil, nil
}

func (t *defaultTable) Set(key, val interface{}) error {
	return nil
}

func (t *defaultTable) Size() int {
	return 0
}
