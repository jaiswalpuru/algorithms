package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func diameter_of_binary_tree(root *TreeNode) int {
	return _diameter(root)
}

func _diameter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l_ht := height(root.Left)
	r_ht := height(root.Right)
	left := _diameter(root.Left)
	right := _diameter(root.Right)
	return max(l_ht+r_ht, max(left, right))
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := height(root.Left)
	r := height(root.Right)
	return 1 + max(l, r)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Right = New(3)
	fmt.Println(diameter_of_binary_tree(root))
}
