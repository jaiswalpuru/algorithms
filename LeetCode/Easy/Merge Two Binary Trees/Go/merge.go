package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func merge(r1, r2 *TreeNode) *TreeNode {
	return _merge(r1, r2)

}

func _merge(r1, r2 *TreeNode) *TreeNode {
	if r1 == nil {
		return r2
	}

	if r2 == nil {
		return r1
	}

	r1.Val += r2.Val
	r1.Left = _merge(r1.Left, r2.Left)
	r1.Right = _merge(r1.Right, r2.Right)
	return r1
}

func main() {
	r1 := New(1)
	r1.Left = New(3)
	r1.Right = New(2)
	r1.Left.Left = New(5)
	r2 := New(2)
	r2.Left = New(1)
	r2.Right = New(3)
	r2.Left.Right = New(4)
	r2.Right.Right = New(7)
	fmt.Println(merge(r1, r2))
}
