package main

import (
	"fmt"
	"strconv"
)

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	visited := make(map[string]int)
	res := []*TreeNode{}
	dfs(root, &visited, &res)

	return res
}

func dfs(root *TreeNode, visited *map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	str := strconv.Itoa(root.Val) + ":" + dfs(root.Left, visited, res) + ":" + dfs(root.Right, visited, res)
	if (*visited)[str] == 1 {
		*res = append(*res, root)
	}
	(*visited)[str]++
	return str
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Left = New(4)
	root.Right = New(3)
	root.Right.Left = New(2)
	root.Right.Left.Left = New(4)
	root.Right.Right = New(4)
	fmt.Println(findDuplicateSubtrees(root))
}
