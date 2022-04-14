package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func search_in_binary_search_tree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	l := search_in_binary_search_tree(root.Left, val)
	r := search_in_binary_search_tree(root.Right, val)
	if l != nil {
		return l
	} else {
		return r
	}
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	root.Right = New(7)
	fmt.Println(search_in_binary_search_tree(root, 5))
}
