package manta

import (
	"container/heap"
	"fmt"
)

// Interface for the tree, only implements Weight
type HuffmanTree interface {
	Weight() int
	IsLeaf() bool
}

// A leaf, contains encoded value
type HuffmanLeaf struct {
	weight int
	value  int
}

// A node with potential left / right nodes or leafs
type HuffmanNode struct {
	weight int
	left   HuffmanTree
	right  HuffmanTree
}

// Return weight for leaf
func (self HuffmanLeaf) Weight() int {
	return self.weight
}

// Return leaf state
func (self HuffmanLeaf) IsLeaf() bool {
	return true
}

// Return weight for node
func (self HuffmanNode) Weight() int {
	return self.weight
}

// Return leaf state
func (self HuffmanNode) IsLeaf() bool {
	return false
}

type treeHeap []HuffmanTree

// Returns the amount of nodes in the tree
func (th treeHeap) Len() int {
	return len(th)
}

// Weight compare function
func (th treeHeap) Less(i int, j int) bool {
	return th[i].Weight() < th[j].Weight()
}

// Append item, required for heap
func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(HuffmanTree))
}

// Remove item, required for heap
func (th *treeHeap) Pop() (popped interface{}) {
	popped = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}

// Swap two items, required for heap
func (th treeHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

// Construct a tree from a map of weight -> item
func buildTree(symFreqs []int) HuffmanTree {
	var trees treeHeap
	for v, w := range symFreqs {
		trees = append(trees, &HuffmanLeaf{w, v})
	}

	heap.Init(&trees)
	for trees.Len() > 1 {
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)

		heap.Push(&trees, &HuffmanNode{a.Weight() + b.Weight(), a, b})
	}

	return heap.Pop(&trees).(HuffmanTree)
}

// Swap two nodes based on the given path
func swapNodes(tree HuffmanTree, path uint32, len uint32) {
	for len > 0 {
		// get current bit
		len--
		one := path & 1
		path = path >> 1

		// check if we are correct
		if tree.IsLeaf() {
			_panicf("Touching leaf in node swap, %d left in path", len)
		}

		// switch on the type
		if one == 1 {
			tree = (tree.(*HuffmanNode).right)
		} else {
			tree = (tree.(*HuffmanNode).left)
		}
	}

	node := tree.(*HuffmanNode)
	tmp := node.left
	node.left = node.right
	node.right = tmp
}

// Print computed tree order
func printCodes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case *HuffmanLeaf:
		fmt.Printf("%v\t%d\t%d\t%s\n", i.value, i.weight, len(prefix), string(prefix))
	case *HuffmanNode:
		prefix = append(prefix, '0')
		printCodes(i.left, prefix)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		printCodes(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}
