package manta

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	hierarchy []*dt
	index     []int32
	tree      *HuffmanTree
	finished  bool
}

// Typedef for a field operation function
type FieldPathOpFcn func(*reader, *fieldpath)

// Initialize a fieldpath object
func fielpath_init(parentTbl *dt) *fieldpath {
	fp := &fieldpath{
		hierarchy: make([]*dt, 0),
		index:     make([]int32, 0),
	}

	fp.hierarchy = append(fp.hierarchy, parentTbl)
	fp.index = append(fp.index, -1) // Always start at -1
	fp.finished = false

	return fp
}

// Walk an encoded fieldpath based on a huffman tree
func (fp *fieldpath) fieldpath_walk(r *reader) []dt_field {
	fields := make([]dt_field, 0)

	// where is do-while when you need it -.-
	node := (*fp.tree).(HuffmanNode)
	for fp.finished == false {
		if r.readBits(1) == 1 {
			switch i := node.right.(type) {
			case HuffmanLeaf:
				i.value.(FieldPathOpFcn)(r, fp)
			case HuffmanNode:
				node = i
			}
		} else {
			switch i := node.left.(type) {
			case HuffmanLeaf:
				i.value.(FieldPathOpFcn)(r, fp)
			case HuffmanNode:
				node = i
			}
		}
	}

	return fields
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
	fp.index[len(fp.index)-1] += 1
}

func PlusTwo(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 2
}

func PlusThree(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 3
}

func PlusFour(r *reader, fp *fieldpath) {
	fp.index[len(fp.index)-1] += 4
}

func PlusN(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaZeroRightZero(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaZeroRightNonZero(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaOneRightZero(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaOneRightNonZero(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaNRightZero(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaNRightNonZero(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaNRightNonZeroPack6Bits(r *reader, fp *fieldpath) {

}

func PushOneLeftDeltaNRightNonZeroPack8Bits(r *reader, fp *fieldpath) {

}

func PushTwoLeftDeltaZero(r *reader, fp *fieldpath) {

}

func PushTwoLeftDeltaOne(r *reader, fp *fieldpath) {

}

func PushTwoLeftDeltaN(r *reader, fp *fieldpath) {

}

func PushTwoPack5LeftDeltaZero(r *reader, fp *fieldpath) {

}

func PushTwoPack5LeftDeltaOne(r *reader, fp *fieldpath) {

}

func PushTwoPack5LeftDeltaN(r *reader, fp *fieldpath) {

}

func PushThreeLeftDeltaZero(r *reader, fp *fieldpath) {

}

func PushThreeLeftDeltaOne(r *reader, fp *fieldpath) {

}

func PushThreeLeftDeltaN(r *reader, fp *fieldpath) {

}

func PushThreePack5LeftDeltaZero(r *reader, fp *fieldpath) {

}

func PushThreePack5LeftDeltaOne(r *reader, fp *fieldpath) {

}

func PushThreePack5LeftDeltaN(r *reader, fp *fieldpath) {

}

func PushN(r *reader, fp *fieldpath) {

}

func PushNAndNonTopological(r *reader, fp *fieldpath) {

}

func PopOnePlusOne(r *reader, fp *fieldpath) {

}

func PopOnePlusN(r *reader, fp *fieldpath) {

}

func PopAllButOnePlusOne(r *reader, fp *fieldpath) {

}

func PopAllButOnePlusN(r *reader, fp *fieldpath) {

}

func PopAllButOnePlusNPackN(r *reader, fp *fieldpath) {

}

func PopAllButOnePlusNPack3Bits(r *reader, fp *fieldpath) {

}

func PopAllButOnePlusNPack6Bits(r *reader, fp *fieldpath) {

}

func PopNPlusOne(r *reader, fp *fieldpath) {

}

func PopNPlusN(r *reader, fp *fieldpath) {

}

func PopNAndNonTopographical(r *reader, fp *fieldpath) {

}

func NonTopoComplex(r *reader, fp *fieldpath) {

}

func NonTopoPenultimatePlusOne(r *reader, fp *fieldpath) {

}

func NonTopoComplexPack4Bits(r *reader, fp *fieldpath) {

}

func FieldPathEncodeFinish(r *reader, fp *fieldpath) {
	fp.finished = true
}
