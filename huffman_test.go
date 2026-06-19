package manta

import "testing"

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
