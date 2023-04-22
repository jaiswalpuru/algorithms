package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func same_tree(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	if r1.Val != r2.Val {
		return false
	}
	return same_tree(r1.Left, r2.Left) && same_tree(r1.Right, r2.Right)
}

func main() {
	r1 := New(1)
	r1.Left = New(2)
	r1.Right = New(3)
	r2 := New(1)
	r2.Left = New(2)
	r2.Right = New(3)
	fmt.Println(same_tree(r1, r2))
}
