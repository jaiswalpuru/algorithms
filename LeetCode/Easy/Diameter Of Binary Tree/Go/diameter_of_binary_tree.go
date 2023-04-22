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
	diameter := 0
	_diameter(root, &diameter)
	return diameter
}

func _diameter(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	l := _diameter(root.Left, diameter)
	r := _diameter(root.Right, diameter)
	*diameter = max(*diameter, l+r)
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
