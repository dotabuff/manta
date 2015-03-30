package manta

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	CoordIntegerBits    = 14
	CoordFractionalBits = 5
	CoordDenominator    = (1 << CoordFractionalBits)
	CoordResolution     = (1.0 / CoordDenominator)

	NormalFractionalBits = 11
	NormalDenominator    = ((1 << NormalFractionalBits) - 1)
	NormalResolution     = (1.0 / NormalDenominator)
)

type BitReader struct {
	buffer     []byte
	currentBit int
}

func NewBitReader(buffer []byte) *BitReader {
	if len(buffer) == 0 {
		panic("empty buffer?")
	}
	return &BitReader{buffer: buffer}
}

func (br *BitReader) Length() int      { return len(br.buffer) }
func (br *BitReader) CurrentBit() int  { return br.currentBit }
func (br *BitReader) CurrentByte() int { return br.currentBit / 8 }
func (br *BitReader) BitsLeft() int    { return (len(br.buffer) * 8) - br.currentBit }
func (br *BitReader) BytesLeft() int   { return len(br.buffer) - (br.currentBit / 8) }

type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) String() string {
	return fmt.Sprintf("{{ x: %f, y: %f, z: %f }}", v.X, v.Y, v.Z)
}

type Vector2 struct {
	X, Y float64
}

func (v Vector2) String() string {
	return fmt.Sprintf("{{ x: %f, y: %f }}", v.X, v.Y)
}

type SeekOrigin int

const (
	Current SeekOrigin = iota
	Begin
	End
)

func (br *BitReader) SeekBits(offset int, origin SeekOrigin) {
	if origin == Current {
		br.currentBit += offset
	} else if origin == Begin {
		br.currentBit = offset
	} else if origin == End {
		br.currentBit = (len(br.buffer) * 8) - offset
	}
	if br.currentBit < 0 || br.currentBit > (len(br.buffer)*8) {
		panic("out of range")
	}
}

func (br *BitReader) ReadUBitsByteAligned(nBits int) uint {
	if nBits%8 != 0 {
		panic("Must be multple of 8")
	}

	if br.currentBit%8 != 0 {
		panic("Current bit is not byte-aligned")
	}

	var result uint
	for i := 0; i < nBits/8; i++ {
		result += uint(br.buffer[br.CurrentByte()] << (uint(i) * 8))
		br.currentBit += 8
	}
	return result
}

func (br *BitReader) ReadUBitsNotByteAligned(nBits int) uint {
	bitOffset := br.currentBit % 8
	nBitsToRead := bitOffset + nBits
	nBytesToRead := nBitsToRead / 8
	if nBitsToRead%8 != 0 {
		nBytesToRead += 1
	}

	var currentValue uint64
	for i := 0; i < nBytesToRead; i++ {
		b := br.buffer[br.CurrentByte()+i]
		currentValue += (uint64(b) << (uint64(i) * 8))
	}
	currentValue >>= uint(bitOffset)
	currentValue &= ((1 << uint64(nBits)) - 1)
	br.currentBit += nBits
	return uint(currentValue)
}

func (br *BitReader) ReadVarInt() (result uint) {
	var b uint
	count := 0

	for {
		if count == 5 {
			return result
		} else if br.CurrentByte() >= len(br.buffer) {
			return result
		}

		b = br.ReadUBits(8)
		result |= (b & 0x7f) << uint(7*count)
		count++

		if (b & 0x80) != 0x80 {
			break
		}
	}

	return result
}

func (br *BitReader) ReadUBits(nBits int) uint {
	if nBits <= 0 || nBits > 32 {
		panic("Value must be a positive integer between 1 and 32 inclusive.")
	}
	if (br.currentBit + nBits) > (len(br.buffer) * 8) {
		panic("Out of range")
	}
	if br.currentBit%8 == 0 && nBits%8 == 0 {
		return br.ReadUBitsByteAligned(nBits)
	}
	return br.ReadUBitsNotByteAligned(nBits)
}

func (br *BitReader) ReadBitsAsBytes(n int) []byte {
	result := make([]byte, (n+7)/8)
	i := 0
	for n > 7 {
		n -= 8
		result[i] = byte(br.ReadUBits(8))
		i++
	}
	if n != 0 {
		result[i] = byte(br.ReadUBits(n))
	}

	return result
}

func (br *BitReader) ReadBits(nBits int) int {
	result := br.ReadUBits(nBits - 1)
	if br.ReadBoolean() {
		result = -((1 << (uint(nBits) - 1)) - result)
	}
	return int(result)
}

func (br *BitReader) ReadBoolean() bool {
	if br.CurrentBit()+1 > br.Length()*8 {
		panic("Out of range")
	}
	currentByte := br.currentBit / 8
	bitOffset := br.currentBit % 8
	result := br.buffer[currentByte]&(1<<uint(bitOffset)) != 0
	br.currentBit++
	return result
}

func (br *BitReader) ReadByte() byte {
	return byte(br.ReadUBits(8))
}

func (br *BitReader) ReadBytes(nBytes uint) []byte {
	result := make([]byte, nBytes)
	for i := uint(0); i < nBytes; i++ {
		result[i] = br.ReadByte()
	}
	return result
}

func (br *BitReader) ReadBitFloat() float32 {
	b := bytes.NewBuffer(br.ReadBytes(4))
	var f float32
	binary.Read(b, binary.LittleEndian, &f)
	return f
}

func (br *BitReader) ReadBitNormal() float32 {
	signbit := br.ReadBoolean()
	fractval := float32(br.ReadUBits(NormalFractionalBits))
	value := fractval * NormalResolution
	if signbit {
		value = -value
	}
	return value
}

func (br *BitReader) ReadBitCellCoord(bits int, integral, lowPrecision bool) (value float32) {
	if integral {
		value = float32(br.ReadBits(bits))
	} else {
		intval := br.ReadBits(bits)
		if lowPrecision {
			fractval := float32(br.ReadBits(3))
			value = float32(intval) + (fractval * (1.0 / (1 << 3)))
		} else {
			fractval := float32(br.ReadBits(5))
			value = float32(intval) + (fractval * (1.0 / (1 << 5)))
		}
	}
	return value
}

func (br *BitReader) ReadBitCoord() (value float32) {
	intFlag := br.ReadBoolean()
	fractFlag := br.ReadBoolean()
	if intFlag || fractFlag {
		negative := br.ReadBoolean()
		if intFlag {
			value += float32(br.ReadUBits(CoordIntegerBits)) + 1
		}
		if fractFlag {
			value += float32(br.ReadUBits(CoordFractionalBits)) * CoordResolution
		}
		if negative {
			value = -value
		}
	}
	return value
}

func (br *BitReader) ReadString() string {
	bs := []byte{}
	for {
		b := br.ReadByte()
		if b == 0 {
			break
		}
		bs = append(bs, b)
	}
	return string(bs)
}

func (br *BitReader) ReadStringN(n int) string {
	buf := []byte{}
	for n > 0 {
		c := br.ReadByte()
		if c == 0 {
			break
		}
		buf = append(buf, c)
		n--
	}
	return string(buf)
}

func (br *BitReader) ReadLengthPrefixedString() string {
	stringLength := uint(br.ReadUBits(9))
	if stringLength > 0 {
		return string(br.ReadBytes(stringLength))
	}
	return ""
}

func (br *BitReader) ReadNextEntityIndex(oldEntity int) int {
	ret := br.ReadUBits(4)
	more1 := br.ReadBoolean()
	more2 := br.ReadBoolean()
	if more1 {
		ret += (br.ReadUBits(4) << 4)
	}
	if more2 {
		ret += (br.ReadUBits(8) << 4)
	}
	return oldEntity + 1 + int(ret)
}

// [2 bits header][X bits type][varint size][8*size content]
// X = (header 00 = 4; header 01 = 8; header 10 = 12).
// Die ersten 4 bits vor den Rest pasten, oder via bitmask machen.
// Du musst 6 bits lesen, die ersten beiden davon entscheiden wie viel du dazu lesen musst
// also [XXYYYY]
// die ersten 4 musst du vor den rest, den du liest, pushen
// also wenn du danach noch mal vier liest musst du das so machen:
// [nimm 4 untere bits von den 6 die du liest], xor (4 neue bits gelesen und nach vorne geshifted)
func (br *BitReader) ReadInnerPacket() (int32, []byte) {
	initial := br.ReadUBits(6)
	header := initial >> 4
	var pType uint
	if header == 0 {
		pType = initial
	} else {
		pType = (initial & 15) | (br.ReadUBits(int(header*4+(((2-header)>>31)&16))) << 4)
	}

	pLen := br.ReadVarInt()
	pBytes := br.ReadBytes(pLen)
	return int32(pType), pBytes
}
