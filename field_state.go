package manta

import (
	"math"
	"slices"
	"unsafe"
)

// cellKind tags the typed value held by a cell.
type cellKind uint8

const (
	cellEmpty cellKind = iota
	cellFloat32
	cellInt32
	cellUint32
	cellUint64 // value held inline in num (fits in 32 bits), reproduced as uint64
	cellInt64  // value held inline in num (fits in int32), reproduced as int64
	cellBool
	cellString    // ptr/num are the string's data and length
	cellFloats    // ptr/num are a []float32's data and length
	cellBytes     // ptr/num are a []byte's data and length
	cellUint64Box // ptr is a *uint64 holding a genuinely 64-bit value
	cellInt64Box  // ptr is a *int64 holding a genuinely 64-bit value
	cellSub       // ptr is a *fieldState
)

// cell is a tagged union holding one decoded field value. The common scalars
// (float32, int32, uint32, bool, and 32-bit-range uint64/int64 values) live in
// num without allocating. Strings, vectors and binary blobs keep just their
// data pointer in ptr and their length in num, so a decode allocates only the
// data itself, not an interface box for it; the rare genuinely-64-bit integers
// (CStrongHandle, fixed64, steam IDs) point at a heap copy; nested tables point
// at their *fieldState. kind says which, so the value is rebuilt with its exact
// original type. This holds the cell to 16 bytes, which keeps the per-entity
// state arrays (cloned from the baseline and grown on write) compact. Values are
// boxed into interface{} lazily, on read through iface() (e.g. Entity.Get),
// which is far rarer than the per-field-per-tick write path.
type cell struct {
	ptr  unsafe.Pointer
	num  uint32
	kind cellKind

	// shared marks a vector, blob or nested state owned by a class baseline
	// template, and so shared with every entity cloned from it (see share).
	// iface hands out copies of shared vectors and blobs, and set copies a
	// shared nested state before writing into it.
	shared bool
}

func floatCell(f float32) cell   { return cell{num: math.Float32bits(f), kind: cellFloat32} }
func int32Cell(v int32) cell     { return cell{num: uint32(v), kind: cellInt32} }
func uint32Cell(v uint32) cell   { return cell{num: v, kind: cellUint32} }
func subCell(s *fieldState) cell { return cell{ptr: unsafe.Pointer(s), kind: cellSub} }

func stringCell(v string) cell {
	return cell{ptr: unsafe.Pointer(unsafe.StringData(v)), num: uint32(len(v)), kind: cellString}
}

func floatsCell(v []float32) cell {
	return cell{ptr: unsafe.Pointer(unsafe.SliceData(v)), num: uint32(len(v)), kind: cellFloats}
}

func bytesCell(v []byte) cell {
	return cell{ptr: unsafe.Pointer(unsafe.SliceData(v)), num: uint32(len(v)), kind: cellBytes}
}

// smallUint64Cell stores a uint64 whose value is known to fit in 32 bits inline,
// reproduced as a uint64 on read. Used for varint handle types whose wire value
// cannot exceed 32 bits. Values that may be wider must use uint64Cell.
func smallUint64Cell(v uint64) cell { return cell{num: uint32(v), kind: cellUint64} }

// uint64Cell stores a uint64, inline when it fits in 32 bits and otherwise in
// a heap copy. Either way it reads back as a uint64.
func uint64Cell(v uint64) cell {
	if v <= math.MaxUint32 {
		return smallUint64Cell(v)
	}
	p := new(uint64)
	*p = v
	return cell{ptr: unsafe.Pointer(p), kind: cellUint64Box}
}

// int64Cell stores an int64, inline when it fits in an int32 and otherwise in
// a heap copy. Either way it reads back as an int64.
func int64Cell(v int64) cell {
	if v >= math.MinInt32 && v <= math.MaxInt32 {
		return cell{num: uint32(int32(v)), kind: cellInt64}
	}
	p := new(int64)
	*p = v
	return cell{ptr: unsafe.Pointer(p), kind: cellInt64Box}
}

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
		return (*fieldState)(c.ptr), true
	}
	return nil, false
}

// asFloat32 returns the float32 held by a scalar float cell. It is used by the
// vector decoders, which assemble a []float32 from per-component float decoders.
func (c cell) asFloat32() float32 { return math.Float32frombits(c.num) }

// floats returns the []float32 held by a cellFloats cell.
func (c cell) floats() []float32 { return unsafe.Slice((*float32)(c.ptr), c.num) }

// bytes returns the []byte held by a cellBytes cell.
func (c cell) bytes() []byte { return unsafe.Slice((*byte)(c.ptr), c.num) }

// iface boxes the cell value into an interface{} for the public Get API,
// reproducing the exact dynamic type the decoder originally produced. This is
// the only place a stored value is heap-boxed.
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
	case cellInt64:
		return int64(int32(c.num))
	case cellBool:
		return c.num != 0
	case cellString:
		return unsafe.String((*byte)(c.ptr), c.num)
	case cellFloats:
		if c.shared {
			return slices.Clone(c.floats())
		}
		return c.floats()
	case cellBytes:
		if c.shared {
			return slices.Clone(c.bytes())
		}
		return c.bytes()
	case cellUint64Box:
		return *(*uint64)(c.ptr)
	case cellInt64Box:
		return *(*int64)(c.ptr)
	case cellSub:
		return (*fieldState)(c.ptr)
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
	return s.getCell(fp).iface()
}

// getCell returns the cell at fp, or an empty cell if fp does not resolve to
// one.
func (s *fieldState) getCell(fp *fieldPath) cell {
	x := s
	for i := int32(0); i <= fp.last; i++ {
		z := int(fp.path[i])
		if len(x.state) < z+2 {
			return cell{}
		}
		if i == fp.last {
			return x.state[z]
		}
		sub, ok := x.state[z].sub()
		if !ok {
			return cell{}
		}
		x = sub
	}
	return cell{}
}

func (s *fieldState) set(fp *fieldPath, c cell) {
	x := s
	for i := int32(0); i <= fp.last; i++ {
		z := int(fp.path[i])
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
		} else if x.state[z].shared {
			// First write into a nested state shared with the baseline
			// template: copy it for this entity (copy on write). Its own
			// nested states and leaves stay shared until written.
			x.state[z] = subCell(&fieldState{state: slices.Clone((*fieldState)(x.state[z].ptr).state)})
		}
		x, _ = x.state[z].sub()
	}
}

// share makes s a class baseline template for clone: every vector and blob
// leaf and every nested state reachable from it is marked shared. s must not
// be modified afterwards. It is called once per decoded baseline.
func (s *fieldState) share() {
	for i := range s.state {
		switch c := &s.state[i]; c.kind {
		case cellSub:
			(*fieldState)(c.ptr).share()
			c.shared = true
		case cellFloats, cellBytes:
			c.shared = true
		}
	}
}

// clone returns a copy of a baseline template (see share) for a new entity,
// which can then be mutated independently of the template and its siblings.
// Only the root cell array is copied. Nested states stay shared until set
// first writes into one, which copies just that state; shared vectors and
// blobs are handed out by iface as copies (a caller may mutate a slice
// returned from Entity.Get/Map); strings and boxed integers are immutable. So
// nothing a clone exposes or modifies aliases template storage, and an entity
// pays only for the parts of the baseline it actually changes.
func (s *fieldState) clone() *fieldState {
	c := &fieldState{}
	s.cloneTo(c)
	return c
}

// cloneTo is clone into an existing root, such as the state held inline in
// an Entity.
func (s *fieldState) cloneTo(dst *fieldState) {
	dst.state = slices.Clone(s.state)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
