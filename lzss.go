package manta

import (
	"encoding/binary"
	"fmt"
)

const (
	LZSS_WINDOW_SIZE      = 4096
	LZSS_MAX_STORE_LENGTH = 18 // 16?
	LZSS_LOOKSHIFT        = 4
	LZSS_THRESHOLD        = 1
)

/*
// Ultimate bad idea version

func decompress4(buf []byte) ([]byte, error) {
	N := 4096
	F := 16
	LOOKAHEAD := (1 << 4)

	text_buf := make([]byte, N + F - 1)

	var i, j, k, r, c int
	var flags uint

	for i = 0 ; i < N - F; i++ {
		text_buf[i] = ' '
	}
	r = N - F
	flags = 0

	b := newReader(buf)
	b.seekBytes(8)
	out := make([]byte, 0)

	eof := func() bool {
		return b.remBytes() < 1
	}

	for {
		flags >>= 1
		if (flags & 256) == 0 {
			if b.eof() {
				break
			}

			flags = int(b.readByte()) | 0xff00
		}
		if (flags & 1) != 0 {
			if b.eof() {
				break
			}
			c = int(b.readByte())
			out = append(out, byte(c))
			text_buf[r] = byte(c)
			r++
			r &= (N - 1)
		} else {
			if b.eof() {
				break
			}
			i = int(b.readByte())
			j = int(b.readByte())
			i |= ((j & 0xf0) << 4)
			j = (j & 0x0f) + 2
		}
	}
}
*/

// Work in progress version
func decompress3(buf []byte) ([]byte, error) {
	r := newReader(buf)

	if s := r.readStringN(4); s != "LZSS" {
		return nil, _errorf("expected LZSS header, got %s", s)
	}

	size := int(r.readLeUint32())
	out := make([]byte, 0)
	ring := make([]byte, LZSS_WINDOW_SIZE+LZSS_MAX_STORE_LENGTH-1)
	R := LZSS_WINDOW_SIZE - LZSS_MAX_STORE_LENGTH
	for i := 0; i < R; i++ {
		ring[i] = ' '
	}

	var cmdByte, getCmdByte byte

	for {
		if getCmdByte == 0 {
			cmdByte = r.readByte()
		}

		getCmdByte = (getCmdByte + 1) & 0x07

		if (cmdByte & 0x01) != 0 {
			c := r.readBytes(2)
			position := int(c[0] | ((c[1] & 0xF) << LZSS_LOOKSHIFT))
			count := int(c[1]&0x0F) + LZSS_THRESHOLD
			if count == 1 {
				break
			}

			for i := 0; i < count; i++ {
				c := ring[position+i&(LZSS_WINDOW_SIZE-1)]
				ring[R] = c
				out = append(out, c)
				R = (R + 1) & (LZSS_WINDOW_SIZE - 1)
			}

		} else {
			c := r.readByte()
			ring[R] = c
			R = (R + 1) & (LZSS_WINDOW_SIZE - 1)
			out = append(out, c)
		}
		cmdByte = cmdByte >> 1
	}

	if len(out) != size {
		return nil, _errorf("expected %d bytes, got %d", size, len(out))
	}

	return out, nil
}

// Stable, mostly working version.
func decompress2(buf []byte) ([]byte, error) {
	r := newReader(buf)

	if s := r.readStringN(4); s != "LZSS" {
		return nil, _errorf("expected LZSS header, got %s", s)
	}

	size := int(r.readLeUint32())
	out := make([]byte, 0)

	var cmdByte, getCmdByte byte

	for {
		if getCmdByte == 0 {
			cmdByte = r.readByte()
		}
		getCmdByte = (getCmdByte + 1) & 0x07

		if (cmdByte & 0x01) != 0 {
			a := r.readByte()
			b := r.readByte()

			position := int((a << LZSS_LOOKSHIFT) | (b >> LZSS_LOOKSHIFT))
			count := int((b & 0x0F) + 1)
			if count == 1 {
				break
			}
			source := len(out) - int(position) - 1
			for i := 0; i < count; i++ {
				out = append(out, out[source+i])
			}
		} else {
			out = append(out, r.readByte())
		}
		cmdByte = cmdByte >> 1
	}

	if len(out) != size {
		return nil, _errorf("expected %d bytes, got %d", size, len(out))
	}

	return out, nil
}

// Stable, previous "best reproduction" of a baseline that appears to be very wrong.
func decompress1(in []byte) ([]byte, error) {
	var cmdByte, getCmdByte int
	var inPos, outPos int

	if string(in[0:4]) != "LZSS" {
		return nil, fmt.Errorf("bad header")
	}
	inPos += 4

	size := int(binary.LittleEndian.Uint32(in[4:8]))
	inPos += 4

	_debugf("size is %d", size)

	out := make([]byte, size)

	for {
		if getCmdByte == 0 {
			_debugf("%d: cmdByte = %x", outPos, in[inPos])
			cmdByte = int(in[inPos])
			inPos += 1

		}

		getCmdByte = (getCmdByte + 1) & 0x07

		if (cmdByte & 0x01) != 0 {
			// Read first byte, get shifted value
			position := uint(in[inPos]) << LZSS_LOOKSHIFT
			inPos += 1

			// Read second byte
			extra := uint(in[inPos])
			inPos += 1

			// Update position with shifted value of second byte
			position |= (extra >> LZSS_LOOKSHIFT)

			// Get count with shifted value of second byte
			count := 1 + (extra & 0x0F)
			if count == 1 {
				break
			}

			srcPos := int(uint(outPos) - position - 1)
			_debugf("%d: position = %d, count = %d, source = %d", outPos, position, count, srcPos)
			for i := 0; i < int(count); i++ {
				_debugf("%d: writing %x", outPos, out[srcPos])
				out[outPos] = out[srcPos]
				outPos += 1
				srcPos += 1
			}

		} else {
			out[outPos] = in[inPos]
			_debugf("%d: copy %x (%d / %d)", outPos, out[outPos], int(out[outPos]), uint(out[outPos]))
			outPos += 1
			inPos += 1
		}

		cmdByte = cmdByte >> 1
	}

	out = out[:outPos]

	if len(out) != size {
		return nil, fmt.Errorf("expected %d, got %d", size, len(out))
	}

	return out, nil
}
