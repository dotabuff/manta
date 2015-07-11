package manta

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	fields    []*dt_field
	hierarchy []*dt
	index     []int32
	tree      *HuffmanTree
	finished  bool
}

// Typedef for a field operation function
type FieldPathOpFcn func(*reader, *fieldpath)

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

	var x FieldPathOpFcn
	x = PlusOne
	_ = x

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
				i.value.(func(*reader, *fieldpath))(r, fp)
				node = (*fp.tree).(HuffmanNode)
			case HuffmanNode:
				node = i
			}
		} else {
			switch i := node.left.(type) {
			case HuffmanLeaf:
				i.value.(func(*reader, *fieldpath))(r, fp)
				node = (*fp.tree).(HuffmanNode)
			case HuffmanNode:
				node = i
			}
		}
	}
}

// Returns a huffman tree based on the operation weights
func fieldpath_huffman() HuffmanTree {
	FieldPathOperations := make(map[int]interface{})

	FieldPathOperations[36271] = PlusOne
	FieldPathOperations[10334] = PlusTwo
	FieldPathOperations[1375] = PlusThree
	FieldPathOperations[646] = PlusFour
	FieldPathOperations[4128] = PlusN
	FieldPathOperations[35] = PushOneLeftDeltaZeroRightZero
	FieldPathOperations[3] = PushOneLeftDeltaZeroRightNonZero
	FieldPathOperations[521] = PushOneLeftDeltaOneRightZero
	FieldPathOperations[2942] = PushOneLeftDeltaOneRightNonZero
	FieldPathOperations[560] = PushOneLeftDeltaNRightZero
	FieldPathOperations[471] = PushOneLeftDeltaNRightNonZero
	FieldPathOperations[10530] = PushOneLeftDeltaNRightNonZeroPack6Bits
	FieldPathOperations[251] = PushOneLeftDeltaNRightNonZeroPack8Bits
	FieldPathOperations[0] = PushTwoLeftDeltaZero
	FieldPathOperations[0] = PushTwoLeftDeltaOne
	FieldPathOperations[0] = PushTwoLeftDeltaN
	FieldPathOperations[0] = PushTwoPack5LeftDeltaZero
	FieldPathOperations[0] = PushTwoPack5LeftDeltaOne
	FieldPathOperations[0] = PushTwoPack5LeftDeltaN
	FieldPathOperations[0] = PushThreeLeftDeltaZero
	FieldPathOperations[0] = PushThreeLeftDeltaOne
	FieldPathOperations[0] = PushThreeLeftDeltaN
	FieldPathOperations[0] = PushThreePack5LeftDeltaZero
	FieldPathOperations[0] = PushThreePack5LeftDeltaOne
	FieldPathOperations[0] = PushThreePack5LeftDeltaN
	FieldPathOperations[0] = PushN
	FieldPathOperations[310] = PushNAndNonTopological
	FieldPathOperations[2] = PopOnePlusOne
	FieldPathOperations[0] = PopOnePlusN
	FieldPathOperations[1837] = PopAllButOnePlusOne
	FieldPathOperations[149] = PopAllButOnePlusN
	FieldPathOperations[300] = PopAllButOnePlusNPack3Bits
	FieldPathOperations[634] = PopAllButOnePlusNPack6Bits
	FieldPathOperations[0] = PopNPlusOne
	FieldPathOperations[0] = PopNPlusN
	FieldPathOperations[1] = PopNAndNonTopographical
	FieldPathOperations[76] = NonTopoComplex
	FieldPathOperations[271] = NonTopoPenultimatePlusOne
	FieldPathOperations[99] = NonTopoComplexPack4Bits
	FieldPathOperations[25474] = FieldPathEncodeFinish

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

	if field.Table != nil {
		_panicf("Trying to push dt as field index")
	}

	// Append the field to our list
	fp.fields = append(fp.fields, field.Field)
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

	if field.Table != nil {
		_panicf("Trying to push dt as field index")
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

	if field.Table != nil {
		_panicf("Trying to push dt as field index")
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

	if field.Table != nil {
		_panicf("Trying to push dt as field index")
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
}

func PushOneLeftDeltaZeroRightNonZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaZeroRightNonZero, %s", fp.hierarchy[0].Name)
	}
}

func PushOneLeftDeltaOneRightZero(r *reader, fp *fieldpath) {
	if debugMode {
		_debugf("Calling PushOneLeftDeltaOneRightZero, %s", fp.hierarchy[0].Name)
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
