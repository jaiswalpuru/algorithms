package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func binary_tree_pruning(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	post_order(root, root)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func post_order(root, prev *TreeNode) {
	if root == nil {
		return
	}
	post_order(root.Left, root)
	post_order(root.Right, root)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		if prev.Left == root {
			prev.Left = nil
		}
		if prev.Right == root {
			prev.Right = nil
		}
	}
}

func main() {
	root := New(1)
	root.Right = New(0)
	root.Right.Left = New(1)
	root.Right.Right = New(1)
	fmt.Println(binary_tree_pruning(root))
}
