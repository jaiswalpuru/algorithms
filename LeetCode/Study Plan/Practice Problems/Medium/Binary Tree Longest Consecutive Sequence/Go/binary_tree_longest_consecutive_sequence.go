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

func binary_tree_longest_consecutive_sequence(root *TreeNode) int {
	res := 1
	traverse(root, 1, &res)
	return res
}

func traverse(root *TreeNode, temp int, res *int) {
	if root == nil {
		return
	}
	*res = max(*res, temp)
	if root.Left != nil {
		if root.Left.Val-root.Val == 1 {
			traverse(root.Left, temp+1, res)
		} else {
			traverse(root.Left, 1, res)
		}
	}
	if root.Right != nil {
		if root.Right.Val-root.Val == 1 {
			traverse(root.Right, temp+1, res)
		} else {
			traverse(root.Right, 1, res)
		}
	}
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	fmt.Println(binary_tree_longest_consecutive_sequence(root))
}
