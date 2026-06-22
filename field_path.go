package manta

import (
	"strconv"
	"strings"
	"sync"
)

var huffTree = newHuffmanTree()

// Flattened representation of huffTree for fast field-path op decoding. Each
// internal node n occupies huffTreeLeft[n] and huffTreeRight[n]; a non-negative
// entry is a child node index and a negative entry -(op+1) is a leaf carrying
// field-path op id op. Walking int arrays avoids the interface-method dispatch
// and pointer chasing of the huffmanTree on the hot path. It is derived from
// manta's own huffTree, so the codes are identical to the interface-tree walk.
var (
	huffTreeLeft  []int32
	huffTreeRight []int32
	huffTreeRoot  int32
)

func init() {
	huffTreeRoot = flattenHuffmanTree(huffTree)
	buildHuffLookup()
}

// huffLookupBits is the width of the field-path op lookup window. Most ops have
// codes no longer than this and resolve in a single table index.
const huffLookupBits = 8

// huffLookup resolves up to huffLookupBits of the field-path op stream in one
// step. Each entry packs the consumed bit count in its low byte (0 meaning the
// code is longer than the window) and, in its high byte, either the resolved op
// id (when the low byte is non-zero) or the flat-tree node to continue from.
var huffLookup [1 << huffLookupBits]uint16

func buildHuffLookup() {
	for v := 0; v < len(huffLookup); v++ {
		node := huffTreeRoot
		resolved := false
		for bit := uint32(0); bit < huffLookupBits; bit++ {
			var child int32
			if (v>>bit)&1 == 1 {
				child = huffTreeRight[node]
			} else {
				child = huffTreeLeft[node]
			}
			if child < 0 {
				huffLookup[v] = uint16(bit+1) | uint16(-child-1)<<8
				resolved = true
				break
			}
			node = child
		}
		if !resolved {
			huffLookup[v] = uint16(node) << 8
		}
	}
}

// flattenHuffmanTree appends the internal nodes of t to huffTreeLeft/huffTreeRight
// and returns the index assigned to t. Leaf children are encoded inline as
// -(op+1).
func flattenHuffmanTree(t huffmanTree) int32 {
	idx := int32(len(huffTreeLeft))
	huffTreeLeft = append(huffTreeLeft, 0)
	huffTreeRight = append(huffTreeRight, 0)

	if l := t.Left(); l.IsLeaf() {
		huffTreeLeft[idx] = -(int32(l.Value()) + 1)
	} else {
		huffTreeLeft[idx] = flattenHuffmanTree(l)
	}
	if r := t.Right(); r.IsLeaf() {
		huffTreeRight[idx] = -(int32(r.Value()) + 1)
	} else {
		huffTreeRight[idx] = flattenHuffmanTree(r)
	}

	return idx
}

// maxFieldPathDepth is the maximum field-path depth, matching clarity's
// S2LongFieldPathFormat (whose BITS_PER_COMPONENT has 7 entries). The fixed-size
// path array enforces this invariant: a deeper push hits a bounds-check panic
// that the parser's top-level recover turns into an error, rather than silently
// corrupting entity state.
const maxFieldPathDepth = 7

type fieldPath struct {
	path [maxFieldPathDepth]int
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
	x.path = fp.path
	x.last = fp.last
	x.done = fp.done
	return x
}

// String returns a string representing the fieldPath
func (fp *fieldPath) String() string {
	ss := make([]string, fp.last+1)
	for i := 0; i <= fp.last; i++ {
		ss[i] = strconv.Itoa(fp.path[i])
	}
	return strings.Join(ss, "/")
}

// newFieldPath returns a new fieldPath ready for use
func newFieldPath() *fieldPath {
	fp := fpPool.Get().(*fieldPath)
	fp.reset()
	return fp
}

var fpPool = &sync.Pool{
	New: func() interface{} {
		return &fieldPath{}
	},
}

// reset resets the fieldPath to the empty value
func (fp *fieldPath) reset() {
	fp.path = [maxFieldPathDepth]int{-1}
	fp.last = 0
	fp.done = false
}

// release returns the fieldPath to the pool for re-use
func (fp *fieldPath) release() {
	fpPool.Put(fp)
}

// readFieldPaths decodes the field-path operation stream from r into the
// provided buffer (which the caller resets to length 0) and returns the grown
// slice. Field paths are cumulative deltas, snapshotted by value into the
// buffer, so there is no per-path allocation or sync.Pool churn and the buffer
// is reused across calls.
func readFieldPaths(r *reader, paths []fieldPath) []fieldPath {
	// Borrow a reusable accumulator from the pool. Taking its address for the
	// field-path op functions (an indirect call) would otherwise force it to
	// escape and heap-allocate on every call; the pool keeps that amortized.
	fp := fpPool.Get().(*fieldPath)
	fp.reset()

	for {
		// Resolve the op via the 8-bit lookup table, falling back to a flat-tree
		// walk for the rare codes longer than the lookup window. peekBits never
		// over-reads, so the final FieldPathEncodeFinish near the end of the
		// buffer is handled correctly.
		var op int32
		entry := huffLookup[r.peekBits(huffLookupBits)]
		if consumed := entry & 0xFF; consumed != 0 {
			r.skipBits(uint32(consumed))
			op = int32(entry >> 8)
		} else {
			r.skipBits(huffLookupBits)
			node := int32(entry >> 8)
			for {
				var child int32
				if r.readBits(1) == 1 {
					child = huffTreeRight[node]
				} else {
					child = huffTreeLeft[node]
				}
				if child < 0 {
					op = -child - 1
					break
				}
				node = child
			}
		}

		fieldPathTable[op].fn(r, fp)
		if fp.done {
			fpPool.Put(fp)
			return paths
		}
		paths = append(paths, *fp)
	}
}

// newHuffmanTree creates a new huffmanTree from the field path table
func newHuffmanTree() huffmanTree {
	freqs := make([]int, len(fieldPathTable))
	for i, op := range fieldPathTable {
		freqs[i] = op.weight
	}
	return buildHuffmanTree(freqs)
}
