package manta

import (
	"strconv"
)

// A huffmanTree for the fieldpath, built at init time based on fieldpath ops
var fpHuf huffmanTree = newFieldpathHuffman()

// Thanks to @spheenik for being resilient in his efforts to figure out the rest of the tree

// A single field to be read
type fieldpathField struct {
	Name  string
	Field *dtField
}

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	parent   *dt
	fields   []*fieldpathField
	index    []int32
	finished bool
}

// Contains the weight and lookup function for a single operation
type fieldpathOp struct {
	Name     string
	Function func(*reader, *fieldpath)
	Weight   int
}

// Global fieldpath lookup array
var fieldpathLookup = []fieldpathOp{
	{"PlusOne", fpPlusOne, 36271},
	{"PlusTwo", fpPlusTwo, 10334},
	{"PlusThree", fpPlusThree, 1375},
	{"PlusFour", fpPlusFour, 646},
	{"PlusN", fpPlusN, 4128},
	{"PushOneLeftDeltaZeroRightZero", fpPushOneLeftDeltaZeroRightZero, 35},
	{"PushOneLeftDeltaZeroRightNonZero", fpPushOneLeftDeltaZeroRightNonZero, 3},
	{"PushOneLeftDeltaOneRightZero", fpPushOneLeftDeltaOneRightZero, 521},
	{"PushOneLeftDeltaOneRightNonZero", fpPushOneLeftDeltaOneRightNonZero, 2942},
	{"PushOneLeftDeltaNRightZero", fpPushOneLeftDeltaNRightZero, 560},
	{"PushOneLeftDeltaNRightNonZero", fpPushOneLeftDeltaNRightNonZero, 471},
	{"PushOneLeftDeltaNRightNonZeroPack6Bits", fpPushOneLeftDeltaNRightNonZeroPack6Bits, 10530},
	{"PushOneLeftDeltaNRightNonZeroPack8Bits", fpPushOneLeftDeltaNRightNonZeroPack8Bits, 251},
	{"PushTwoLeftDeltaZero", fpPushTwoLeftDeltaZero, 0},
	{"PushTwoPack5LeftDeltaZero", fpPushTwoPack5LeftDeltaZero, 0},
	{"PushThreeLeftDeltaZero", fpPushThreeLeftDeltaZero, 0},
	{"PushThreePack5LeftDeltaZero", fpPushThreePack5LeftDeltaZero, 0},
	{"PushTwoLeftDeltaOne", fpPushTwoLeftDeltaOne, 0},
	{"PushTwoPack5LeftDeltaOne", fpPushTwoPack5LeftDeltaOne, 0},
	{"PushThreeLeftDeltaOne", fpPushThreeLeftDeltaOne, 0},
	{"PushThreePack5LeftDeltaOne", fpPushThreePack5LeftDeltaOne, 0},
	{"PushTwoLeftDeltaN", fpPushTwoLeftDeltaN, 0},
	{"PushTwoPack5LeftDeltaN", fpPushTwoPack5LeftDeltaN, 0},
	{"PushThreeLeftDeltaN", fpPushThreeLeftDeltaN, 0},
	{"PushThreePack5LeftDeltaN", fpPushThreePack5LeftDeltaN, 0},
	{"PushN", fpPushN, 0},
	{"PushNAndNonTopological", fpPushNAndNonTopological, 310},
	{"PopOnePlusOne", fpPopOnePlusOne, 2},
	{"PopOnePlusN", fpPopOnePlusN, 0},
	{"PopAllButOnePlusOne", fpPopAllButOnePlusOne, 1837},
	{"PopAllButOnePlusN", fpPopAllButOnePlusN, 149},
	{"PopAllButOnePlusNPack3Bits", fpPopAllButOnePlusNPack3Bits, 300},
	{"PopAllButOnePlusNPack6Bits", fpPopAllButOnePlusNPack6Bits, 634},
	{"PopNPlusOne", fpPopNPlusOne, 0},
	{"PopNPlusN", fpPopNPlusN, 0},
	{"PopNAndNonTopographical", fpPopNAndNonTopographical, 1},
	{"NonTopoComplex", fpNonTopoComplex, 76},
	{"NonTopoPenultimatePlusOne", fpNonTopoPenultimatePlusOne, 271},
	{"NonTopoComplexPack4Bits", fpNonTopoComplexPack4Bits, 99},
	{"FieldPathEncodeFinish", fpFieldPathEncodeFinish, 25474},
}

// Initialize a fieldpath object
func newFieldpath(parentTbl *dt) *fieldpath {
	fp := &fieldpath{
		parent:   parentTbl,
		fields:   make([]*fieldpathField, 0, 2),
		index:    make([]int32, 1, 2),
		finished: false,
	}
	fp.index[0] = -1 // always start at -1

	return fp
}

// Walk an encoded fieldpath based on a huffman tree
func (fp *fieldpath) walk(r *reader) {
	var node huffmanTree = fpHuf
	var next huffmanTree

	for !fp.finished {
		if r.readBits(1) == 1 {
			next = node.Right()
		} else {
			next = node.Left()
		}

		if next.IsLeaf() {
			node = fpHuf
			fieldpathLookup[next.Value()].Function(r, fp)
			if !fp.finished {
				fp.addField()
			}
		} else {
			node = next
		}
	}
}

// Adds a field based on the current index
func (fp *fieldpath) addField() {
	cDt := fp.parent

	var name string
	var i int

	if v(6) {
		var path string
		for i := 0; i < len(fp.index)-1; i++ {
			path += strconv.Itoa(int(fp.index[i])) + "/"
		}
		_debugf("adding field with path: %s%d", path, fp.index[len(fp.index)-1])
	}

	for i = 0; i < len(fp.index)-1; i++ {
		if cDt.Properties[fp.index[i]].Table != nil {
			cDt = cDt.Properties[fp.index[i]].Table
			name += cDt.Name + "."
		} else {

			// Hint:
			// If this panics, the property in question migh have a type that doesn't premit automatic array deduction (e.g. no CUtlVector prefix, or [] suffix).
			// Adjust the type manualy in property_serializers.go

			_panicf("expected table for type %s fp properties: %v, %v", cDt.Name, cDt.Properties[fp.index[i]].Field.Name, cDt.Properties[fp.index[i]].Field.Type)
		}
	}

	fp.fields = append(fp.fields, &fieldpathField{name + cDt.Properties[fp.index[i]].Field.Name, cDt.Properties[fp.index[i]].Field})
}

// Returns a huffman tree based on the operation weights
func newFieldpathHuffman() huffmanTree {
	// Generate freq map
	huffmanlist := make([]int, 40)
	for i, fpo := range fieldpathLookup {
		huffmanlist[i] = fpo.Weight
	}

	return buildHuffmanTree(huffmanlist)
}

// All different fieldops below

func fpPlusOne(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 1
}

func fpPlusTwo(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 2
}

func fpPlusThree(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 3
}

func fpPlusFour(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 4
}

func fpPlusN(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 5
}

func fpPushOneLeftDeltaZeroRightZero(r *reader, fp *fieldpath) {
	fp.index = append(fp.index, 0)
}

func fpPushOneLeftDeltaZeroRightNonZero(r *reader, fp *fieldpath) {
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushOneLeftDeltaOneRightZero(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 1
	fp.index = append(fp.index, 0)
}

func fpPushOneLeftDeltaOneRightNonZero(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 1
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushOneLeftDeltaNRightZero(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP())
	fp.index = append(fp.index, 0)
}

func fpPushOneLeftDeltaNRightNonZero(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 2
	fp.index = append(fp.index, int32(r.readUBitVarFP())+1)
}

func fpPushOneLeftDeltaNRightNonZeroPack6Bits(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readBits(3)) + 2
	fp.index = append(fp.index, int32(r.readBits(3))+1)
}

func fpPushOneLeftDeltaNRightNonZeroPack8Bits(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readBits(4)) + 2
	fp.index = append(fp.index, int32(r.readBits(4))+1)
}

func fpPushTwoLeftDeltaZero(r *reader, fp *fieldpath) {
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushTwoLeftDeltaOne(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushTwoLeftDeltaN(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushTwoPack5LeftDeltaZero(r *reader, fp *fieldpath) {
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func fpPushTwoPack5LeftDeltaOne(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func fpPushTwoPack5LeftDeltaN(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func fpPushThreeLeftDeltaZero(r *reader, fp *fieldpath) {
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushThreeLeftDeltaOne(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushThreeLeftDeltaN(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func fpPushThreePack5LeftDeltaZero(r *reader, fp *fieldpath) {
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func fpPushThreePack5LeftDeltaOne(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func fpPushThreePack5LeftDeltaN(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func fpPushN(r *reader, fp *fieldpath) {
	n := int(r.readUBitVar())
	fp.index[len(fp.index)-1] += int32(r.readUBitVar())

	for i := 0; i < n; i++ {
		fp.index = append(fp.index, int32(r.readUBitVarFP()))
	}
}

func fpPushNAndNonTopological(r *reader, fp *fieldpath) {
	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += r.readVarInt32() + 1
		}
	}

	count := int(r.readUBitVar())
	for j := 0; j < count; j++ {
		fp.index = append(fp.index, int32(r.readUBitVarFP()))
	}
}

func fpPopOnePlusOne(r *reader, fp *fieldpath) {
	fp.index = fp.index[:len(fp.index)-1]
	fp.index[len(fp.index)-1] += 1
}

func fpPopOnePlusN(r *reader, fp *fieldpath) {
	fp.index = fp.index[:len(fp.index)-1]
	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 1
}

func fpPopAllButOnePlusOne(r *reader, fp *fieldpath) {
	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += 1
}

func fpPopAllButOnePlusN(r *reader, fp *fieldpath) {
	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 1
}

func fpPopAllButOnePlusNPackN(r *reader, fp *fieldpath) {
	_panicf("Name: %s", fp.parent.Name)
}

func fpPopAllButOnePlusNPack3Bits(r *reader, fp *fieldpath) {
	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += int32(r.readBits(3)) + 1
}

func fpPopAllButOnePlusNPack6Bits(r *reader, fp *fieldpath) {
	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += int32(r.readBits(6)) + 1
}

func fpPopNPlusOne(r *reader, fp *fieldpath) {
	fp.index = fp.index[:len(fp.index)-(int(r.readUBitVarFP()))]
	fp.index[len(fp.index)-1] += 1
}

func fpPopNPlusN(r *reader, fp *fieldpath) {
	fp.index = fp.index[:len(fp.index)-(int(r.readUBitVarFP()))]
	fp.index[len(fp.index)-1] += r.readVarInt32()
}

func fpPopNAndNonTopographical(r *reader, fp *fieldpath) {
	fp.index = fp.index[:len(fp.index)-(int(r.readUBitVarFP()))]

	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += r.readVarInt32()
		}
	}
}

func fpNonTopoComplex(r *reader, fp *fieldpath) {
	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += r.readVarInt32()
		}
	}
}

func fpNonTopoPenultimatePlusOne(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-2] += 1
}

func fpNonTopoComplexPack4Bits(r *reader, fp *fieldpath) {
	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += int32(r.readBits(4)) - 7
		}
	}
}

func fpFieldPathEncodeFinish(r *reader, fp *fieldpath) {
	fp.finished = true
}
