package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func buildTree(inorder []int, postorder []int) *TreeNode {
	inIndex := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inIndex[inorder[i]] = i
	}
	postIndex := len(postorder) - 1
	return recur(0, len(inorder)-1, inorder, postorder, inIndex, &postIndex)
}

func recur(left, right int, inorder, postorder []int, inIndex map[int]int, postIndex *int) *TreeNode {
	if left > right {
		return nil
	}

	root := New(postorder[*postIndex])
	ind := inIndex[root.Val]
	*postIndex--
	root.Right = recur(ind+1, right, inorder, postorder, inIndex, postIndex)
	root.Left = recur(left, ind-1, inorder, postorder, inIndex, postIndex)
	return root
}

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
