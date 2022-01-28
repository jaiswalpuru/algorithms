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
	if root == nil {
		return 0
	}

	max_sum := math.MinInt64
	_get_max(root, &max_sum)
	return max_sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _get_max(root *TreeNode, max_sum *int) int {
	if root == nil {
		return 0
	}
	left_gain := max(_get_max(root.Left, max_sum), 0)
	right_gain := max(_get_max(root.Right, max_sum), 0)
	path_sum := root.Val + left_gain + right_gain
	*max_sum = max(*max_sum, path_sum)
	return root.Val + max(left_gain, right_gain)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	fmt.Println(get_max_path_sum(root))
}
