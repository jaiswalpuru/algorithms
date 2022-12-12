package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func successor(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Right == nil {
		return root.Val
	}
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Right.Left = New(4)
	fmt.Println(successor(root))
}
