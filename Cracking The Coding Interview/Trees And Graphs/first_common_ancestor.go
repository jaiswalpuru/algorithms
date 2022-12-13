package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func lca(root *TreeNode, n1, n2 int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == n1 || root.Val == n2 {
		return root
	}
	lca_l := lca(root.Left, n1, n2)
	lca_r := lca(root.Right, n1, n2)
	if lca_l != nil && lca_r != nil {
		return root
	}
	if lca_l != nil {
		return lca_l
	}
	return lca_r
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Right.Left = New(6)
	root.Right.Right = New(7)
	fmt.Println(lca(root, 4, 5))
}
