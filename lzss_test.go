package manta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompressLZSS(t *testing.T) {
	assert := assert.New(t)

	fixtures := []string{"4162", "18726", "22356"}

	for _, f := range fixtures {
		// Read a compressed fixture
		compressed := _read_fixture(_sprintf("lzss/%s_compressed", f))

		res1, err := decompress2(compressed)
		assert.Nil(err)

		res2, err := decompress3(compressed)
		assert.Nil(err)

		assert.True(byteCompareVerbose(t, res1, res2))

		/*
			// Read a reference fixture decompressed with LZSS.CPP
			reference := _read_fixture(_sprintf("lzss/%s_reference_out", f))

			// Decompress using our internal lib and write a comparison fixture
			result, err := decompress(compressed)
			_dump_fixture(_sprintf("lzss/%s_internal_out", f), result)

			assert.Nil(err)
			assert.Equal(reference, result)
		*/
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
