package manta

// A fieldpath, used to walk through the flattened table hierarchy
type fieldpath struct {
	hierarchy []*dt
	index     []int32
}

func fielpath_init(parentTbl *dt) *fieldpath {
	fp := &fieldpath{
		hierarchy: make([]*dt, 0),
		index:     make([]int32, 0),
	}

	fp.hierarchy = append(fp.hierarchy, parentTbl)
	fp.index = append(fp.index, -1) // Always start at -1

	return fp
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

}
