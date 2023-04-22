package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := []*TreeNode{root}
	nullNode := false

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr == nil {
			nullNode = true
		} else {
			if nullNode {
				return false
			}
			q = append(q, curr.Left)
			q = append(q, curr.Right)
		}
	}

	return true
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Right.Left = New(6)
	fmt.Println(isCompleteTree(root))
}
