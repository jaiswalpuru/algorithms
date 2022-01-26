package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//To do it iteratively just use BFS and check left.Left.Val == right.Right.Val && left.Right.Val == right.Left.Val and
//push into queue

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func recur_is_symmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return _recur_is_symmetric(root.Left, root.Right)
}

func _recur_is_symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return _recur_is_symmetric(left.Left, right.Right) && _recur_is_symmetric(left.Right, right.Left)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(2)
	root.Left.Left = New(3)
	root.Left.Right = New(4)
	root.Right.Left = New(4)
	root.Right.Right = New(3)
	fmt.Println(recur_is_symmetric(root))
}
