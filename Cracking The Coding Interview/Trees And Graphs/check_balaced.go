package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func get_height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(get_height(root.Left), get_height(root.Right))
}

func check_balanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := get_height(root.Left)
	r := get_height(root.Right)
	if abs(l-r) > 1 {
		return false
	}
	return check_balanced(root.Left) && check_balanced(root.Right)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Right.Right = New(4)
	root.Right.Right.Right = New(5)
	fmt.Println(check_balanced(root))
}
