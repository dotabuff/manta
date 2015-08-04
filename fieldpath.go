package manta

import (
	"strconv"
)

// This is a list of encoding functions that are know to resolve to a correct layout.
// All of these have been verified
// ----------------------------------------------------------------------------------
//
// ID  NAME                           WEIGHT  ORIG  LEN  BITS
//  0  PlusOne                         36271          1   0
//  1  EncodingFinish                  25474          2   10
//  2  PlusTwo                         10334          4   1110
//  3  PlusN						    4128          5   11010
//  4  PlusThree                        1375          6   110010
//  5  PopAllButOnePlusOne              1837          6   110011
//  6  PushOneLeftDeltaOneRightZero      521          8   11011010
//  7  NonTopoComplexPack4Bits            99         10   1101100010
//  8  NonTopoComplex                     76         11   11011000111
//  9  PushOneLeftDeltaZeroRightZero      35         12   110110001101
// 10  PopOnePlusOne                       1     2   15   110110001100001
// 11  PushTwoLeftDeltaZero                0         27   110110001100100110000000000

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	parent   *dt
	fields   []*dt_field
	index    []int32
	tree     *HuffmanTree
	treeS    HuffmanTree // static version
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
	{"PlusOne", PlusOne, 36271},
	{"FieldPathEncodeFinish", FieldPathEncodeFinish, 25474},
	{"PushOneLeftDeltaNRightNonZeroPack6Bits", PushOneLeftDeltaNRightNonZeroPack6Bits, 10530},
	{"PlusTwo", PlusTwo, 10334},
	{"PlusN", PlusN, 4128},
	{"PushOneLeftDeltaOneRightNonZero", PushOneLeftDeltaOneRightNonZero, 2942},
	{"PopAllButOnePlusOne", PopAllButOnePlusOne, 1837},
	{"PlusThree", PlusThree, 1375},
	{"PlusFour", PlusFour, 646},
	{"PopAllButOnePlusNPack6Bits", PopAllButOnePlusNPack6Bits, 634},
	{"PushOneLeftDeltaNRightZero", PushOneLeftDeltaNRightZero, 560},
	{"PushOneLeftDeltaOneRightZero", PushOneLeftDeltaOneRightZero, 521},
	{"PushOneLeftDeltaNRightNonZero", PushOneLeftDeltaNRightNonZero, 471},
	{"PushNAndNonTopological", PushNAndNonTopological, 310},
	{"PopAllButOnePlusNPack3Bits", PopAllButOnePlusNPack3Bits, 300},
	{"NonTopoPenultimatePlusOne", NonTopoPenultimatePlusOne, 271},
	{"PushOneLeftDeltaNRightNonZeroPack8Bits", PushOneLeftDeltaNRightNonZeroPack8Bits, 251},
	{"PopAllButOnePlusN", PopAllButOnePlusN, 149},
	{"NonTopoComplexPack4Bits", NonTopoComplexPack4Bits, 99},
	{"NonTopoComplex", NonTopoComplex, 76},
	{"PushOneLeftDeltaZeroRightZero", PushOneLeftDeltaZeroRightZero, 35},
	{"PushOneLeftDeltaZeroRightNonZero", PushOneLeftDeltaZeroRightNonZero, 3},
	{"PopOnePlusOne", PopOnePlusOne, 1},                     // should be 2 but our algorithm is wrong
	{"PopNAndNonTopographical", PopNAndNonTopographical, 2}, // should be 1 but our algorithm is wrong
	{"PopNPlusN", PopNPlusN, 0},
	{"PopNPlusOne", PopNPlusOne, 0},
	{"PopOnePlusN", PopOnePlusN, 0},
	{"PushN", PushN, 0},
	{"PushThreePack5LeftDeltaN", PushThreePack5LeftDeltaN, 0},
	{"PushThreePack5LeftDeltaOne", PushThreePack5LeftDeltaOne, 0},
	{"PushThreePack5LeftDeltaZero", PushThreePack5LeftDeltaZero, 0},
	{"PushThreeLeftDeltaN", PushThreeLeftDeltaN, 0},
	{"PushThreeLeftDeltaOne", PushThreeLeftDeltaOne, 0},
	{"PushThreeLeftDeltaZero", PushThreeLeftDeltaZero, 0},
	{"PushTwoPack5LeftDeltaN", PushTwoPack5LeftDeltaN, 0},
	{"PushTwoPack5LeftDeltaOne", PushTwoPack5LeftDeltaOne, 0},
	{"PushTwoPack5LeftDeltaZero", PushTwoPack5LeftDeltaZero, 0},
	{"PushTwoLeftDeltaN", PushTwoLeftDeltaN, 0},
	{"PushTwoLeftDeltaOne", PushTwoLeftDeltaOne, 0},
	{"PushTwoLeftDeltaZero", PushTwoLeftDeltaZero, 0},
}

// Initialize a fieldpath object
func newFieldpath(parentTbl *dt, huf *HuffmanTree) *fieldpath {
	fp := &fieldpath{
		parent:   parentTbl,
		fields:   make([]*dt_field, 0),
		index:    make([]int32, 0),
		tree:     huf,
		treeS:    newFieldpathHuffmanStatic(),
		finished: false,
	}

	fp.index = append(fp.index, -1) // Always start at -1

	return fp
}

// Walk an encoded fieldpath based on a huffman tree
func (fp *fieldpath) walk(r *reader) {
	cnt := 0
	root := fp.treeS
	node := root

	for fp.finished == false {
		cnt++
		if r.readBits(1) == 1 {
			if i := node.Right(); i.IsLeaf() {
				node = root
				fieldpathLookup[i.Value()].Function(r, fp)
				fp.addField()

				_debugf("Reached in %d bits, %s", cnt, fp.fields[len(fp.fields)-1].Name)
				cnt = 0
			} else {
				node = i
			}
		} else {
			if i := node.Left(); i.IsLeaf() {
				node = root
				fieldpathLookup[i.Value()].Function(r, fp)
				fp.addField()

				_debugf("Reached in %d bits, %s", cnt, fp.fields[len(fp.fields)-1].Name)
				cnt = 0
			} else {
				node = i
			}
		}
	}

	// Will always add one additional field for the finishEncoding operation, remove it
	fp.fields = fp.fields[:len(fp.fields)-1]
}

// Adds a field based on the current index
func (fp *fieldpath) addField() {
	cDt := fp.parent

	var path string
	i := 0

	for i = 0; i < len(fp.index)-1; i++ {
		path += strconv.Itoa(int(fp.index[i])) + "/"
	}

	_debugf("Adding field with path: %s%d", path, fp.index[len(fp.index)-1])

	for i = 0; i < len(fp.index)-1; i++ {
		if cDt.Properties[fp.index[i]].Table != nil {
			cDt = cDt.Properties[fp.index[i]].Table
		} else {
			_panicf("expected table in fp properties")
		}
	}

	fp.fields = append(fp.fields, cDt.Properties[fp.index[i]].Field)
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

// Returns the static huffman tree based on our observed tree states
func newFieldpathHuffmanStatic() HuffmanTree {
	var h HuffmanTree
	h = &HuffmanNode{0, nil, nil}

	addNode(h, 0, 1, 0)        // PlusOne
	addNode(h, 1, 2, 1)        // EncodingFinish
	addNode(h, 7, 4, 3)        // PlusTwo
	addNode(h, 11, 5, 4)       // PlusN
	addNode(h, 51, 6, 6)       // PopAllButOnePlusOne
	addNode(h, 19, 6, 7)       // PlusThree
	addNode(h, 251, 8, 8)      // PlusFour
	addNode(h, 91, 8, 11)      // PushOneLeftDeltaOneRightZero
	addNode(h, 283, 10, 18)    // NonTopoComplexPack4Bits
	addNode(h, 1819, 11, 19)   // NonTopoComplex
	addNode(h, 2843, 12, 20)   // PushOneLeftDeltaZeroRightZero
	addNode(h, 17179, 15, 22)  // PopOnePlusOne
	addNode(h, 103195, 27, 39) // PushTwoLeftDeltaZero

	return h
}

func PlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 1
}

func PlusTwo(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 2
}

func PlusThree(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 3
}

func PlusFour(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 4
}

func PlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// This one reads a variable-length header encoding the number of bits
	// to read for N. Five is always added to the result.

	rBits := []int{2, 4, 10, 17, 30}

	for _, bits := range rBits {
		if r.readBits(1) == 1 {
			// Always add 5 to the result
			fp.index[len(fp.index)-1] += int32(r.readBits(bits)) + 5
			return
		}
	}
}

func PushOneLeftDeltaZeroRightZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Get current field and index
	fp.index = append(fp.index, 0)
}

func PushOneLeftDeltaZeroRightNonZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// should be correct, not encountered however
	/*rBits := []int{2, 4, 10, 17, 30}

	for _, bits := range rBits {
		if r.readBits(1) == 1 {
			fp.index = append(fp.index, int32(r.readBits(bits)))
			_debugf("Index: %v, BitsL %v", fp.index, bits)
			return
		}
	}*/
}

func PushOneLeftDeltaOneRightZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Push +1, set index to 0
	fp.index[len(fp.index)-1] += 1
	fp.index = append(fp.index, 0)
}

func PushOneLeftDeltaOneRightNonZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushOneLeftDeltaNRightZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushOneLeftDeltaNRightNonZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

}

func PushOneLeftDeltaNRightNonZeroPack6Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushOneLeftDeltaNRightNonZeroPack8Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushTwoLeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	fp.index = append(fp.index, 0)
	fp.index = append(fp.index, 0)
}

func PushTwoLeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushTwoLeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushTwoPack5LeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushTwoPack5LeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushTwoPack5LeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushThreeLeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushThreeLeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushThreeLeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushThreePack5LeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushThreePack5LeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushThreePack5LeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PushNAndNonTopological(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopOnePlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Check if we can pop an element
	if len(fp.index) <= 1 {
		_panicf("Trying to pop last element")
	}

	fp.index = fp.index[:len(fp.index)-1]
	fp.index[len(fp.index)-1] += 1
}

func PopOnePlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopAllButOnePlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// Remove all hierarchy and index element
	fp.index = fp.index[:1]
	fp.index[len(fp.index)-1] += 1
}

func PopAllButOnePlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopAllButOnePlusNPackN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopAllButOnePlusNPack3Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopAllButOnePlusNPack6Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopNPlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopNPlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func PopNAndNonTopographical(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func NonTopoComplex(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// See NonTopoComplexPack4Bits

	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += r.readVarInt32()
		}
	}
}

func NonTopoPenultimatePlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)
}

func NonTopoComplexPack4Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	// NonTopological = Disregard the hierarchy, work directly on the field
	// indizies for now
	//
	// Variables:
	// v4 = 0; // Incremented by 1 each loop
	// v3 = CFieldPath;
	//
	// Assumptions:
	// - Path data (array with MaxDepth) is first element of CFieldPath
	// - Current depth has an offset of 8 from CFieldPath
	//
	// Each loop does the following:
	// - Read 1 bit, if it is set, break
	// - Read 4 bits, substract 7 = v5
	// - Apply the data read to the v4'th index: v3[v4] += v5
	//
	// End condition:
	// - r.readBits(1) == 1
	// - Reached current depth (see assumption)

	for i := 0; i < len(fp.index); i++ {
		if r.readBoolean() {
			fp.index[i] += int32(r.readBits(4)) - 7
		}
	}
}

func FieldPathEncodeFinish(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.parent.Name)

	fp.finished = true
}
