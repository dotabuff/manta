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
