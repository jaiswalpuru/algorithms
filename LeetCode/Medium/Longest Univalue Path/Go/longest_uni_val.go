package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root *TreeNode, path *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, path)
	r := dfs(root.Right, path)

	curr := 1
	if root.Left != nil && root.Left.Val == root.Val {
		curr += l
	} else {
		l = 0
	}

	if root.Right != nil && root.Right.Val == root.Val {
		curr += r
	} else {
		r = 0
	}

	*path = max(*path, curr-1)
	//this backtracking step
	return max(l+1, r+1)

}

func longest_uni_val(root *TreeNode) int {
	path := 0
	dfs(root, &path)
	return path
}

func main() {
	root := New(5)
	root.Left = New(4)
	root.Left.Left = New(1)
	root.Left.Right = New(1)
	root.Right = New(5)
	root.Right.Right = New(5)
	fmt.Println(longest_uni_val(root))
}
