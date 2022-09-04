package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func get_max_path_sum(root *TreeNode) int {
	ans := math.MinInt64
	recur(root, &ans)
	return ans
}

func recur(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left, ans)
	right := recur(root.Right, ans)
	l_sum, r_sum := max(root.Val, root.Val+left), max(root.Val, root.Val+right)
	total_sum := l_sum + r_sum - root.Val
	*ans = max(*ans, total_sum)
	return max(l_sum, r_sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	fmt.Println(get_max_path_sum(root))
}
