package goose

type Table interface {
	Get(key interface{}) (interface{}, error)
	Set(key, val interface{}) error
	Size() int
}

type DefaultTable struct{}

func NewDefaultTable() Table {
	return &DefaultTable{}
}

func (t *DefaultTable) Get(key interface{}) (interface{}, error) {
	return nil, nil
}

func (t *DefaultTable) Set(key, val interface{}) error {
	return nil
}

func (t *DefaultTable) Size() int {
	return 0
}
