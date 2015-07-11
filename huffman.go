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
	value  interface{}
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
func buildTree(symFreqs map[int]interface{}) HuffmanTree {
	var trees treeHeap
	for w, v := range symFreqs {
		trees = append(trees, HuffmanLeaf{w, v})
	}

	heap.Init(&trees)
	for trees.Len() > 1 {
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)

		heap.Push(&trees, HuffmanNode{a.Weight() + b.Weight(), a, b})
	}

	return heap.Pop(&trees).(HuffmanTree)
}

// Print computed tree order
func printCodes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		fmt.Printf("%v\t%d\t%s\n", i.value, i.weight, string(prefix))
	case HuffmanNode:
		prefix = append(prefix, '0')
		printCodes(i.left, prefix)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		printCodes(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}
