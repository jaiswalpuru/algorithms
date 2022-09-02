package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func root_equal_sum_of_children(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}

func main() {
	root := New(10)
	root.Left = New(4)
	root.Right = New(6)
	fmt.Println(root_equal_sum_of_children(root))
}
