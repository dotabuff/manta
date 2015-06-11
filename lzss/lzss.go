package lzss

// #include "./lzss.h"
// #cgo CFLAGS: -std=c11
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var magic = [4]byte{'L', 'Z', 'S', 'S'}

type Header struct {
	Magic [4]byte
	Size  int32
}

func Uncompress(input []byte) (output []byte, err error) {
	buf := bytes.NewBuffer(input)
	header := &Header{}
	err = binary.Read(buf, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	if header.Magic != magic {
		return nil, fmt.Errorf("header was `%s`, expected `%s`", header.Magic, magic)
	}

	if header.Size < int32(len(input)) {
		return nil, fmt.Errorf("assertion failed: len(input) < size: %d < %d", len(input), header.Size)
	}

	var totalBytes int32
	var cmdByte, getCmdByte byte
	var p, q, i byte

	for {
		if getCmdByte == 0 {
			cmdByte, err = buf.ReadByte()
			if err != nil {
				return nil, err
			}
		}

		getCmdByte = (getCmdByte + 1) & 0x07

		if cmdByte&0x01 != 0 {
			// int32_t position = *pInput++ << LZSS_LOOKSHIFT;
			position := cmdByte << 4

			if p, err = buf.ReadByte(); err != nil {
				return nil, err
			}

			// position |= ( *pInput >> LZSS_LOOKSHIFT );
			position |= (p >> 4)

			if q, err = buf.ReadByte(); err != nil {
				return nil, err
			}

			count := (q & 0x0f) + 1
			if count == 1 {
				break
			}

			for i = 0; i < count; i++ {
				offset := (len(output) - int(position) - 1) + int(i)
				source := output[offset]
				output = append(output, source)
			}
		} else {
			if p, err = buf.ReadByte(); err != nil {
				return nil, err
			}
			output = append(output, p)
		}

		cmdByte = cmdByte >> 1
	}

	if int32(len(output)) != header.Size {
		return nil, fmt.Errorf("Unexpected failure: total(%d) != actual(%d)", totalBytes, header.Size)
	}

	return output, nil
}

func UncompressReference(sample string) []byte {
	outRaw := ""
	out := C.CString(outRaw)
	in := C.CString(sample)
	f := C.lzss_uncompress(in, out, 99999)

	if f > 0 {
		return []byte(C.GoStringN(out, C.int(f)))
	} else {
		spew.Dump(f)
	}
	return nil
}
