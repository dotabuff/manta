package lzss

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

	for {
		if getCmdByte == 0 {
			cmdByte, err = buf.ReadByte()
			if err != nil {
				return nil, err
			}
		}

		getCmdByte = (getCmdByte + 1) & 0x07

		if cmdByte&0x01 == 0x01 {
			p, err := buf.ReadByte()
			if err != nil {
				return nil, err
			}

			position := p << 4

			q, err := buf.ReadByte()
			if err != nil {
				return nil, err
			}

			position |= (q >> 4)

			count := (q & 0x0f) + 1
			if count == 1 {
				break
			}

			source := byte(len(input)-buf.Len()) - position - 1
			for i := byte(0); i < count; i++ {
				output = append(output, source)
			}

			totalBytes += int32(count)
		} else {
			b, err := buf.ReadByte()
			if err != nil {
				return nil, err
			}
			output = append(output, b)
			totalBytes++
		}

		cmdByte = cmdByte >> 1
	}

	if totalBytes != header.Size {
		return nil, fmt.Errorf("Unexpected failure: total(%d) != actual(%d)", totalBytes, header.Size)
	}

	return output, nil
}
