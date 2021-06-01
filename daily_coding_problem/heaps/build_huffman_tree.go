/*
	Huffman coding is a method of encoding characters based on their frequency. Each letter is assigned a variable
	length binary string, such as 0101 or 111110, where shorter lengths.
	Only the leaf nodes have letters.
	Given a dictionary of characters frequencies, build a Huffman tree, and use it to determine a mapping between
	characters and their encoded binary strings.
*/

package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	char  rune
	left  *Node
	right *Node
}

type Pair struct {
	char rune
	cnt  int
}

type NodeFreq struct {
	freq int
	node *Node
}

type MinHeap []NodeFreq

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(val interface{}) { *h = append(*h, val.(NodeFreq)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func NewNode(letter rune, left, right *Node) *Node {
	return &Node{char: letter, left: left, right: right}
}

func build_tree(freq []Pair) *Node {
	nodes := make(MinHeap, 0)
	heap.Init(&nodes)

	for i := range freq {
		heap.Push(&nodes, NodeFreq{freq: freq[i].cnt, node: NewNode(freq[i].char, nil, nil)})
	}

	for len(nodes) > 1 {
		f1 := heap.Pop(&nodes).(NodeFreq)
		f2 := heap.Pop(&nodes).(NodeFreq)
		node := NewNode('*', f1.node, f2.node)
		heap.Push(&nodes, NodeFreq{freq: f1.freq + f2.freq, node: node})
	}

	root := nodes[0].node
	return root
}

func encode(root *Node, str string, mapping *map[rune]string) {

	if root == nil {
		return
	}

	if root.left == nil && root.right == nil {
		(*mapping)[root.char] = str
	}

	encode(root.left, str+"0", mapping)
	encode(root.right, str+"1", mapping)
}

func in_order(root *Node) {
	if root == nil {
		return
	}
	in_order(root.left)
	fmt.Println(string(root.char))
	in_order(root.right)
}

func main() {
	let_freq := []Pair{
		{'a', 3},
		{'c', 6},
		{'e', 8},
		{'f', 2},
	}

	root := build_tree(let_freq)
	//in_order(root)
	mp := make(map[rune]string)
	encode(root, "", &mp)
	for k, v := range mp {
		fmt.Println(string(k), "->", v)
	}
}
