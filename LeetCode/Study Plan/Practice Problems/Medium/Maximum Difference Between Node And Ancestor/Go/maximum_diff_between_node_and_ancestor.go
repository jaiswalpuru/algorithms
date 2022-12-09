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

func New(val int) *TreeNode { return &TreeNode{Val: val} }

//-----------------------brute force-------------------------------
func maximum_diff_between_node_and_ancestor(root *TreeNode) int {
	res := int(math.MinInt64)
	recur(root, &res)
	return res
}

func recur(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	l_max := max_val(root.Left)
	r_max := max_val(root.Right)
	l_min := min_val(root.Left)
	r_min := min_val(root.Right)
	if l_max != -1 {
		*res = max(*res, abs(l_max-root.Val))
	}
	if r_max != -1 {
		*res = max(*res, abs(r_max-root.Val))
	}
	if l_min != int(math.MaxInt64) {
		*res = max(*res, abs(l_min-root.Val))
	}
	if r_min != int(math.MaxInt64) {
		*res = max(*res, abs(r_min-root.Val))
	}
	recur(root.Left, res)
	recur(root.Right, res)
}

func max_val(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return max(root.Val, max(max_val(root.Left), max_val(root.Right)))
}

func min_val(root *TreeNode) int {
	if root == nil {
		return int(math.MaxInt64)
	}
	return min(root.Val, min(min_val(root.Left), min_val(root.Right)))
}

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

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

//-----------------------brute force-------------------------------

//--------------------a better approach----------------------------
func maximum_diff_between_node_and_ancestor_eff(root *TreeNode) int {
	return recur_eff(root, root.Val, root.Val)
}

func recur_eff(root *TreeNode, curr_max int, curr_min int) int {
	if root == nil {
		return curr_max - curr_min
	}
	curr_max = max(curr_max, root.Val)
	curr_min = min(curr_min, root.Val)
	left := recur_eff(root.Left, curr_max, curr_min)
	right := recur_eff(root.Right, curr_max, curr_min)
	return max(left, right)
}

//--------------------a better approach----------------------------

func main() {
	root := New(8)
	root.Left = New(3)
	root.Right = New(10)
	root.Left.Left = New(1)
	root.Left.Right = New(6)
	root.Left.Right.Right = New(7)
	root.Left.Right.Left = New(4)
	root.Right.Right = New(14)
	root.Right.Right.Left = New(13)
	fmt.Println(maximum_diff_between_node_and_ancestor_eff(root))
}
