package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func is_balanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(height(root.Left)-height(root.Right)) > 1 {
		return false
	} else {
		return is_balanced(root.Left) && is_balanced(root.Right)
	}
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(is_balanced(root))
}
