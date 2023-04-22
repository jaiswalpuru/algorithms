package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func distribute_coins_in_binary_tree(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, ans)
	r := dfs(root.Right, ans)
	*ans += abs(l) + abs(r)
	return root.Val + l + r - 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	root := New(1)
	root.Left = New(4)
	root.Right = New(0)
	root.Left.Left = New(0)
	root.Left.Right = New(0)
	fmt.Println(distribute_coins_in_binary_tree(root))
}
