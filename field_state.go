package manta

import (
	"sync"
)

var fieldStatePool = &sync.Pool{
	New: func() interface{} {
		return &fieldState{
			values: []interface{}{},
		}
	},
}

type fieldState struct {
	values []interface{}
	size   int
}

func newFieldState() *fieldState {
	return fieldStatePool.Get().(*fieldState)
}

func (s *fieldState) release() {
	s.values = []interface{}{}
	fieldStatePool.Put(s)
}

func (s *fieldState) grow(n int) {
	if s.size < n {
		x := make([]interface{}, n*2)
		copy(x, s.values)
		s.values = x
		s.size = n * 2
	}
}

func (s *fieldState) get(path ...int) interface{} {
	return s.values[path[0]]
}

func (s *fieldState) set(v interface{}, path ...int) {
	s.grow(path[0] + 1)
	s.values[path[0]] = v
}
