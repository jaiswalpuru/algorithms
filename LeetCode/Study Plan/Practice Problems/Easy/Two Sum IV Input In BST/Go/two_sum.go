package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func two_sum(root *TreeNode, k int) bool {
	seen := make(map[int]bool)
	return inorder(root, &seen, k)
}

func inorder(root *TreeNode, seen *map[int]bool, k int) bool {
	if root == nil {
		return false
	}
	if (*seen)[k-root.Val] {
		return true
	}
	(*seen)[root.Val] = true
	return inorder(root.Left, seen, k) || inorder(root.Right, seen, k)
}

func main() {
	root := New(5)
	root.Left = New(3)
	root.Left.Left = New(2)
	root.Left.Right = New(4)
	root.Right = New(6)
	root.Right.Right = New(7)
	fmt.Println(two_sum(root, 9))
}
