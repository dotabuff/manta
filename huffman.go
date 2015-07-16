package manta

import (
	"container/heap"
	"fmt"
)

// Interface for the tree, only implements Weight
type HuffmanTree interface {
	Weight() int
	IsLeaf() bool
	Value() int
	Left() HuffmanTree
	Right() HuffmanTree
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

// Return value for leaf
func (self HuffmanLeaf) Value() int {
	return self.value
}

func (self HuffmanLeaf) Right() HuffmanTree {
	_panicf("HuffmanLeaf doesn't have right node")
	return nil
}

func (self HuffmanLeaf) Left() HuffmanTree {
	_panicf("HuffmanLeaf doesn't have left node")
	return nil
}

// Return weight for node
func (self HuffmanNode) Weight() int {
	return self.weight
}

// Return leaf state
func (self HuffmanNode) IsLeaf() bool {
	return false
}

// Return value for node
func (self HuffmanNode) Value() int {
	_panicf("HuffmanNode doesn't have a value")
	return 0
}

func (self HuffmanNode) Left() HuffmanTree {
	return HuffmanTree(self.left)
}

func (self HuffmanNode) Right() HuffmanTree {
	return HuffmanTree(self.right)
}

type treeHeap []HuffmanTree

// Returns the amount of nodes in the tree
func (th treeHeap) Len() int {
	return len(th)
}

// Weight compare function
func (th treeHeap) Less(i int, j int) bool {
	return th[i].Weight() <= th[j].Weight()
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
			tree = tree.Right()
		} else {
			tree = tree.Left()
		}
	}

	node := tree.(*HuffmanNode)
	node.left, node.right = node.right, node.left
}

// Print computed tree order
func printCodes(tree HuffmanTree, prefix []byte) {
	if tree.IsLeaf() {
		fmt.Printf("%v\t%d\t%d\t%s\n", tree.Value(), tree.Weight(), len(prefix), string(prefix))
	} else {
		prefix = append(prefix, '0')
		printCodes(tree.Left(), prefix)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		printCodes(tree.Right(), prefix)
		prefix = prefix[:len(prefix)-1]
	}
}
