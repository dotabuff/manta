package manta

// This is a list of encoding functions that are know to resolve to a correct layout.
// All of these have been verified
// ----------------------------------------------------------------------------------
//
// ID  NAME                           WEIGHT  ORIG  LEN  BITS
//  0  PlusOne                         36271          1   0
//  1  EncodingFinish                  25474          2   10
//  2  PlusTwo                         10334          4   1110
//  3  PushOneLeftDeltaZeroRightZero      35         12   110110001101
//  4  PlusThree                        1375          6   110010
//  5  PushOneLeftDeltaOneRightZero      521          8   11011010
//  6  PopOnePlusOne                       1     2   15   110110001100001
//  7  PopAllButOnePlusOne              1837          6   110011
//

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	fields    []*dt_field
	hierarchy []*dt
	index     []int32
	tree      *HuffmanTree
	finished  bool
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
		fields:    make([]*dt_field, 0),
		hierarchy: make([]*dt, 0),
		index:     make([]int32, 0),
	}

	fp.hierarchy = append(fp.hierarchy, parentTbl)
	fp.index = append(fp.index, -1) // Always start at -1
	fp.tree = huf
	fp.finished = false

	return fp
}

// Walk an encoded fieldpath based on a huffman tree
func (fp *fieldpath) walk(r *reader) {
	// where is do-while when you need it -.-
	// @todo: Refactor this using node.IsLeaf()

	cnt := 0
	root := HuffmanTree(*fp.tree)
	node := root

	for fp.finished == false {
		cnt++
		if r.readBits(1) == 1 {
			if i := node.Right(); i.IsLeaf() {
				fieldpathLookup[i.Value()].Function(r, fp)
				node = root

				_debugf("Reached in %d bits, %s", cnt, fp.fields[len(fp.fields)-1].Name)
				cnt = 0
			} else {
				node = i
			}
		} else {
			if i := node.Left(); i.IsLeaf() {
				fieldpathLookup[i.Value()].Function(r, fp)
				node = root

				_debugf("Reached in %d bits, %s", cnt, fp.fields[len(fp.fields)-1].Name)
				cnt = 0
			} else {
				node = i
			}
		}
	}
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

func PlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 1

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow")
	}

	// It's likely that we should actually push the tables
	// CWorld baseline advances from CPhysicsComponent.m_bCollisionActivationDisabled
	// to CRenderComponent and calls Finish without actually reading any element.
	// @todo: Investigate data, probably a handle
	fp.fields = append(fp.fields, field.Field)
}

func PlusTwo(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 2

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow")
	}

	fp.fields = append(fp.fields, field.Field)
}

func PlusThree(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 3

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow")
	}

	fp.fields = append(fp.fields, field.Field)
}

func PlusFour(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// Increment the index
	fp.index[len(fp.index)-1] += 4

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow")
	}

	fp.fields = append(fp.fields, field.Field)
}

func PlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// This one reads a variable-length header encoding the number of bits
	// to read for N. Five is always added to the result.

	rBits := []int{2, 4, 10, 17, 30}

	for _, bits := range rBits {
		if r.readBits(1) == 1 {
			// add 4 to the result and abuse PlusOne to append the field
			fp.index[len(fp.index)-1] += int32(r.readBits(bits)) + 4
			PlusOne(r, fp)
			return
		}
	}
}

func PushOneLeftDeltaZeroRightZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// Get current field and index
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field.Table == nil {
		_panicf("Trying to push field as table")
	}

	// Push the table, reset position to -1
	fp.hierarchy = append(fp.hierarchy, field.Table)
	fp.index = append(fp.index, -1)

	// We abuse PlusOne instead of copying the verification code
	PlusOne(r, fp)
}

func PushOneLeftDeltaZeroRightNonZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushOneLeftDeltaOneRightZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// PlusOne to advance the hierarchy to the next datatable
	fp.index[len(fp.index)-1] += 1

	// Get current field and index
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	// Are we pushing a field?
	if field.Table != nil {
		fp.hierarchy = append(fp.hierarchy, field.Table)
		fp.index = append(fp.index, -1)

		// We abuse PlusOne instead of copying the verification code
		PlusOne(r, fp)

		return
	}

	// Are we pushing an array?
	if field.Field.Serializer.IsArray {
		_debugf("Entering array subroutine")

		// Add our own temp table for the array
		tmpDt := &dt{
			Name:       field.Field.Name,
			Flags:      nil,
			Version:    0,
			Properties: make([]*dt_property, 0),
		}

		// Add each array field to the table
		for i := uint32(0); i < field.Field.Serializer.Length; i++ {
			tmpDt.Properties = append(tmpDt.Properties, &dt_property{
				Field: &dt_field{
					Name:       field.Field.Name,
					Type:       "",
					Index:      int32(i),
					Flags:      field.Field.Flags,
					BitCount:   field.Field.BitCount,
					LowValue:   field.Field.LowValue,
					HighValue:  field.Field.HighValue,
					Version:    field.Field.Version,
					Serializer: field.Field.Serializer.ArraySerializer,
				},
				Table: nil,
			})
		}

		fp.hierarchy = append(fp.hierarchy, tmpDt)
		fp.index = append(fp.index, -1)

		PlusOne(r, fp)

		return
	}

	_panicf("Type: %s is neither Array not Table", field.Field.Name)
}

func PushOneLeftDeltaOneRightNonZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushOneLeftDeltaNRightZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushOneLeftDeltaNRightNonZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushOneLeftDeltaNRightNonZeroPack6Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushOneLeftDeltaNRightNonZeroPack8Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushTwoLeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushTwoLeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushTwoLeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushTwoPack5LeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushTwoPack5LeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushTwoPack5LeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushThreeLeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushThreeLeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushThreeLeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushThreePack5LeftDeltaZero(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushThreePack5LeftDeltaOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushThreePack5LeftDeltaN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PushNAndNonTopological(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopOnePlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// Check if we can pop an element
	if len(fp.index) <= 1 {
		_panicf("Trying to pop last element")
	}

	// Pop last index and table
	fp.hierarchy = fp.hierarchy[:len(fp.hierarchy)-1]
	fp.index = fp.index[:len(fp.index)-1]

	// Read next element
	PlusOne(r, fp)
}

func PopOnePlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopAllButOnePlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// Remove all hierarchy and index element
	fp.hierarchy = fp.hierarchy[:1]
	fp.index = fp.index[:1]

	PlusOne(r, fp)
}

func PopAllButOnePlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopAllButOnePlusNPackN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopAllButOnePlusNPack3Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopAllButOnePlusNPack6Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopNPlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopNPlusN(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func PopNAndNonTopographical(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func NonTopoComplex(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	// See NonTopoComplexPack4Bits

	depth := 0
	for depth != len(fp.index) && r.readBits(1) == 1 {
		_debugf("Apply to index %d, %d", depth, r.readVarInt32())
		depth++
	}

	fp.hierarchy = fp.hierarchy[:len(fp.hierarchy)-1]
	fp.index = fp.index[:len(fp.index)-1]

	PushOneLeftDeltaOneRightZero(r, fp)
}

func NonTopoPenultimatePlusOne(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)
}

func NonTopoComplexPack4Bits(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

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

	depth := 0
	for depth != len(fp.index) && r.readBits(1) == 1 {
		_debugf("Todo: Apply to index %d, %d", depth, r.readBits(4)-7)
		depth++
	}

	// HACK FOR baseline's only:
	// --------------------------
	// For baselines's, we should switch from current/2 to current+1/0 in 99%
	// of the cases.
	// (reads +1, -X, (advance index [0]+1, advance index [1]-X)

	fp.hierarchy = fp.hierarchy[:len(fp.hierarchy)-1]
	fp.index = fp.index[:len(fp.index)-1]

	PushOneLeftDeltaOneRightZero(r, fp)
}

func FieldPathEncodeFinish(r *reader, fp *fieldpath) {
	_debugf("Name: %s", fp.hierarchy[0].Name)

	fp.finished = true
}
