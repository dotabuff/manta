package manta

import (
	"strconv"
)

// Thanks to @spheenik for being resilient in his efforts to figure out the rest of the tree

// A single field to be read
type fieldpath_field struct {
	Name  string
	Field *dt_field
}

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	parent   *dt
	fields   []*fieldpath_field
	index    []int32
	tree     *HuffmanTree
	finished bool
}

// Contains the weight and lookup function for a single operation
type fieldpathOp struct {
	Name     string
	Function func(*Reader, *fieldpath)
	Weight   int
}

// Global fieldpath lookup array
var fieldpathLookup = []fieldpathOp{
	{"PlusOne", PlusOne, 36271},
	{"PlusTwo", PlusTwo, 10334},
	{"PlusThree", PlusThree, 1375},
	{"PlusFour", PlusFour, 646},
	{"PlusN", PlusN, 4128},
	{"PushOneLeftDeltaZeroRightZero", PushOneLeftDeltaZeroRightZero, 35},
	{"PushOneLeftDeltaZeroRightNonZero", PushOneLeftDeltaZeroRightNonZero, 3},
	{"PushOneLeftDeltaOneRightZero", PushOneLeftDeltaOneRightZero, 521},
	{"PushOneLeftDeltaOneRightNonZero", PushOneLeftDeltaOneRightNonZero, 2942},
	{"PushOneLeftDeltaNRightZero", PushOneLeftDeltaNRightZero, 560},
	{"PushOneLeftDeltaNRightNonZero", PushOneLeftDeltaNRightNonZero, 471},
	{"PushOneLeftDeltaNRightNonZeroPack6Bits", PushOneLeftDeltaNRightNonZeroPack6Bits, 10530},
	{"PushOneLeftDeltaNRightNonZeroPack8Bits", PushOneLeftDeltaNRightNonZeroPack8Bits, 251},
	{"PushTwoLeftDeltaZero", PushTwoLeftDeltaZero, 0},
	{"PushTwoPack5LeftDeltaZero", PushTwoPack5LeftDeltaZero, 0},
	{"PushThreeLeftDeltaZero", PushThreeLeftDeltaZero, 0},
	{"PushThreePack5LeftDeltaZero", PushThreePack5LeftDeltaZero, 0},
	{"PushTwoLeftDeltaOne", PushTwoLeftDeltaOne, 0},
	{"PushTwoPack5LeftDeltaOne", PushTwoPack5LeftDeltaOne, 0},
	{"PushThreeLeftDeltaOne", PushThreeLeftDeltaOne, 0},
	{"PushThreePack5LeftDeltaOne", PushThreePack5LeftDeltaOne, 0},
	{"PushTwoLeftDeltaN", PushTwoLeftDeltaN, 0},
	{"PushTwoPack5LeftDeltaN", PushTwoPack5LeftDeltaN, 0},
	{"PushThreeLeftDeltaN", PushThreeLeftDeltaN, 0},
	{"PushThreePack5LeftDeltaN", PushThreePack5LeftDeltaN, 0},
	{"PushN", PushN, 0},
	{"PushNAndNonTopological", PushNAndNonTopological, 310},
	{"PopOnePlusOne", PopOnePlusOne, 2},
	{"PopOnePlusN", PopOnePlusN, 0},
	{"PopAllButOnePlusOne", PopAllButOnePlusOne, 1837},
	{"PopAllButOnePlusN", PopAllButOnePlusN, 149},
	{"PopAllButOnePlusNPack3Bits", PopAllButOnePlusNPack3Bits, 300},
	{"PopAllButOnePlusNPack6Bits", PopAllButOnePlusNPack6Bits, 634},
	{"PopNPlusOne", PopNPlusOne, 0},
	{"PopNPlusN", PopNPlusN, 0},
	{"PopNAndNonTopographical", PopNAndNonTopographical, 1},
	{"NonTopoComplex", NonTopoComplex, 76},
	{"NonTopoPenultimatePlusOne", NonTopoPenultimatePlusOne, 271},
	{"NonTopoComplexPack4Bits", NonTopoComplexPack4Bits, 99},
	{"FieldPathEncodeFinish", FieldPathEncodeFinish, 25474},
}

// Initialize a fieldpath object
func newFieldpath(parentTbl *dt, huf *HuffmanTree) *fieldpath {
	fp := &fieldpath{
		parent:   parentTbl,
		fields:   make([]*fieldpath_field, 0),
		index:    make([]int32, 0),
		tree:     huf,
		finished: false,
	}

	fp.index = append(fp.index, -1) // Always start at -1

	return fp
}

// Walk an encoded fieldpath based on a huffman tree
func (fp *fieldpath) walk(r *Reader) {
	cnt := 0
	root := fp.tree
	node := root

	for fp.finished == false {
		cnt++
		if r.readBits(1) == 1 {
			if i := (*node).Right(); i.IsLeaf() {
				node = root
				fieldpathLookup[i.Value()].Function(r, fp)

				if fp.finished == false {
					fp.addField()
					_debugfl(6, "Reached in %d bits, %s, %d", cnt, fp.fields[len(fp.fields)-1].Name, r.pos)
				}

				cnt = 0
			} else {
				node = &i
			}
		} else {
			if i := (*node).Left(); i.IsLeaf() {
				node = root
				fieldpathLookup[i.Value()].Function(r, fp)

				if fp.finished == false {
					fp.addField()
					_debugfl(6, "Reached in %d bits, %s, %d", cnt, fp.fields[len(fp.fields)-1].Name, r.pos)
				}

				cnt = 0
			} else {
				node = &i
			}
		}
	}
}

// Adds a field based on the current index
func (fp *fieldpath) addField() {
	cDt := fp.parent

	var name string
	var i int

	if debugLevel >= 6 {
		var path string
		for i := 0; i < len(fp.index)-1; i++ {
			path += strconv.Itoa(int(fp.index[i])) + "/"
		}

		_debugfl(6, "Adding field with path: %s%d", path, fp.index[len(fp.index)-1])
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

	fp.fields = append(fp.fields, &fieldpath_field{name + cDt.Properties[fp.index[i]].Field.Name, cDt.Properties[fp.index[i]].Field})
}

// Returns a huffman tree based on the operation weights
func newFieldpathHuffman() HuffmanTree {
	// Generate feq map
	huffmanlist := make([]int, 40)
	for i, fpo := range fieldpathLookup {
		huffmanlist[i] = fpo.Weight
	}

	return buildTree(huffmanlist)
}

// All different fieldops below

func PlusOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += 1
}

func PlusTwo(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += 2
}

func PlusThree(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += 3
}

func PlusFour(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += 4
}

func PlusN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 5
}

func PushOneLeftDeltaZeroRightZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = append(fp.index, 0)
}

func PushOneLeftDeltaZeroRightNonZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushOneLeftDeltaOneRightZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += 1
	fp.index = append(fp.index, 0)
}

func PushOneLeftDeltaOneRightNonZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += 1
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushOneLeftDeltaNRightZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP())
	fp.index = append(fp.index, 0)
}

func PushOneLeftDeltaNRightNonZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 2
	fp.index = append(fp.index, int32(r.readUBitVarFP())+1)
}

func PushOneLeftDeltaNRightNonZeroPack6Bits(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readBits(3)) + 2
	fp.index = append(fp.index, int32(r.readBits(3))+1)
}

func PushOneLeftDeltaNRightNonZeroPack8Bits(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readBits(4)) + 2
	fp.index = append(fp.index, int32(r.readBits(4))+1)
}

func PushTwoLeftDeltaZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushTwoLeftDeltaOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushTwoLeftDeltaN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushTwoPack5LeftDeltaZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func PushTwoPack5LeftDeltaOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func PushTwoPack5LeftDeltaN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func PushThreeLeftDeltaZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushThreeLeftDeltaOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushThreeLeftDeltaN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
	fp.index = append(fp.index, int32(r.readUBitVarFP()))
}

func PushThreePack5LeftDeltaZero(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func PushThreePack5LeftDeltaOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1]++
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func PushThreePack5LeftDeltaN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-1] += int32(r.readUBitVar()) + 2
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
	fp.index = append(fp.index, int32(r.readBits(5)))
}

func PushN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	n := int(r.readUBitVar())
	fp.index[len(fp.index)-1] += int32(r.readUBitVar())

	for i := 0; i < n; i++ {
		fp.index = append(fp.index, int32(r.readUBitVarFP()))
	}
}

func PushNAndNonTopological(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

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

func PopOnePlusOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:len(fp.index)-1]
	fp.index[len(fp.index)-1] += 1
}

func PopOnePlusN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:len(fp.index)-1]
	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 1
}

func PopAllButOnePlusOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += 1
}

func PopAllButOnePlusN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += int32(r.readUBitVarFP()) + 1
}

func PopAllButOnePlusNPackN(r *Reader, fp *fieldpath) {
	_panicf("Name: %s", fp.parent.Name)
}

func PopAllButOnePlusNPack3Bits(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += int32(r.readBits(3)) + 1
}

func PopAllButOnePlusNPack6Bits(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += int32(r.readBits(6)) + 1
}

func PopNPlusOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:len(fp.index)-(int(r.readUBitVarFP()))]
	fp.index[len(fp.index)-1] += 1
}

func PopNPlusN(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:len(fp.index)-(int(r.readUBitVarFP()))]
	fp.index[len(fp.index)-1] += r.readVarInt32()
}

func PopNAndNonTopographical(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index = fp.index[:len(fp.index)-(int(r.readUBitVarFP()))]

	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += r.readVarInt32()
		}
	}
}

func NonTopoComplex(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += r.readVarInt32()
		}
	}
}

func NonTopoPenultimatePlusOne(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.index[len(fp.index)-2] += 1
}

func NonTopoComplexPack4Bits(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += int32(r.readBits(4)) - 7
		}
	}
}

func FieldPathEncodeFinish(r *Reader, fp *fieldpath) {
	_debugfl(10, "Name: %s", fp.parent.Name)

	fp.finished = true
}
