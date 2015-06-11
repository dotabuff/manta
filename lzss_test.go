package manta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompressLZSS(t *testing.T) {
	assert := assert.New(t)

	fixtures := []string{
		"18726",
		"22356",
		"4162",
	}

	for _, f := range fixtures {
		compressed := _read_fixture(_sprintf("lzss/%s_compressed", f))
		expect := _read_fixture(_sprintf("lzss/%s_decompressed", f))
		got, err := unlzss(compressed)
		assert.Nil(err)
		assert.True(byteCompareVerbose(t, expect, got))
	}
}

// Compares two byte buffers,
func byteCompareVerbose(t *testing.T, bufA, bufB []byte) bool {
	getByte := func(buf []byte, pos int) byte {
		if pos > len(buf) {
			return 0x00
		}
		return buf[pos]
	}

	size := len(bufA)
	if len(bufB) > size {
		size = len(bufB)
	}

	matched := 0
	for i := 0; i < size; i++ {
		a := getByte(bufA, i)
		b := getByte(bufB, i)
		if a != b {
			t.Logf("byte %d: A=%02x, B=%02x", i, a, b)
		} else {
			matched += 1
		}
	}

	if matched != size {
		t.Logf("%d/%d bytes matched", matched, size)
		return false
	}

	return true
}
