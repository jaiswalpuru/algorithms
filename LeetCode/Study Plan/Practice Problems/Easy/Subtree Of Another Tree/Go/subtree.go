package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func is_subtree(root *TreeNode, sub_tree *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == sub_tree.Val {
		if _is_identical(root, sub_tree) {
			return true
		}
	}
	return is_subtree(root.Left, sub_tree) || is_subtree(root.Right, sub_tree)
}

func _is_identical(root, sub_tree *TreeNode) bool {
	if root == nil && sub_tree == nil {
		return true
	}
	if root == nil || sub_tree == nil {
		return false
	}
	if root.Val != sub_tree.Val {
		return false
	}
	return _is_identical(root.Left, sub_tree.Left) && _is_identical(root.Right, sub_tree.Right)
}

func main() {
	root := New(3)
	root.Left = New(4)
	root.Right = New(5)
	root.Left.Left = New(1)
	root.Left.Right = New(2)
	subtree := New(4)
	subtree.Left = New(1)
	subtree.Right = New(2)
	fmt.Println(is_subtree(root, subtree))
}
