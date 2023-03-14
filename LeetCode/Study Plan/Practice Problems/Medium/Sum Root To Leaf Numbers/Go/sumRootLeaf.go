package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	sum := 0
	recur(root, 0, &sum)
	return sum
}

func recur(root *TreeNode, num int, sum *int) {
	if root == nil {
		return
	}

	num = (num * 10) + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += num
	}
	recur(root.Left, num, sum)
	recur(root.Right, num, sum)
}

func main() {
	root := New(1)
	root.Left = New(0)
	// root.Right = New(0)
	// root.Left.Left = New(5)
	// root.Left.Right = New(1)
	fmt.Println(sumNumbers(root))
}
