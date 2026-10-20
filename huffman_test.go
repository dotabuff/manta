package manta

import (
	"math/rand"
	"testing"
)

// TestHuffmanFlatMatchesTree verifies the flattened field-path huffman arrays
// (huffTreeLeft/huffTreeRight) decode identically to the interface tree they are
// derived from, so the fast path can never silently desync from the canonical
// tree if the op table or weights change.
func TestHuffmanFlatMatchesTree(t *testing.T) {
	var walk func(node int32, itree huffmanTree)
	walk = func(node int32, itree huffmanTree) {
		if c := huffTreeLeft[node]; c < 0 {
			l := itree.Left()
			if !l.IsLeaf() {
				t.Fatalf("node %d left: flat=leaf op %d, tree=internal", node, -c-1)
			}
			if got := int32(l.Value()); got != -c-1 {
				t.Fatalf("node %d left: flat op %d != tree op %d", node, -c-1, got)
			}
		} else {
			if itree.Left().IsLeaf() {
				t.Fatalf("node %d left: flat=internal %d, tree=leaf op %d", node, c, itree.Left().Value())
			}
			walk(c, itree.Left())
		}

		if c := huffTreeRight[node]; c < 0 {
			r := itree.Right()
			if !r.IsLeaf() {
				t.Fatalf("node %d right: flat=leaf op %d, tree=internal", node, -c-1)
			}
			if got := int32(r.Value()); got != -c-1 {
				t.Fatalf("node %d right: flat op %d != tree op %d", node, -c-1, got)
			}
		} else {
			if itree.Right().IsLeaf() {
				t.Fatalf("node %d right: flat=internal %d, tree=leaf op %d", node, c, itree.Right().Value())
			}
			walk(c, itree.Right())
		}
	}
	walk(huffTreeRoot, huffTree)
}

// decodeOneHuffOpLookup decodes a single field-path op code (without executing
// the op) using the 8-bit lookup fast path.
func decodeOneHuffOpLookup(r *reader) int32 {
	entry := huffLookup[r.peekBits(huffLookupBits)]
	if consumed := entry & 0xFF; consumed != 0 {
		r.skipBits(uint32(consumed))
		return int32(entry >> 8)
	}
	r.skipBits(huffLookupBits)
	node := int32(entry >> 8)
	for {
		var child int32
		if r.readBits(1) == 1 {
			child = huffTreeRight[node]
		} else {
			child = huffTreeLeft[node]
		}
		if child < 0 {
			return -child - 1
		}
		node = child
	}
}

// decodeOneHuffOpWalk decodes a single field-path op code using only the flat
// tree walk.
func decodeOneHuffOpWalk(r *reader) int32 {
	node := huffTreeRoot
	for {
		var child int32
		if r.readBits(1) == 1 {
			child = huffTreeRight[node]
		} else {
			child = huffTreeLeft[node]
		}
		if child < 0 {
			return -child - 1
		}
		node = child
	}
}

// TestReadFieldPathsTruncated verifies that a truncated field-path op stream
// fails cleanly instead of underflowing the bit accumulator. An all-zero buffer
// decodes as an unbounded run of PlusOne ops (code "0") that never reaches
// FieldPathEncodeFinish; once the bits run out, the lookup would consume past
// the end of the buffer. With the skipBits guard this panics (the parser
// recovers it into an error); without it, bitCount underflows and the decode
// loops forever appending field paths.
func TestReadFieldPathsTruncated(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic on a truncated field-path stream")
		}
	}()
	readFieldPaths(newReader([]byte{0x00, 0x00}), nil)
}

// TestHuffmanLookupMatchesWalk verifies the 8-bit lookup fast path decodes the
// same op and consumes the same number of bits as the pure flat-tree walk for
// many random streams. An 8-byte buffer comfortably holds the longest code
// (17 bits), so neither path reaches the buffer end.
func TestHuffmanLookupMatchesWalk(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for trial := 0; trial < 5000; trial++ {
		buf := make([]byte, 8)
		for i := range buf {
			buf[i] = byte(rng.Intn(256))
		}
		ra := newReader(buf)
		rb := newReader(buf)
		opA := decodeOneHuffOpLookup(ra)
		opB := decodeOneHuffOpWalk(rb)
		consumedA := ra.pos*8 - ra.bitCount
		consumedB := rb.pos*8 - rb.bitCount
		if opA != opB || consumedA != consumedB {
			t.Fatalf("trial %d buf %x: lookup op=%d bits=%d, walk op=%d bits=%d", trial, buf, opA, consumedA, opB, consumedB)
		}
	}
}

// fieldPathOpCodes returns each op's huffman code as (bits, length), with the
// first bit read in the least significant position.
func fieldPathOpCodes() (codes [fpOpFieldPathEncodeFinish + 1]uint32, lens [fpOpFieldPathEncodeFinish + 1]uint32) {
	var walk func(node int32, code, depth uint32)
	walk = func(node int32, code, depth uint32) {
		for bit, child := range []int32{huffTreeLeft[node], huffTreeRight[node]} {
			c := code | uint32(bit)<<depth
			if child < 0 {
				codes[-child-1], lens[-child-1] = c, depth+1
			} else {
				walk(child, c, depth+1)
			}
		}
	}
	walk(huffTreeRoot, 0, 0)
	return
}

// bitWriter appends bits LSB-first, matching reader's bit order.
type bitWriter struct {
	buf []byte
	n   uint32
}

func (w *bitWriter) write(v uint32, n uint32) {
	for i := uint32(0); i < n; i++ {
		if w.n%8 == 0 {
			w.buf = append(w.buf, 0)
		}
		w.buf[w.n/8] |= byte((v>>i)&1) << (w.n % 8)
		w.n++
	}
}

func (w *bitWriter) truncate(n uint32) {
	w.buf = w.buf[:(n+7)/8]
	if n%8 != 0 {
		w.buf[n/8] &= byte(1)<<(n%8) - 1
	}
	w.n = n
}

// readerAtBit returns a reader over buf advanced by n bits.
func readerAtBit(buf []byte, n uint32) *reader {
	r := newReader(buf)
	for n > 0 {
		k := min(n, 32)
		r.readBits(k)
		n -= k
	}
	return r
}

// readFieldPathsReference is the canonical field-path decoder: a bit-by-bit
// tree walk applying each op through fieldPathTable's function values.
func readFieldPathsReference(r *reader) []fieldPath {
	var paths []fieldPath
	fp := newFieldPath()
	defer fp.release()
	for {
		op := decodeOneHuffOpWalk(r)
		if op == fpOpFieldPathEncodeFinish {
			return paths
		}
		fieldPathTable[op].fn(r, fp)
		paths = append(paths, *fp)
	}
}

// TestReadFieldPathsMatchesTable verifies readFieldPaths' switch dispatch
// applies the same op functions as fieldPathTable, by decoding streams built
// from random sequences of every op (with random operand bits) and comparing
// against readFieldPathsReference. Each op is appended only if the reference
// decoder applies it cleanly, which keeps the path depth in range.
func TestReadFieldPathsMatchesTable(t *testing.T) {
	codes, lens := fieldPathOpCodes()
	rng := rand.New(rand.NewSource(1))
	seen := make([]int, len(fieldPathTable))

	for trial := 0; trial < 2000; trial++ {
		w := &bitWriter{}
		fp := newFieldPath()
		for k := 0; k < 1+rng.Intn(40); k++ {
			for attempt := 0; attempt < 50; attempt++ {
				op := rng.Intn(fpOpFieldPathEncodeFinish)
				save := w.n
				w.write(codes[op], lens[op])
				for i := 0; i < 8; i++ {
					w.write(rng.Uint32(), 32)
				}
				r := readerAtBit(w.buf, save+lens[op])
				next := *fp
				ok := func() (ok bool) {
					defer func() { ok = recover() == nil }()
					fieldPathTable[op].fn(r, &next)
					return
				}()
				if ok && next.last >= 0 {
					w.truncate(r.pos*8 - r.bitCount)
					*fp = next
					seen[op]++
					break
				}
				w.truncate(save)
			}
		}
		fp.release()
		w.write(codes[fpOpFieldPathEncodeFinish], lens[fpOpFieldPathEncodeFinish])
		for i := rng.Intn(3); i > 0; i-- {
			w.write(rng.Uint32(), 8)
		}

		ra, rb := newReader(w.buf), newReader(w.buf)
		got := readFieldPaths(ra, nil)
		want := readFieldPathsReference(rb)
		if len(got) != len(want) {
			t.Fatalf("trial %d: %d paths, want %d", trial, len(got), len(want))
		}
		for i := range want {
			if got[i].last != want[i].last || got[i].path != want[i].path {
				t.Fatalf("trial %d path %d: got %v/%d, want %v/%d", trial, i, got[i].path, got[i].last, want[i].path, want[i].last)
			}
		}
		if a, b := ra.pos*8-ra.bitCount, rb.pos*8-rb.bitCount; a != b {
			t.Fatalf("trial %d: consumed %d bits, want %d", trial, a, b)
		}
	}

	for op := 0; op < fpOpFieldPathEncodeFinish; op++ {
		if seen[op] == 0 {
			t.Errorf("op %s never exercised", fieldPathTable[op].name)
		}
	}
}
