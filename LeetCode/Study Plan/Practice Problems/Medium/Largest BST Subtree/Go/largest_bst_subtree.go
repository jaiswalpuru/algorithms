package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largest_bst_subtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if is_valid_bst(root, nil, nil) {
		return count_node(root)
	}

	return max(largest_bst_subtree(root.Left), largest_bst_subtree(root.Right))
}

func count_node(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + count_node(root.Left) + count_node(root.Right)
}

func is_valid_bst(root, left, right *TreeNode) bool {
	if root == nil {
		return true
	}

	if left != nil && left.Val >= root.Val {
		return false
	}

	if right != nil && right.Val <= root.Val {
		return false
	}

	return is_valid_bst(root.Left, left, root) && is_valid_bst(root.Right, root, right)
}

func main() {
	root := New(10)
	root.Left = New(5)
	root.Right = New(15)
	root.Left.Left = New(1)
	root.Left.Right = New(8)
	root.Right.Right = New(7)
	fmt.Println(largest_bst_subtree(root))
}
