package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func delete_leaves_with_given_value(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = delete_leaves_with_given_value(root.Left, target)
	root.Right = delete_leaves_with_given_value(root.Right, target)
	if root.Val == target && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Left = New(2)
	root.Right = New(3)
	root.Right.Right = New(4)
	root.Right.Left = New(2)
	fmt.Println(delete_leaves_with_given_value(root, 2))
}
