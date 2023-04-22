package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode {return &TreeNode{Val:val, Left : nil, Right : nil}}

func split_bst(root *TreeNode, target int) []*TreeNode {
	if root == nil {
		return []*TreeNode{nil, nil}
	}else if root.Val <= target {
		bst := split_bst(root.Right, target)
		root.Right = bst[0]
		bst[0] = root
		return bst
	}else {
		bst := split_bst(root.Left, target)
		root.Left = bst[1]
		bst[1] = root
		return bst
	}
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(6)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	root.Right.Left = New(5)
	root.Right.Right = New(7)
	fmt.Println(split_bst(root, 2))
}