package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func check_subtree(root *TreeNode, sub *TreeNode) bool {
	node := find(root, sub.Val)
	if node == nil {
		return false
	}
	return match(node, sub)
}

func find(r1 *TreeNode, val int) *TreeNode {
	if r1 == nil {
		return nil
	}
	if r1.Val == val {
		return r1
	}
	v1 := find(r1.Left, val)
	v2 := find(r1.Right, val)
	if v1 == nil {
		return v2
	}
	return v1
}

func match(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.Val != r2.Val {
		return false
	}
	return match(r1.Left, r2.Left) && match(r1.Right, r2.Right)
}

func main() {
	r := New(50)
	r.Left = New(20)
	r.Right = New(60)
	r.Left.Left = New(10)
	r.Left.Right = New(25)
	r.Left.Left.Right = New(15)
	r.Left.Left.Left = New(5)
	r.Right.Right = New(70)
	r.Right.Right.Left = New(65)
	r.Right.Right.Right = New(80)
	sub := New(90)
	fmt.Println(check_subtree(r, sub))
}
