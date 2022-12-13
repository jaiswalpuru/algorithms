package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

//---------------------brute force-----------------------
func paths_with_sum(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	path := path_with_sum_from_node(root, target, 0)
	left := paths_with_sum(root.Left, target)
	right := paths_with_sum(root.Right, target)
	return path + left + right
}

func path_with_sum_from_node(root *TreeNode, target int, curr_sum int) int {
	if root == nil {
		return 0
	}
	curr_sum += root.Val
	total_paths := 0
	if curr_sum == target {
		total_paths++
	}
	total_paths += path_with_sum_from_node(root.Left, target, curr_sum)
	total_paths += path_with_sum_from_node(root.Right, target, curr_sum)
	return total_paths
}

//---------------------brute force-----------------------

//---------------------efficient-----------------------
func path_with_sum_eff(root *TreeNode, target int) int {
	return path_with_sum_eff_recur(root, target, 0, map[int]int{})
}

func path_with_sum_eff_recur(root *TreeNode, target, run_sum int, mp map[int]int) int {
	if root == nil {
		return 0
	}
	run_sum += root.Val
	remaining := run_sum - target
	total_paths := mp[remaining]
	if run_sum == target {
		total_paths++
	}
	mp[run_sum]++
	total_paths += path_with_sum_eff_recur(root.Left, target, run_sum, mp)
	total_paths += path_with_sum_eff_recur(root.Right, target, run_sum, mp)
	return total_paths
}

//---------------------efficient-----------------------

func main() {
	root := New(10)
	root.Left = New(5)
	root.Left.Left = New(3)
	root.Left.Right = New(2)
	root.Left.Right.Right = New(1)
	root.Left.Left.Right = New(-2)
	root.Left.Left.Left = New(3)
	root.Right = New(-3)
	root.Right.Right = New(11)
	fmt.Println(path_with_sum_eff(root, 8))
}
