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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mode(root *TreeNode) []int {
	mp := make(map[int]int)
	max_val := math.MinInt64
	min_val := 1
	_mode(root, &mp, &max_val, &min_val)
	res := []int{}
	for k, v := range mp {
		if v == max_val {
			res = append(res, k)
		}
	}
	return res
}

func _mode(root *TreeNode, mp *map[int]int, max_val *int, min_val *int) {
	if root == nil {
		return
	}
	_mode(root.Left, mp, max_val, min_val)
	(*mp)[root.Val]++
	*max_val = max(*max_val, (*mp)[root.Val])
	*min_val = min(*min_val, (*mp)[root.Val])
	_mode(root.Right, mp, max_val, min_val)
}

func main() {
	root := New(1)
	root.Right = New(2)
	root.Right.Left = New(2)
	fmt.Println(mode(root))
}
