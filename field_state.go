package manta

import "sync"

type fieldState struct {
	state []interface{}
}

// Size classes for field state pools to optimize for common sizes
var (
	fieldStatePool8   = &sync.Pool{New: func() interface{} { return &fieldState{state: make([]interface{}, 8)} }}
	fieldStatePool16  = &sync.Pool{New: func() interface{} { return &fieldState{state: make([]interface{}, 16)} }}
	fieldStatePool32  = &sync.Pool{New: func() interface{} { return &fieldState{state: make([]interface{}, 32)} }}
	fieldStatePool64  = &sync.Pool{New: func() interface{} { return &fieldState{state: make([]interface{}, 64)} }}
	fieldStatePool128 = &sync.Pool{New: func() interface{} { return &fieldState{state: make([]interface{}, 128)} }}
)

func newFieldState() *fieldState {
	return getPooledFieldState(8)
}

func newFieldStateWithSize(size int) *fieldState {
	return getPooledFieldState(size)
}

func getPooledFieldState(minSize int) *fieldState {
	var fs *fieldState
	
	switch {
	case minSize <= 8:
		fs = fieldStatePool8.Get().(*fieldState)
	case minSize <= 16:
		fs = fieldStatePool16.Get().(*fieldState)
	case minSize <= 32:
		fs = fieldStatePool32.Get().(*fieldState)
	case minSize <= 64:
		fs = fieldStatePool64.Get().(*fieldState)
	case minSize <= 128:
		fs = fieldStatePool128.Get().(*fieldState)
	default:
		// For very large sizes, don't use pool
		return &fieldState{state: make([]interface{}, minSize)}
	}
	
	// Reset the field state for reuse
	fs.reset()
	return fs
}

func (s *fieldState) reset() {
	// Clear all values but keep the slice capacity
	for i := range s.state {
		s.state[i] = nil
	}
}

func (s *fieldState) release() {
	// Return to appropriate pool based on capacity
	cap := cap(s.state)
	switch {
	case cap <= 8:
		fieldStatePool8.Put(s)
	case cap <= 16:
		fieldStatePool16.Put(s)
	case cap <= 32:
		fieldStatePool32.Put(s)
	case cap <= 64:
		fieldStatePool64.Put(s)
	case cap <= 128:
		fieldStatePool128.Put(s)
	// Large field states are not pooled
	}
}

func (s *fieldState) releaseRecursive() {
	// Release any nested field states first
	for _, v := range s.state {
		if nested, ok := v.(*fieldState); ok {
			nested.releaseRecursive()
		}
	}
	// Reset this state and return to pool
	s.reset()
	s.release()
}

func (s *fieldState) get(fp *fieldPath) interface{} {
	x := s
	z := 0
	for i := 0; i <= fp.last; i++ {
		z = fp.path[i]
		if len(x.state) < z+2 {
			return nil
		}
		if i == fp.last {
			return x.state[z]
		}
		if _, ok := x.state[z].(*fieldState); !ok {
			return nil
		}
		x = x.state[z].(*fieldState)
	}
	return nil
}

func (s *fieldState) set(fp *fieldPath, v interface{}) {
	x := s
	z := 0
	for i := 0; i <= fp.last; i++ {
		z = fp.path[i]
		if y := len(x.state); y < z+2 {
			// Simple growth strategy: grow slice in place if possible
			newSize := max(z+2, y*2)
			newState := make([]interface{}, newSize)
			copy(newState, x.state)
			x.state = newState
		}
		if i == fp.last {
			if _, ok := x.state[z].(*fieldState); !ok {
				x.state[z] = v
			}
			return
		}
		if _, ok := x.state[z].(*fieldState); !ok {
			x.state[z] = newFieldState()
		}
		x = x.state[z].(*fieldState)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
