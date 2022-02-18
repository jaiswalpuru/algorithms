package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func find_leaves(root *TreeNode) [][]int {
	res := [][]int{}

	ht := make(map[int][]int)
	min_val, max_val := math.MaxInt64, math.MinInt64
	get_ht(root, &ht, &min_val, &max_val)
	for i := min_val; i <= max_val; i++ {
		res = append(res, ht[i])
	}
	return res
}

func get_ht(root *TreeNode, res *map[int][]int, min_val, max_val *int) int {
	if root == nil {
		return -1
	}

	left_ht := get_ht(root.Left, res, min_val, max_val)
	right_ht := get_ht(root.Right, res, min_val, max_val)

	curr_ht := 1 + max(left_ht, right_ht)
	*min_val = min(*min_val, curr_ht)
	*max_val = max(*max_val, curr_ht)
	(*res)[curr_ht] = append((*res)[curr_ht], root.Val)
	return curr_ht
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	fmt.Println(find_leaves(root))
}
