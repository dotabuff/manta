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

// A fieldPath is int32-sized throughout (field and element indices are small)
// to keep it at 32 bytes: readFieldPaths snapshots one per decoded field.
type fieldPath struct {
	path [maxFieldPathDepth]int32
	last int32
}

type fieldPathOp struct {
	name   string
	weight int
	fn     func(r *reader, fp *fieldPath)
}

// Field-path op ids. An op's id is its huffman leaf value, so these follow the
// order of fieldPathTable (which is keyed by them).
const (
	fpOpPlusOne = iota
	fpOpPlusTwo
	fpOpPlusThree
	fpOpPlusFour
	fpOpPlusN
	fpOpPushOneLeftDeltaZeroRightZero
	fpOpPushOneLeftDeltaZeroRightNonZero
	fpOpPushOneLeftDeltaOneRightZero
	fpOpPushOneLeftDeltaOneRightNonZero
	fpOpPushOneLeftDeltaNRightZero
	fpOpPushOneLeftDeltaNRightNonZero
	fpOpPushOneLeftDeltaNRightNonZeroPack6Bits
	fpOpPushOneLeftDeltaNRightNonZeroPack8Bits
	fpOpPushTwoLeftDeltaZero
	fpOpPushTwoPack5LeftDeltaZero
	fpOpPushThreeLeftDeltaZero
	fpOpPushThreePack5LeftDeltaZero
	fpOpPushTwoLeftDeltaOne
	fpOpPushTwoPack5LeftDeltaOne
	fpOpPushThreeLeftDeltaOne
	fpOpPushThreePack5LeftDeltaOne
	fpOpPushTwoLeftDeltaN
	fpOpPushTwoPack5LeftDeltaN
	fpOpPushThreeLeftDeltaN
	fpOpPushThreePack5LeftDeltaN
	fpOpPushN
	fpOpPushNAndNonTopological
	fpOpPopOnePlusOne
	fpOpPopOnePlusN
	fpOpPopAllButOnePlusOne
	fpOpPopAllButOnePlusN
	fpOpPopAllButOnePlusNPack3Bits
	fpOpPopAllButOnePlusNPack6Bits
	fpOpPopNPlusOne
	fpOpPopNPlusN
	fpOpPopNAndNonTopographical
	fpOpNonTopoComplex
	fpOpNonTopoPenultimatePlusOne
	fpOpNonTopoComplexPack4Bits
	fpOpFieldPathEncodeFinish
)

// fieldPathTable defines each field-path op: its name, its huffman weight (which
// fixes the op codes), and its effect. readFieldPaths dispatches the same op
// functions through a switch of direct calls rather than through fn, so its
// accumulator does not escape; TestReadFieldPathsMatchesTable checks the two
// agree.
var fieldPathTable = []fieldPathOp{
	fpOpPlusOne:                                {"PlusOne", 36271, fpPlusOne},
	fpOpPlusTwo:                                {"PlusTwo", 10334, fpPlusTwo},
	fpOpPlusThree:                              {"PlusThree", 1375, fpPlusThree},
	fpOpPlusFour:                               {"PlusFour", 646, fpPlusFour},
	fpOpPlusN:                                  {"PlusN", 4128, fpPlusN},
	fpOpPushOneLeftDeltaZeroRightZero:          {"PushOneLeftDeltaZeroRightZero", 35, fpPushOneLeftDeltaZeroRightZero},
	fpOpPushOneLeftDeltaZeroRightNonZero:       {"PushOneLeftDeltaZeroRightNonZero", 3, fpPushOneLeftDeltaZeroRightNonZero},
	fpOpPushOneLeftDeltaOneRightZero:           {"PushOneLeftDeltaOneRightZero", 521, fpPushOneLeftDeltaOneRightZero},
	fpOpPushOneLeftDeltaOneRightNonZero:        {"PushOneLeftDeltaOneRightNonZero", 2942, fpPushOneLeftDeltaOneRightNonZero},
	fpOpPushOneLeftDeltaNRightZero:             {"PushOneLeftDeltaNRightZero", 560, fpPushOneLeftDeltaNRightZero},
	fpOpPushOneLeftDeltaNRightNonZero:          {"PushOneLeftDeltaNRightNonZero", 471, fpPushOneLeftDeltaNRightNonZero},
	fpOpPushOneLeftDeltaNRightNonZeroPack6Bits: {"PushOneLeftDeltaNRightNonZeroPack6Bits", 10530, fpPushOneLeftDeltaNRightNonZeroPack6Bits},
	fpOpPushOneLeftDeltaNRightNonZeroPack8Bits: {"PushOneLeftDeltaNRightNonZeroPack8Bits", 251, fpPushOneLeftDeltaNRightNonZeroPack8Bits},
	fpOpPushTwoLeftDeltaZero:                   {"PushTwoLeftDeltaZero", 0, fpPushTwoLeftDeltaZero},
	fpOpPushTwoPack5LeftDeltaZero:              {"PushTwoPack5LeftDeltaZero", 0, fpPushTwoPack5LeftDeltaZero},
	fpOpPushThreeLeftDeltaZero:                 {"PushThreeLeftDeltaZero", 0, fpPushThreeLeftDeltaZero},
	fpOpPushThreePack5LeftDeltaZero:            {"PushThreePack5LeftDeltaZero", 0, fpPushThreePack5LeftDeltaZero},
	fpOpPushTwoLeftDeltaOne:                    {"PushTwoLeftDeltaOne", 0, fpPushTwoLeftDeltaOne},
	fpOpPushTwoPack5LeftDeltaOne:               {"PushTwoPack5LeftDeltaOne", 0, fpPushTwoPack5LeftDeltaOne},
	fpOpPushThreeLeftDeltaOne:                  {"PushThreeLeftDeltaOne", 0, fpPushThreeLeftDeltaOne},
	fpOpPushThreePack5LeftDeltaOne:             {"PushThreePack5LeftDeltaOne", 0, fpPushThreePack5LeftDeltaOne},
	fpOpPushTwoLeftDeltaN:                      {"PushTwoLeftDeltaN", 0, fpPushTwoLeftDeltaN},
	fpOpPushTwoPack5LeftDeltaN:                 {"PushTwoPack5LeftDeltaN", 0, fpPushTwoPack5LeftDeltaN},
	fpOpPushThreeLeftDeltaN:                    {"PushThreeLeftDeltaN", 0, fpPushThreeLeftDeltaN},
	fpOpPushThreePack5LeftDeltaN:               {"PushThreePack5LeftDeltaN", 0, fpPushThreePack5LeftDeltaN},
	fpOpPushN:                                  {"PushN", 0, fpPushN},
	fpOpPushNAndNonTopological:                 {"PushNAndNonTopological", 310, fpPushNAndNonTopological},
	fpOpPopOnePlusOne:                          {"PopOnePlusOne", 2, fpPopOnePlusOne},
	fpOpPopOnePlusN:                            {"PopOnePlusN", 0, fpPopOnePlusN},
	fpOpPopAllButOnePlusOne:                    {"PopAllButOnePlusOne", 1837, fpPopAllButOnePlusOne},
	fpOpPopAllButOnePlusN:                      {"PopAllButOnePlusN", 149, fpPopAllButOnePlusN},
	fpOpPopAllButOnePlusNPack3Bits:             {"PopAllButOnePlusNPack3Bits", 300, fpPopAllButOnePlusNPack3Bits},
	fpOpPopAllButOnePlusNPack6Bits:             {"PopAllButOnePlusNPack6Bits", 634, fpPopAllButOnePlusNPack6Bits},
	fpOpPopNPlusOne:                            {"PopNPlusOne", 0, fpPopNPlusOne},
	fpOpPopNPlusN:                              {"PopNPlusN", 0, fpPopNPlusN},
	fpOpPopNAndNonTopographical:                {"PopNAndNonTopographical", 1, fpPopNAndNonTopographical},
	fpOpNonTopoComplex:                         {"NonTopoComplex", 76, fpNonTopoComplex},
	fpOpNonTopoPenultimatePlusOne:              {"NonTopoPenultimatePlusOne", 271, fpNonTopoPenultimatePlusOne},
	fpOpNonTopoComplexPack4Bits:                {"NonTopoComplexPack4Bits", 99, fpNonTopoComplexPack4Bits},
	fpOpFieldPathEncodeFinish:                  {"FieldPathEncodeFinish", 25474, fpFieldPathEncodeFinish},
}

func fpPlusOne(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 1
}

func fpPlusTwo(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 2
}

func fpPlusThree(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 3
}

func fpPlusFour(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 4
}

func fpPlusN(r *reader, fp *fieldPath) {
	fp.path[fp.last] += r.readUBitVarFieldPath() + 5
}

func fpPushOneLeftDeltaZeroRightZero(r *reader, fp *fieldPath) {
	fp.last++
	fp.path[fp.last] = 0
}

func fpPushOneLeftDeltaZeroRightNonZero(r *reader, fp *fieldPath) {
	fp.last++
	fp.path[fp.last] = r.readUBitVarFieldPath()
}

func fpPushOneLeftDeltaOneRightZero(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 1
	fp.last++
	fp.path[fp.last] = 0
}

func fpPushOneLeftDeltaOneRightNonZero(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 1
	fp.last++
	fp.path[fp.last] = r.readUBitVarFieldPath()
}

func fpPushOneLeftDeltaNRightZero(r *reader, fp *fieldPath) {
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] = 0
}

func fpPushOneLeftDeltaNRightNonZero(r *reader, fp *fieldPath) {
	fp.path[fp.last] += r.readUBitVarFieldPath() + 2
	fp.last++
	fp.path[fp.last] = r.readUBitVarFieldPath() + 1
}

func fpPushOneLeftDeltaNRightNonZeroPack6Bits(r *reader, fp *fieldPath) {
	fp.path[fp.last] += int32(r.readBits(3)) + 2
	fp.last++
	fp.path[fp.last] = int32(r.readBits(3)) + 1
}

func fpPushOneLeftDeltaNRightNonZeroPack8Bits(r *reader, fp *fieldPath) {
	fp.path[fp.last] += int32(r.readBits(4)) + 2
	fp.last++
	fp.path[fp.last] = int32(r.readBits(4)) + 1
}

func fpPushTwoLeftDeltaZero(r *reader, fp *fieldPath) {
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
}

func fpPushTwoPack5LeftDeltaZero(r *reader, fp *fieldPath) {
	fp.last++
	fp.path[fp.last] = int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] = int32(r.readBits(5))
}

func fpPushThreeLeftDeltaZero(r *reader, fp *fieldPath) {
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
}

func fpPushThreePack5LeftDeltaZero(r *reader, fp *fieldPath) {
	fp.last++
	fp.path[fp.last] = int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] = int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] = int32(r.readBits(5))
}

func fpPushTwoLeftDeltaOne(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 1
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
}

func fpPushTwoPack5LeftDeltaOne(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 1
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
}

func fpPushThreeLeftDeltaOne(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 1
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
}

func fpPushThreePack5LeftDeltaOne(r *reader, fp *fieldPath) {
	fp.path[fp.last] += 1
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
}

func fpPushTwoLeftDeltaN(r *reader, fp *fieldPath) {
	fp.path[fp.last] += int32(r.readUBitVar()) + 2
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
}

func fpPushTwoPack5LeftDeltaN(r *reader, fp *fieldPath) {
	fp.path[fp.last] += int32(r.readUBitVar()) + 2
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
}

func fpPushThreeLeftDeltaN(r *reader, fp *fieldPath) {
	fp.path[fp.last] += int32(r.readUBitVar()) + 2
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
	fp.last++
	fp.path[fp.last] += r.readUBitVarFieldPath()
}

func fpPushThreePack5LeftDeltaN(r *reader, fp *fieldPath) {
	fp.path[fp.last] += int32(r.readUBitVar()) + 2
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
	fp.last++
	fp.path[fp.last] += int32(r.readBits(5))
}

func fpPushN(r *reader, fp *fieldPath) {
	n := int(r.readUBitVar())
	fp.path[fp.last] += int32(r.readUBitVar())
	for i := 0; i < n; i++ {
		fp.last++
		fp.path[fp.last] += r.readUBitVarFieldPath()
	}
}

func fpPushNAndNonTopological(r *reader, fp *fieldPath) {
	for i := int32(0); i <= fp.last; i++ {
		if r.readBoolean() {
			fp.path[i] += r.readVarInt32() + 1
		}
	}
	count := int(r.readUBitVar())
	for i := 0; i < count; i++ {
		fp.last++
		fp.path[fp.last] = r.readUBitVarFieldPath()
	}
}

func fpPopOnePlusOne(r *reader, fp *fieldPath) {
	fp.pop(1)
	fp.path[fp.last] += 1
}

func fpPopOnePlusN(r *reader, fp *fieldPath) {
	fp.pop(1)
	fp.path[fp.last] += r.readUBitVarFieldPath() + 1
}

func fpPopAllButOnePlusOne(r *reader, fp *fieldPath) {
	fp.pop(fp.last)
	fp.path[0] += 1
}

func fpPopAllButOnePlusN(r *reader, fp *fieldPath) {
	fp.pop(fp.last)
	fp.path[0] += r.readUBitVarFieldPath() + 1
}

func fpPopAllButOnePlusNPack3Bits(r *reader, fp *fieldPath) {
	fp.pop(fp.last)
	fp.path[0] += int32(r.readBits(3)) + 1
}

func fpPopAllButOnePlusNPack6Bits(r *reader, fp *fieldPath) {
	fp.pop(fp.last)
	fp.path[0] += int32(r.readBits(6)) + 1
}

func fpPopNPlusOne(r *reader, fp *fieldPath) {
	fp.pop(r.readUBitVarFieldPath())
	fp.path[fp.last] += 1
}

func fpPopNPlusN(r *reader, fp *fieldPath) {
	fp.pop(r.readUBitVarFieldPath())
	fp.path[fp.last] += r.readVarInt32()
}

func fpPopNAndNonTopographical(r *reader, fp *fieldPath) {
	fp.pop(r.readUBitVarFieldPath())
	for i := int32(0); i <= fp.last; i++ {
		if r.readBoolean() {
			fp.path[i] += r.readVarInt32()
		}
	}
}

func fpNonTopoComplex(r *reader, fp *fieldPath) {
	for i := int32(0); i <= fp.last; i++ {
		if r.readBoolean() {
			fp.path[i] += r.readVarInt32()
		}
	}
}

func fpNonTopoPenultimatePlusOne(r *reader, fp *fieldPath) {
	fp.path[fp.last-1] += 1
}

func fpNonTopoComplexPack4Bits(r *reader, fp *fieldPath) {
	for i := int32(0); i <= fp.last; i++ {
		if r.readBoolean() {
			fp.path[i] += int32(r.readBits(4)) - 7
		}
	}
}

// fpFieldPathEncodeFinish ends the op stream; readFieldPaths stops on its op
// id, so it has no effect on the path.
func fpFieldPathEncodeFinish(r *reader, fp *fieldPath) {}

// pop reduces the last element by n, zeroing values in the popped path
func (fp *fieldPath) pop(n int32) {
	for i := int32(0); i < n; i++ {
		fp.path[fp.last] = 0
		fp.last--
	}
}

// copy returns a copy of the fieldPath
func (fp *fieldPath) copy() *fieldPath {
	x := fpPool.Get().(*fieldPath)
	x.path = fp.path
	x.last = fp.last
	return x
}

// String returns a string representing the fieldPath
func (fp *fieldPath) String() string {
	ss := make([]string, fp.last+1)
	for i := int32(0); i <= fp.last; i++ {
		ss[i] = strconv.Itoa(int(fp.path[i]))
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
	fp.path = [maxFieldPathDepth]int32{-1}
	fp.last = 0
}

// release returns the fieldPath to the pool for re-use
func (fp *fieldPath) release() {
	fpPool.Put(fp)
}

// readFieldPaths decodes the field-path operation stream from r into the
// provided buffer (which the caller resets to length 0) and returns the grown
// slice. Field paths are cumulative deltas, snapshotted by value into the
// buffer, so there is no per-path allocation and the buffer is reused across
// calls.
//
// Ops are dispatched through a switch of direct calls to the fieldPathTable op
// functions (most of which inline) rather than through the table's function
// values. Besides avoiding an indirect call per op, this lets escape analysis
// keep the accumulator on the stack; calling through fn forced it to the heap,
// which cost a sync.Pool round trip per call.
func readFieldPaths(r *reader, paths []fieldPath) []fieldPath {
	var fp fieldPath
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

		switch op {
		case fpOpPlusOne:
			fpPlusOne(r, &fp)
		case fpOpPlusTwo:
			fpPlusTwo(r, &fp)
		case fpOpPlusThree:
			fpPlusThree(r, &fp)
		case fpOpPlusFour:
			fpPlusFour(r, &fp)
		case fpOpPlusN:
			fpPlusN(r, &fp)
		case fpOpPushOneLeftDeltaZeroRightZero:
			fpPushOneLeftDeltaZeroRightZero(r, &fp)
		case fpOpPushOneLeftDeltaZeroRightNonZero:
			fpPushOneLeftDeltaZeroRightNonZero(r, &fp)
		case fpOpPushOneLeftDeltaOneRightZero:
			fpPushOneLeftDeltaOneRightZero(r, &fp)
		case fpOpPushOneLeftDeltaOneRightNonZero:
			fpPushOneLeftDeltaOneRightNonZero(r, &fp)
		case fpOpPushOneLeftDeltaNRightZero:
			fpPushOneLeftDeltaNRightZero(r, &fp)
		case fpOpPushOneLeftDeltaNRightNonZero:
			fpPushOneLeftDeltaNRightNonZero(r, &fp)
		case fpOpPushOneLeftDeltaNRightNonZeroPack6Bits:
			fpPushOneLeftDeltaNRightNonZeroPack6Bits(r, &fp)
		case fpOpPushOneLeftDeltaNRightNonZeroPack8Bits:
			fpPushOneLeftDeltaNRightNonZeroPack8Bits(r, &fp)
		case fpOpPushTwoLeftDeltaZero:
			fpPushTwoLeftDeltaZero(r, &fp)
		case fpOpPushTwoPack5LeftDeltaZero:
			fpPushTwoPack5LeftDeltaZero(r, &fp)
		case fpOpPushThreeLeftDeltaZero:
			fpPushThreeLeftDeltaZero(r, &fp)
		case fpOpPushThreePack5LeftDeltaZero:
			fpPushThreePack5LeftDeltaZero(r, &fp)
		case fpOpPushTwoLeftDeltaOne:
			fpPushTwoLeftDeltaOne(r, &fp)
		case fpOpPushTwoPack5LeftDeltaOne:
			fpPushTwoPack5LeftDeltaOne(r, &fp)
		case fpOpPushThreeLeftDeltaOne:
			fpPushThreeLeftDeltaOne(r, &fp)
		case fpOpPushThreePack5LeftDeltaOne:
			fpPushThreePack5LeftDeltaOne(r, &fp)
		case fpOpPushTwoLeftDeltaN:
			fpPushTwoLeftDeltaN(r, &fp)
		case fpOpPushTwoPack5LeftDeltaN:
			fpPushTwoPack5LeftDeltaN(r, &fp)
		case fpOpPushThreeLeftDeltaN:
			fpPushThreeLeftDeltaN(r, &fp)
		case fpOpPushThreePack5LeftDeltaN:
			fpPushThreePack5LeftDeltaN(r, &fp)
		case fpOpPushN:
			fpPushN(r, &fp)
		case fpOpPushNAndNonTopological:
			fpPushNAndNonTopological(r, &fp)
		case fpOpPopOnePlusOne:
			fpPopOnePlusOne(r, &fp)
		case fpOpPopOnePlusN:
			fpPopOnePlusN(r, &fp)
		case fpOpPopAllButOnePlusOne:
			fpPopAllButOnePlusOne(r, &fp)
		case fpOpPopAllButOnePlusN:
			fpPopAllButOnePlusN(r, &fp)
		case fpOpPopAllButOnePlusNPack3Bits:
			fpPopAllButOnePlusNPack3Bits(r, &fp)
		case fpOpPopAllButOnePlusNPack6Bits:
			fpPopAllButOnePlusNPack6Bits(r, &fp)
		case fpOpPopNPlusOne:
			fpPopNPlusOne(r, &fp)
		case fpOpPopNPlusN:
			fpPopNPlusN(r, &fp)
		case fpOpPopNAndNonTopographical:
			fpPopNAndNonTopographical(r, &fp)
		case fpOpNonTopoComplex:
			fpNonTopoComplex(r, &fp)
		case fpOpNonTopoPenultimatePlusOne:
			fpNonTopoPenultimatePlusOne(r, &fp)
		case fpOpNonTopoComplexPack4Bits:
			fpNonTopoComplexPack4Bits(r, &fp)
		case fpOpFieldPathEncodeFinish:
			return paths
		}
		paths = append(paths, fp)
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
