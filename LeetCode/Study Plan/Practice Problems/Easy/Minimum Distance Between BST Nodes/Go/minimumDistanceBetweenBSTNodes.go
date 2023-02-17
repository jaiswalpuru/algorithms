package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
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
func minDiffInBST(root *TreeNode) int {
	sortedElements := []int{}
	inorder(&sortedElements, root)

	minDiff := sortedElements[1] - sortedElements[0]
	for i := 2; i < len(sortedElements); i++ {
		minDiff = min(minDiff, sortedElements[i]-sortedElements[i-1])
	}
	return minDiff
}

func inorder(sortedElements *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorder(sortedElements, root.Left)
	*sortedElements = append(*sortedElements, root.Val)
	inorder(sortedElements, root.Right)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(6)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	fmt.Println(minDiffInBST(root))
}
