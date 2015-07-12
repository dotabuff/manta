package manta

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	fields    []*dt_field
	hierarchy []*dt
	index     []int32
	tree      *HuffmanTree
	finished  bool
}

// Global fieldpath lookup array
var fieldpath_lookup []func(*reader, *fieldpath)

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

	// Get's initialized only once
	if fieldpath_lookup == nil {
		fieldpath_lookup = make([]func(*reader, *fieldpath), 40)
		fieldpath_lookup[0] = PlusOne
		fieldpath_lookup[1] = PlusTwo
		fieldpath_lookup[2] = PlusThree
		fieldpath_lookup[3] = PlusFour
		fieldpath_lookup[4] = PlusN
		fieldpath_lookup[5] = PushOneLeftDeltaZeroRightZero
		fieldpath_lookup[6] = PushOneLeftDeltaZeroRightNonZero
		fieldpath_lookup[7] = PushOneLeftDeltaOneRightZero
		fieldpath_lookup[8] = PushOneLeftDeltaOneRightNonZero
		fieldpath_lookup[9] = PushOneLeftDeltaNRightZero
		fieldpath_lookup[10] = PushOneLeftDeltaNRightNonZero
		fieldpath_lookup[11] = PushOneLeftDeltaNRightNonZeroPack6Bits
		fieldpath_lookup[12] = PushOneLeftDeltaNRightNonZeroPack8Bits
		fieldpath_lookup[13] = PushTwoLeftDeltaZero
		fieldpath_lookup[14] = PushTwoLeftDeltaOne
		fieldpath_lookup[15] = PushTwoLeftDeltaN
		fieldpath_lookup[16] = PushTwoPack5LeftDeltaZero
		fieldpath_lookup[17] = PushTwoPack5LeftDeltaOne
		fieldpath_lookup[18] = PushTwoPack5LeftDeltaN
		fieldpath_lookup[19] = PushThreeLeftDeltaZero
		fieldpath_lookup[20] = PushThreeLeftDeltaOne
		fieldpath_lookup[21] = PushThreeLeftDeltaN
		fieldpath_lookup[22] = PushThreePack5LeftDeltaZero
		fieldpath_lookup[23] = PushThreePack5LeftDeltaOne
		fieldpath_lookup[24] = PushThreePack5LeftDeltaN
		fieldpath_lookup[25] = PushN
		fieldpath_lookup[26] = PushNAndNonTopological
		fieldpath_lookup[27] = PopOnePlusOne
		fieldpath_lookup[28] = PopOnePlusN
		fieldpath_lookup[29] = PopAllButOnePlusOne
		fieldpath_lookup[30] = PopAllButOnePlusN
		fieldpath_lookup[31] = PopAllButOnePlusNPack3Bits
		fieldpath_lookup[32] = PopAllButOnePlusNPack6Bits
		fieldpath_lookup[33] = PopNPlusOne
		fieldpath_lookup[34] = PopNPlusN
		fieldpath_lookup[35] = PopNAndNonTopographical
		fieldpath_lookup[36] = NonTopoComplex
		fieldpath_lookup[37] = NonTopoPenultimatePlusOne
		fieldpath_lookup[38] = NonTopoComplexPack4Bits
		fieldpath_lookup[39] = FieldPathEncodeFinish
	}

	return fp
}

// Walk an encoded fieldpath based on a huffman tree
func (fp *fieldpath) fieldpath_walk(r *reader) {
	// where is do-while when you need it -.-
	// @todo: Refactor this using node.IsLeaf()
	node := (*fp.tree).(HuffmanNode)
	for fp.finished == false {
		if r.readBits(1) == 1 {
			switch i := node.right.(type) {
			case HuffmanLeaf:
				fieldpath_lookup[i.value](r, fp)
				node = (*fp.tree).(HuffmanNode)
			case HuffmanNode:
				node = i
			}
		} else {
			switch i := node.left.(type) {
			case HuffmanLeaf:
				fieldpath_lookup[i.value](r, fp)
				node = (*fp.tree).(HuffmanNode)
			case HuffmanNode:
				node = i
			}
		}
	}
}

// Returns a huffman tree based on the operation weights
func fieldpath_huffman() HuffmanTree {
	FieldPathOperations := make(map[int]int)

	FieldPathOperations[0] = 36271
	FieldPathOperations[1] = 10334
	FieldPathOperations[2] = 1375
	FieldPathOperations[3] = 646
	FieldPathOperations[4] = 4128
	FieldPathOperations[5] = 35
	FieldPathOperations[6] = 3
	FieldPathOperations[7] = 521
	FieldPathOperations[8] = 2942
	FieldPathOperations[9] = 560
	FieldPathOperations[10] = 471
	FieldPathOperations[11] = 10530
	FieldPathOperations[12] = 251
	FieldPathOperations[13] = 0
	FieldPathOperations[14] = 0
	FieldPathOperations[15] = 0
	FieldPathOperations[16] = 0
	FieldPathOperations[17] = 0
	FieldPathOperations[18] = 0
	FieldPathOperations[19] = 0
	FieldPathOperations[20] = 0
	FieldPathOperations[21] = 0
	FieldPathOperations[22] = 0
	FieldPathOperations[23] = 0
	FieldPathOperations[24] = 0
	FieldPathOperations[25] = 0
	FieldPathOperations[26] = 310
	FieldPathOperations[27] = 2
	FieldPathOperations[28] = 0
	FieldPathOperations[29] = 1837
	FieldPathOperations[30] = 149
	FieldPathOperations[31] = 300
	FieldPathOperations[32] = 634
	FieldPathOperations[33] = 0
	FieldPathOperations[34] = 0
	FieldPathOperations[35] = 1
	FieldPathOperations[36] = 76
	FieldPathOperations[37] = 271
	FieldPathOperations[38] = 99
	FieldPathOperations[39] = 25474

	return buildTree(FieldPathOperations)
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
		printCodes(*fp.tree, []byte{})

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
