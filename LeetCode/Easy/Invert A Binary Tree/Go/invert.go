package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func invert(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	left, right := invert(root.Left), invert(root.Right)
	root.Left = left
	root.Right = right
	return root
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(7)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	root.Right.Left = New(6)
	root.Right.Right = New(9)
	invert(root)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}
