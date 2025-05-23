package manta

import (
	"strconv"
	"strings"
	"sync"
)

var huffTree = newHuffmanTree()

type fieldPath struct {
	path []int
	last int
	done bool
}

type fieldPathOp struct {
	name   string
	weight int
	fn     func(r *reader, fp *fieldPath)
}

var fieldPathTable = []fieldPathOp{
	fieldPathOp{"PlusOne", 36271, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 1
	}},
	fieldPathOp{"PlusTwo", 10334, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 2
	}},
	fieldPathOp{"PlusThree", 1375, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 3
	}},
	fieldPathOp{"PlusFour", 646, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 4
	}},
	fieldPathOp{"PlusN", 4128, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += r.readUBitVarFieldPath() + 5
	}},
	fieldPathOp{"PushOneLeftDeltaZeroRightZero", 35, func(r *reader, fp *fieldPath) {
		fp.last++
		fp.path[fp.last] = 0
	}},
	fieldPathOp{"PushOneLeftDeltaZeroRightNonZero", 3, func(r *reader, fp *fieldPath) {
		fp.last++
		fp.path[fp.last] = r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushOneLeftDeltaOneRightZero", 521, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 1
		fp.last++
		fp.path[fp.last] = 0
	}},
	fieldPathOp{"PushOneLeftDeltaOneRightNonZero", 2942, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 1
		fp.last++
		fp.path[fp.last] = r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushOneLeftDeltaNRightZero", 560, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] = 0
	}},
	fieldPathOp{"PushOneLeftDeltaNRightNonZero", 471, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += r.readUBitVarFieldPath() + 2
		fp.last++
		fp.path[fp.last] = r.readUBitVarFieldPath() + 1
	}},
	fieldPathOp{"PushOneLeftDeltaNRightNonZeroPack6Bits", 10530, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += int(r.readBits(3)) + 2
		fp.last++
		fp.path[fp.last] = int(r.readBits(3)) + 1
	}},
	fieldPathOp{"PushOneLeftDeltaNRightNonZeroPack8Bits", 251, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += int(r.readBits(4)) + 2
		fp.last++
		fp.path[fp.last] = int(r.readBits(4)) + 1
	}},
	fieldPathOp{"PushTwoLeftDeltaZero", 0, func(r *reader, fp *fieldPath) {
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushTwoPack5LeftDeltaZero", 0, func(r *reader, fp *fieldPath) {
		fp.last++
		fp.path[fp.last] = int(r.readBits(5))
		fp.last++
		fp.path[fp.last] = int(r.readBits(5))
	}},
	fieldPathOp{"PushThreeLeftDeltaZero", 0, func(r *reader, fp *fieldPath) {
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushThreePack5LeftDeltaZero", 0, func(r *reader, fp *fieldPath) {
		fp.last++
		fp.path[fp.last] = int(r.readBits(5))
		fp.last++
		fp.path[fp.last] = int(r.readBits(5))
		fp.last++
		fp.path[fp.last] = int(r.readBits(5))
	}},
	fieldPathOp{"PushTwoLeftDeltaOne", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 1
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushTwoPack5LeftDeltaOne", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 1
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
	}},
	fieldPathOp{"PushThreeLeftDeltaOne", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 1
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushThreePack5LeftDeltaOne", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += 1
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
	}},
	fieldPathOp{"PushTwoLeftDeltaN", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += int(r.readUBitVar()) + 2
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushTwoPack5LeftDeltaN", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += int(r.readUBitVar()) + 2
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
	}},
	fieldPathOp{"PushThreeLeftDeltaN", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += int(r.readUBitVar()) + 2
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
	}},
	fieldPathOp{"PushThreePack5LeftDeltaN", 0, func(r *reader, fp *fieldPath) {
		fp.path[fp.last] += int(r.readUBitVar()) + 2
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
		fp.last++
		fp.path[fp.last] += int(r.readBits(5))
	}},
	fieldPathOp{"PushN", 0, func(r *reader, fp *fieldPath) {
		n := int(r.readUBitVar())
		fp.path[fp.last] += int(r.readUBitVar())
		for i := 0; i < n; i++ {
			fp.last++
			fp.path[fp.last] += r.readUBitVarFieldPath()
		}
	}},
	fieldPathOp{"PushNAndNonTopological", 310, func(r *reader, fp *fieldPath) {
		for i := 0; i <= fp.last; i++ {
			if r.readBoolean() {
				fp.path[i] += int(r.readVarInt32()) + 1
			}
		}
		count := int(r.readUBitVar())
		for i := 0; i < count; i++ {
			fp.last++
			fp.path[fp.last] = r.readUBitVarFieldPath()
		}
	}},
	fieldPathOp{"PopOnePlusOne", 2, func(r *reader, fp *fieldPath) {
		fp.pop(1)
		fp.path[fp.last] += 1
	}},
	fieldPathOp{"PopOnePlusN", 0, func(r *reader, fp *fieldPath) {
		fp.pop(1)
		fp.path[fp.last] += r.readUBitVarFieldPath() + 1
	}},
	fieldPathOp{"PopAllButOnePlusOne", 1837, func(r *reader, fp *fieldPath) {
		fp.pop(fp.last)
		fp.path[0] += 1
	}},
	fieldPathOp{"PopAllButOnePlusN", 149, func(r *reader, fp *fieldPath) {
		fp.pop(fp.last)
		fp.path[0] += r.readUBitVarFieldPath() + 1
	}},
	fieldPathOp{"PopAllButOnePlusNPack3Bits", 300, func(r *reader, fp *fieldPath) {
		fp.pop(fp.last)
		fp.path[0] += int(r.readBits(3)) + 1
	}},
	fieldPathOp{"PopAllButOnePlusNPack6Bits", 634, func(r *reader, fp *fieldPath) {
		fp.pop(fp.last)
		fp.path[0] += int(r.readBits(6)) + 1
	}},
	fieldPathOp{"PopNPlusOne", 0, func(r *reader, fp *fieldPath) {
		fp.pop(r.readUBitVarFieldPath())
		fp.path[fp.last] += 1
	}},
	fieldPathOp{"PopNPlusN", 0, func(r *reader, fp *fieldPath) {
		fp.pop(r.readUBitVarFieldPath())
		fp.path[fp.last] += int(r.readVarInt32())
	}},
	fieldPathOp{"PopNAndNonTopographical", 1, func(r *reader, fp *fieldPath) {
		fp.pop(r.readUBitVarFieldPath())
		for i := 0; i <= fp.last; i++ {
			if r.readBoolean() {
				fp.path[i] += int(r.readVarInt32())
			}
		}
	}},
	fieldPathOp{"NonTopoComplex", 76, func(r *reader, fp *fieldPath) {
		for i := 0; i <= fp.last; i++ {
			if r.readBoolean() {
				fp.path[i] += int(r.readVarInt32())
			}
		}
	}},
	fieldPathOp{"NonTopoPenultimatePlusOne", 271, func(r *reader, fp *fieldPath) {
		fp.path[fp.last-1] += 1
	}},
	fieldPathOp{"NonTopoComplexPack4Bits", 99, func(r *reader, fp *fieldPath) {
		for i := 0; i <= fp.last; i++ {
			if r.readBoolean() {
				fp.path[i] += int(r.readBits(4)) - 7
			}
		}
	}},
	fieldPathOp{"FieldPathEncodeFinish", 25474, func(r *reader, fp *fieldPath) {
		fp.done = true
	}},
}

// pop reduces the last element by n, zeroing values in the popped path
func (fp *fieldPath) pop(n int) {
	for i := 0; i < n; i++ {
		fp.path[fp.last] = 0
		fp.last--
	}
}

// copy returns a copy of the fieldPath
func (fp *fieldPath) copy() *fieldPath {
	x := fpPool.Get().(*fieldPath)
	copy(x.path, fp.path)
	x.last = fp.last
	x.done = fp.done
	return x
}

// String returns a string representing the fieldPath - optimized
func (fp *fieldPath) String() string {
	if fp.last == 0 {
		return strconv.Itoa(fp.path[0])
	}

	// Use strings.Builder for better performance
	var builder strings.Builder
	builder.Grow(fp.last * 4) // Estimate 4 chars per element

	builder.WriteString(strconv.Itoa(fp.path[0]))
	for i := 1; i <= fp.last; i++ {
		builder.WriteByte('/')
		builder.WriteString(strconv.Itoa(fp.path[i]))
	}
	return builder.String()
}

// newFieldPath returns a new fieldPath ready for use
func newFieldPath() *fieldPath {
	fp := fpPool.Get().(*fieldPath)
	fp.reset()
	return fp
}

// Optimized field path pool with better allocation patterns
var fpPool = &sync.Pool{
	New: func() interface{} {
		return &fieldPath{
			path: make([]int, 7),
			last: 0,
			done: false,
		}
	},
}

// Pool for field path slices to reduce allocations in readFieldPaths
var fpSlicePool = &sync.Pool{
	New: func() interface{} {
		// Pre-allocate with reasonable capacity based on typical usage
		return make([]*fieldPath, 0, 64)
	},
}

// Pre-warm the pool with some field paths to reduce early allocation pressure
func init() {
	// Pre-allocate some field paths to reduce initial allocation overhead
	for i := 0; i < 100; i++ {
		fp := &fieldPath{
			path: make([]int, 7),
			last: 0,
			done: false,
		}
		fpPool.Put(fp)
	}
}

// reset resets the fieldPath to the empty value - optimized version
func (fp *fieldPath) reset() {
	// Fast reset: only clear what we need
	fp.path[0] = -1
	for i := 1; i <= fp.last && i < len(fp.path); i++ {
		fp.path[i] = 0
	}
	fp.last = 0
	fp.done = false
}

// release returns the fieldPath to the pool for re-use
func (fp *fieldPath) release() {
	fpPool.Put(fp)
}

// readFieldPaths reads a new slice of fieldPath values from the given reader
func readFieldPaths(r *reader) []*fieldPath {
	fp := newFieldPath()

	node, next := huffTree, huffTree

	// Get pooled slice instead of allocating new one
	paths := fpSlicePool.Get().([]*fieldPath)
	paths = paths[:0] // Reset length but keep capacity

	for !fp.done {
		if r.readBits(1) == 1 {
			next = node.Right()
		} else {
			next = node.Left()
		}

		if next.IsLeaf() {
			node = huffTree
			fieldPathTable[next.Value()].fn(r, fp)
			if !fp.done {
				paths = append(paths, fp.copy())
			}
		} else {
			node = next
		}
	}

	fp.release()

	return paths
}

// releaseFieldPaths returns the field path slice to the pool after all paths are released
func releaseFieldPaths(fps []*fieldPath) {
	// Reset the slice for reuse but keep the capacity
	for i := range fps {
		fps[i] = nil // Clear references to help GC
	}
	fps = fps[:0]
	fpSlicePool.Put(fps)
}

// newHuffmanTree creates a new huffmanTree from the field path table
func newHuffmanTree() huffmanTree {
	freqs := make([]int, len(fieldPathTable))
	for i, op := range fieldPathTable {
		freqs[i] = op.weight
	}
	return buildHuffmanTree(freqs)
}
