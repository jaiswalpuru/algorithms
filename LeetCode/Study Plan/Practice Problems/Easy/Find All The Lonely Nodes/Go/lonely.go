package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func lonely(r *TreeNode) []int {
	res := []int{}
	_lonely(r, &res)
	return res
}

func _lonely(r *TreeNode, res *[]int) {
	if r == nil {
		return
	}

	if r.Left == nil && r.Right != nil {
		*res = append(*res, r.Right.Val)
	}

	if r.Left != nil && r.Right == nil {
		*res = append(*res, r.Left.Val)
	}

	_lonely(r.Left, res)
	_lonely(r.Right, res)
}

func main() {
	r := New(1)
	r.Left = New(2)
	r.Right = New(3)
	r.Left.Right = New(4)
	fmt.Println(lonely(r))
}
