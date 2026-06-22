package manta

import "math"

// cellKind tags the typed value held by a cell.
type cellKind uint8

const (
	cellEmpty cellKind = iota
	cellFloat32
	cellInt32
	cellUint32
	cellUint64 // value held inline in num (fits in 32 bits), reproduced as uint64
	cellBool
	cellRef // ref holds a string, []float32, []byte, or a boxed 64-bit integer
	cellSub // ref holds a *fieldState
)

// cell is a tagged union holding one decoded field value inline. The common
// scalars (float32, int32, uint32, bool, and 32-bit-range uint64 handles) live
// in num without allocating. Reference values (strings, vectors, binary blobs)
// and the rare genuinely-64-bit integers (CStrongHandle, fixed64, int64) use the
// interface field. Keeping num a uint32 holds the cell to 24 bytes, which keeps
// the per-entity state arrays (cloned from the baseline and grown on write)
// compact. Scalars are boxed into interface{} lazily, on read through iface()
// (e.g. Entity.Get), which is far rarer than the per-field-per-tick write path.
type cell struct {
	ref  interface{}
	num  uint32
	kind cellKind
}

func floatCell(f float32) cell   { return cell{num: math.Float32bits(f), kind: cellFloat32} }
func int32Cell(v int32) cell     { return cell{num: uint32(v), kind: cellInt32} }
func uint32Cell(v uint32) cell   { return cell{num: v, kind: cellUint32} }
func refCell(v interface{}) cell { return cell{ref: v, kind: cellRef} }
func subCell(s *fieldState) cell { return cell{ref: s, kind: cellSub} }

// smallUint64Cell stores a uint64 whose value is known to fit in 32 bits inline,
// reproduced as a uint64 on read. Used for varint handle types whose wire value
// cannot exceed 32 bits. Genuinely 64-bit values must use refCell instead.
func smallUint64Cell(v uint64) cell { return cell{num: uint32(v), kind: cellUint64} }

func boolCell(b bool) cell {
	c := cell{kind: cellBool}
	if b {
		c.num = 1
	}
	return c
}

func (c cell) isEmpty() bool { return c.kind == cellEmpty }

// sub returns the nested fieldState if this cell holds one.
func (c cell) sub() (*fieldState, bool) {
	if c.kind == cellSub {
		return c.ref.(*fieldState), true
	}
	return nil, false
}

// asFloat32 returns the float32 held by a scalar float cell. It is used by the
// vector decoders, which assemble a []float32 from per-component float decoders.
func (c cell) asFloat32() float32 { return math.Float32frombits(c.num) }

// iface boxes the cell value into an interface{} for the public Get API,
// reproducing the exact dynamic type the decoder originally produced. This is
// the only place a stored scalar is heap-boxed.
func (c cell) iface() interface{} {
	switch c.kind {
	case cellFloat32:
		return math.Float32frombits(c.num)
	case cellInt32:
		return int32(c.num)
	case cellUint32:
		return c.num
	case cellUint64:
		return uint64(c.num)
	case cellBool:
		return c.num != 0
	case cellRef, cellSub:
		return c.ref
	}
	return nil
}

type fieldState struct {
	state []cell
}

func newFieldState() *fieldState {
	return &fieldState{
		state: make([]cell, 8),
	}
}

func (s *fieldState) get(fp *fieldPath) interface{} {
	x := s
	for i := 0; i <= fp.last; i++ {
		z := fp.path[i]
		if len(x.state) < z+2 {
			return nil
		}
		if i == fp.last {
			return x.state[z].iface()
		}
		sub, ok := x.state[z].sub()
		if !ok {
			return nil
		}
		x = sub
	}
	return nil
}

func (s *fieldState) set(fp *fieldPath, c cell) {
	x := s
	for i := 0; i <= fp.last; i++ {
		z := fp.path[i]
		if y := len(x.state); y < z+2 {
			grown := make([]cell, max(z+2, y*2))
			copy(grown, x.state)
			x.state = grown
		}
		if i == fp.last {
			// Do not overwrite an existing sub-table with a leaf value.
			if x.state[z].kind != cellSub {
				x.state[z] = c
			}
			return
		}
		if x.state[z].kind != cellSub {
			x.state[z] = subCell(newFieldState())
		}
		x, _ = x.state[z].sub()
	}
}

// clone returns a deep copy of the fieldState so an entity cloned from a cached
// class baseline can be mutated independently of the template and its siblings.
// Nested fieldStates are copied recursively. Mutable leaf values (vector
// []float32 and binary-blob []byte) are also deep-copied, since a caller may
// mutate a slice returned from Entity.Get/Map; strings and boxed integers are
// immutable and shared by value.
func (s *fieldState) clone() *fieldState {
	c := &fieldState{state: make([]cell, len(s.state))}
	copy(c.state, s.state)
	for i := range c.state {
		switch v := c.state[i].ref.(type) {
		case *fieldState:
			c.state[i].ref = v.clone()
		case []float32:
			c.state[i].ref = append([]float32(nil), v...)
		case []byte:
			c.state[i].ref = append([]byte(nil), v...)
		}
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
