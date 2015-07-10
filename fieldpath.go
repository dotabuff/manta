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

func (fp *fieldpath) PlusOne(r *reader) {

}

func (fp *fieldpath) PlusTwo(r *reader) {

}

func (fp *fieldpath) PlusThree(r *reader) {

}

func (fp *fieldpath) PlusFour(r *reader) {

}

func (fp *fieldpath) PlusN(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaZeroRightZero(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaZeroRightNonZero(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaOneRightZero(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaOneRightNonZero(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaNRightZero(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaNRightNonZero(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaNRightNonZeroPack6Bits(r *reader) {

}

func (fp *fieldpath) PushOneLeftDeltaNRightNonZeroPack8Bits(r *reader) {

}

func (fp *fieldpath) PushTwoLeftDeltaZero(r *reader) {

}

func (fp *fieldpath) PushTwoLeftDeltaOne(r *reader) {

}

func (fp *fieldpath) PushTwoLeftDeltaN(r *reader) {

}

func (fp *fieldpath) PushTwoPack5LeftDeltaZero(r *reader) {

}

func (fp *fieldpath) PushTwoPack5LeftDeltaOne(r *reader) {

}

func (fp *fieldpath) PushTwoPack5LeftDeltaN(r *reader) {

}

func (fp *fieldpath) PushThreeLeftDeltaZero(r *reader) {

}

func (fp *fieldpath) PushThreeLeftDeltaOne(r *reader) {

}

func (fp *fieldpath) PushThreeLeftDeltaN(r *reader) {

}

func (fp *fieldpath) PushThreePack5LeftDeltaZero(r *reader) {

}

func (fp *fieldpath) PushThreePack5LeftDeltaOne(r *reader) {

}

func (fp *fieldpath) PushThreePack5LeftDeltaN(r *reader) {

}

func (fp *fieldpath) PushN(r *reader) {

}

func (fp *fieldpath) PushNAndNonTopological(r *reader) {

}

func (fp *fieldpath) PopOnePlusOne(r *reader) {

}

func (fp *fieldpath) PopOnePlusN(r *reader) {

}

func (fp *fieldpath) PopAllButOnePlusOne(r *reader) {

}

func (fp *fieldpath) PopAllButOnePlusNPackN(r *reader) {

}

func (fp *fieldpath) PopAllButOnePlusNPack3Bits(r *reader) {

}

func (fp *fieldpath) PopAllButOnePlusNPack6Bits(r *reader) {

}

func (fp *fieldpath) PopNPlusOne(r *reader) {

}

func (fp *fieldpath) PopNPlusN(r *reader) {

}

func (fp *fieldpath) PopNAndNonTopographical(r *reader) {

}

func (fp *fieldpath) NonTopoComplex(r *reader) {

}

func (fp *fieldpath) NonTopoPenultimatePlusOne(r *reader) {

}

func (fp *fieldpath) NonTopoComplexPack4Bits(r *reader) {

}

func (fp *fieldpath) FieldPathEncodeFinish(r *reader) {

}
