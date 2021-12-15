package main

type BinarySearchTreeNode struct {
	data        int
	left, right *BinarySearchTreeNode
}

func NewNode(data int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{data: data, left: nil, right: nil}
}
