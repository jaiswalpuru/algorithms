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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return recur(root.Left, root.Right)
}

func recur(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(2)
	root.Left.Left = New(3)
	root.Left.Right = New(4)
	root.Right.Left = New(4)
	root.Right.Right = New(3)
	fmt.Println(isSymmetric(root))
}
