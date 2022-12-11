package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func validate_bst(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val > root.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val < root.Val {
			return false
		}
	}
	return validate_bst(root.Left) && validate_bst(root.Right)
}

func main() {
	root := New(1)
	root.Left = New(0)
	root.Right = New(3)
	fmt.Println(validate_bst(root))
}
