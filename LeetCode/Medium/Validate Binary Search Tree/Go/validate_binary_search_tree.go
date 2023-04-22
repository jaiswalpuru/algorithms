package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func _validate_bst(root, left, right *TreeNode) bool {
	if root == nil {
		return true
	}

	if left != nil && left.Val >= root.Val {
		return false
	}

	if right != nil && right.Val <= root.Val {
		return false
	}

	return _validate_bst(root.Left, left, root) && _validate_bst(root.Right, root, right)
}

func validate_bst(root *TreeNode) bool {
	return _validate_bst(root, nil, nil)
}

func main() {
	root := New(2)
	root.Left = New(1)
	root.Right = New(3)
	fmt.Println(validate_bst(root))
}
