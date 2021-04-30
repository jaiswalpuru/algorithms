package main

type BinaryTreeNode struct {
	data        int
	left, right *BinaryTree
}

func NewBT(data int) *BinaryTreeNode {
	return &BinaryTreeNode{data, nil, nil}
}
