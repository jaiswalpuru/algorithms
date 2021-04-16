package main

type BinaryTree struct {
	data        int
	left, right *BinaryTree
}

func NewBT(data int) *BinaryTree {
	return &BinaryTree{data, nil, nil}
}
