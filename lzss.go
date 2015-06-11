package manta

// Decompress a Valve LZSS compressed buffer
func unlzss(buf []byte) ([]byte, error) {
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

			position := (int(a) << 4) | (int(b) >> 4)
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
