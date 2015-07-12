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
type fieldpathop struct {
	Name     string
	Function func(*reader, *fieldpath)
	Weight   int
}

// Global fieldpath lookup array
var fieldpath_lookup []fieldpathop

// Initialize a fieldpath object
func fieldpath_init(parentTbl *dt, huf *HuffmanTree) *fieldpath {
	fp := &fieldpath{
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
func (fp *fieldpath) fieldpath_walk(r *reader) {
	// where is do-while when you need it -.-
	// @todo: Refactor this using node.IsLeaf()
	cnt := 0
	node := (*fp.tree).(*HuffmanNode)
	for fp.finished == false {
		cnt++
		if r.readBits(1) == 1 {
			switch i := node.right.(type) {
			case *HuffmanLeaf:
				_debugf("Reached in %d bits", cnt)
				cnt = 0

				fieldpath_lookup[i.value].Function(r, fp)
				node = (*fp.tree).(*HuffmanNode)
			case *HuffmanNode:
				node = i
			}
		} else {
			switch i := node.left.(type) {
			case *HuffmanLeaf:
				_debugf("Reached in %d bits", cnt)
				cnt = 0

				fieldpath_lookup[i.value].Function(r, fp)
				node = (*fp.tree).(*HuffmanNode)
			case *HuffmanNode:
				node = i
			}
		}
	}
}

// Returns a huffman tree based on the operation weights
func fieldpath_huffman() HuffmanTree {
	// Get's initialized only once
	if fieldpath_lookup == nil {
		fieldpath_lookup = make([]fieldpathop, 40)
		fieldpath_lookup[0] = fieldpathop{"PlusOne", PlusOne, 36271}
		fieldpath_lookup[1] = fieldpathop{"FieldPathEncodeFinish", FieldPathEncodeFinish, 25474}
		fieldpath_lookup[2] = fieldpathop{"PushOneLeftDeltaNRightNonZeroPack6Bits", PushOneLeftDeltaNRightNonZeroPack6Bits, 10530}
		fieldpath_lookup[3] = fieldpathop{"PlusTwo", PlusTwo, 10334}
		fieldpath_lookup[4] = fieldpathop{"PlusN", PlusN, 4128}
		fieldpath_lookup[5] = fieldpathop{"PushOneLeftDeltaOneRightNonZero", PushOneLeftDeltaOneRightNonZero, 2942}
		fieldpath_lookup[6] = fieldpathop{"PopAllButOnePlusOne", PopAllButOnePlusOne, 1837}
		fieldpath_lookup[7] = fieldpathop{"PlusThree", PlusThree, 1375}
		fieldpath_lookup[8] = fieldpathop{"PlusFour", PlusFour, 646}
		fieldpath_lookup[9] = fieldpathop{"PopAllButOnePlusNPack6Bits", PopAllButOnePlusNPack6Bits, 634}
		fieldpath_lookup[10] = fieldpathop{"PushOneLeftDeltaNRightZero", PushOneLeftDeltaNRightZero, 560}
		fieldpath_lookup[11] = fieldpathop{"PushOneLeftDeltaOneRightZero", PushOneLeftDeltaOneRightZero, 521}
		fieldpath_lookup[12] = fieldpathop{"PushOneLeftDeltaNRightNonZero", PushOneLeftDeltaNRightNonZero, 471}
		fieldpath_lookup[13] = fieldpathop{"PushNAndNonTopological", PushNAndNonTopological, 310}
		fieldpath_lookup[14] = fieldpathop{"PopAllButOnePlusNPack3Bits", PopAllButOnePlusNPack3Bits, 300}
		fieldpath_lookup[15] = fieldpathop{"NonTopoPenultimatePlusOne", NonTopoPenultimatePlusOne, 271}
		fieldpath_lookup[16] = fieldpathop{"PushOneLeftDeltaNRightNonZeroPack8Bits", PushOneLeftDeltaNRightNonZeroPack8Bits, 251}
		fieldpath_lookup[17] = fieldpathop{"PopAllButOnePlusN", PopAllButOnePlusN, 149}
		fieldpath_lookup[18] = fieldpathop{"NonTopoComplexPack4Bits", NonTopoComplexPack4Bits, 99}
		fieldpath_lookup[19] = fieldpathop{"NonTopoComplex", NonTopoComplex, 76}
		fieldpath_lookup[20] = fieldpathop{"PushOneLeftDeltaZeroRightZero", PushOneLeftDeltaZeroRightZero, 35}
		fieldpath_lookup[21] = fieldpathop{"PushOneLeftDeltaZeroRightNonZero", PushOneLeftDeltaZeroRightNonZero, 3}
		fieldpath_lookup[22] = fieldpathop{"PopOnePlusOne", PopOnePlusOne, 1}                     // should be 2 but our algorithm is wrong
		fieldpath_lookup[23] = fieldpathop{"PopNAndNonTopographical", PopNAndNonTopographical, 2} // should be 1 but our algorithm is wrong
		fieldpath_lookup[24] = fieldpathop{"PopNPlusN", PopNPlusN, 0}
		fieldpath_lookup[25] = fieldpathop{"PopNPlusOne", PopNPlusOne, 0}
		fieldpath_lookup[26] = fieldpathop{"PopOnePlusN", PopOnePlusN, 0}
		fieldpath_lookup[27] = fieldpathop{"PushN", PushN, 0}
		fieldpath_lookup[28] = fieldpathop{"PushThreePack5LeftDeltaN", PushThreePack5LeftDeltaN, 0}
		fieldpath_lookup[29] = fieldpathop{"PushThreePack5LeftDeltaOne", PushThreePack5LeftDeltaOne, 0}
		fieldpath_lookup[30] = fieldpathop{"PushThreePack5LeftDeltaZero", PushThreePack5LeftDeltaZero, 0}
		fieldpath_lookup[31] = fieldpathop{"PushThreeLeftDeltaN", PushThreeLeftDeltaN, 0}
		fieldpath_lookup[32] = fieldpathop{"PushThreeLeftDeltaOne", PushThreeLeftDeltaOne, 0}
		fieldpath_lookup[33] = fieldpathop{"PushThreeLeftDeltaZero", PushThreeLeftDeltaZero, 0}
		fieldpath_lookup[34] = fieldpathop{"PushTwoPack5LeftDeltaN", PushTwoPack5LeftDeltaN, 0}
		fieldpath_lookup[35] = fieldpathop{"PushTwoPack5LeftDeltaOne", PushTwoPack5LeftDeltaOne, 0}
		fieldpath_lookup[36] = fieldpathop{"PushTwoPack5LeftDeltaZero", PushTwoPack5LeftDeltaZero, 0}
		fieldpath_lookup[37] = fieldpathop{"PushTwoLeftDeltaN", PushTwoLeftDeltaN, 0}
		fieldpath_lookup[38] = fieldpathop{"PushTwoLeftDeltaOne", PushTwoLeftDeltaOne, 0}
		fieldpath_lookup[39] = fieldpathop{"PushTwoLeftDeltaZero", PushTwoLeftDeltaZero, 0}
	}

	// Generate feq map
	huffmanlist := make([]int, 40)
	for i, fpo := range fieldpath_lookup {
		huffmanlist[i] = fpo.Weight
	}

	tree := buildTree(huffmanlist)

	return tree
}

func PlusOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PlusOne, %s", fp.hierarchy[0].Name)
	}

	// Increment the index
	fp.index[len(fp.index)-1] += 1

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow in PlusOne")
	}

	if field.Table == nil {
		fp.fields = append(fp.fields, field.Field)
	}
}

func PlusTwo(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PlusTwo, %s", fp.hierarchy[0].Name)
	}

	// Increment the index
	fp.index[len(fp.index)-1] += 2

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow in PlusOne")
	}

	if field.Table == nil {
		fp.fields = append(fp.fields, field.Field)
	}

	// Append the field to our list
	fp.fields = append(fp.fields, field.Field)
}

func PlusThree(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PlusThree, %s", fp.hierarchy[0].Name)
	}

	// Increment the index
	fp.index[len(fp.index)-1] += 3

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow in PlusOne")
	}

	if field.Table == nil {
		fp.fields = append(fp.fields, field.Field)
	}

	// Append the field to our list
	fp.fields = append(fp.fields, field.Field)
}

func PlusFour(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PlusFour, %s", fp.hierarchy[0].Name)
	}

	// Increment the index
	fp.index[len(fp.index)-1] += 4

	// Verify that the field exists
	tbl := fp.hierarchy[len(fp.index)-1]
	field := tbl.Properties[fp.index[len(fp.index)-1]]

	if field == nil {
		_panicf("Overflow in PlusOne")
	}

	if field.Table == nil {
		fp.fields = append(fp.fields, field.Field)
	}

	// Append the field to our list
	fp.fields = append(fp.fields, field.Field)
}

func PlusN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PlusN, %s", fp.hierarchy[0].Name)
	}
}

func PushOneLeftDeltaZeroRightZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaZeroRightZero, %s", fp.hierarchy[0].Name)
	}

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
	if debugMode {
		_debugf("Calling PushOneLeftDeltaZeroRightNonZero, %s, %d", fp.hierarchy[0].Name, fp.index[len(fp.index)-1])

		//
		// This get's called for the following fp_trace:
		//
		// '3/28/23','3/29',15
		//
		// Move:
		// -----
		// FROM: table["m_animationController.m_flPoseParameter"][23]
		// TO: table["m_animationController.m_AnimOverlay"]
		//
		// This doesn't seem correct, I'd expected PopOnePlusOne to be called here
	}
}

func PushOneLeftDeltaOneRightZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaOneRightZero, %s", fp.hierarchy[0].Name)
	}

	// PlusOne to advance the hierarchy to the next datatable
	PlusOne(r, fp)

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
					Index:      i,
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
}

func PushOneLeftDeltaOneRightNonZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaOneRightNonZero, %s", fp.hierarchy[0].Name)
	}
}

func PushOneLeftDeltaNRightZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaNRightZero, %s", fp.hierarchy[0].Name)
	}
}

func PushOneLeftDeltaNRightNonZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaNRightNonZero, %s", fp.hierarchy[0].Name)
	}
}

func PushOneLeftDeltaNRightNonZeroPack6Bits(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaNRightNonZeroPack6Bits, %s", fp.hierarchy[0].Name)
	}
}

func PushOneLeftDeltaNRightNonZeroPack8Bits(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaNRightNonZeroPack8Bits, %s", fp.hierarchy[0].Name)
	}
}

func PushTwoLeftDeltaZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushTwoLeftDeltaZero, %s", fp.hierarchy[0].Name)
	}
}

func PushTwoLeftDeltaOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushTwoLeftDeltaOne, %s", fp.hierarchy[0].Name)
	}
}

func PushTwoLeftDeltaN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushTwoLeftDeltaN, %s", fp.hierarchy[0].Name)
	}
}

func PushTwoPack5LeftDeltaZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushTwoPack5LeftDeltaZero, %s", fp.hierarchy[0].Name)
	}
}

func PushTwoPack5LeftDeltaOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushTwoPack5LeftDeltaOne, %s", fp.hierarchy[0].Name)
	}
}

func PushTwoPack5LeftDeltaN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushTwoPack5LeftDeltaN, %s", fp.hierarchy[0].Name)
	}
}

func PushThreeLeftDeltaZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushThreeLeftDeltaZero, %s", fp.hierarchy[0].Name)
	}
}

func PushThreeLeftDeltaOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushThreeLeftDeltaOne, %s", fp.hierarchy[0].Name)
	}
}

func PushThreeLeftDeltaN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushThreeLeftDeltaN, %s", fp.hierarchy[0].Name)
	}
}

func PushThreePack5LeftDeltaZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushThreePack5LeftDeltaZero, %s", fp.hierarchy[0].Name)
	}
}

func PushThreePack5LeftDeltaOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushThreePack5LeftDeltaOne, %s", fp.hierarchy[0].Name)
	}
}

func PushThreePack5LeftDeltaN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushThreePack5LeftDeltaN, %s", fp.hierarchy[0].Name)
	}
}

func PushN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushN, %s", fp.hierarchy[0].Name)
	}
}

func PushNAndNonTopological(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushNAndNonTopological, %s", fp.hierarchy[0].Name)
	}
}

func PopOnePlusOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopOnePlusOne, %s", fp.hierarchy[0].Name)
	}
}

func PopOnePlusN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopOnePlusN, %s", fp.hierarchy[0].Name)
	}
}

func PopAllButOnePlusOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopAllButOnePlusOne, %s", fp.hierarchy[0].Name)
	}
}

func PopAllButOnePlusN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopAllButOnePlusN, %s", fp.hierarchy[0].Name)
	}
}

func PopAllButOnePlusNPackN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopAllButOnePlusNPackN, %s", fp.hierarchy[0].Name)
	}
}

func PopAllButOnePlusNPack3Bits(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopAllButOnePlusNPack3Bits, %s", fp.hierarchy[0].Name)
	}
}

func PopAllButOnePlusNPack6Bits(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopAllButOnePlusNPack6Bits, %s", fp.hierarchy[0].Name)
	}
}

func PopNPlusOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopNPlusOne, %s", fp.hierarchy[0].Name)
	}
}

func PopNPlusN(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopNPlusN, %s", fp.hierarchy[0].Name)
	}
}

func PopNAndNonTopographical(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PopNAndNonTopographical, %s", fp.hierarchy[0].Name)
	}
}

func NonTopoComplex(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling NonTopoComplex, %s", fp.hierarchy[0].Name)
	}
}

func NonTopoPenultimatePlusOne(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling NonTopoPenultimatePlusOne, %s", fp.hierarchy[0].Name)
	}
}

func NonTopoComplexPack4Bits(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling NonTopoComplexPack4Bits, %s", fp.hierarchy[0].Name)
	}
}

func FieldPathEncodeFinish(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling FieldPathEncodeFinish, %s", fp.hierarchy[0].Name)
	}

	fp.finished = true
}
