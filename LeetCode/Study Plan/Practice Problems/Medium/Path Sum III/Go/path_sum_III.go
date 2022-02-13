package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

//---------------------using recursion---------------------------
func target_path(root *TreeNode, target int) int {
	total := 0
	_target_path(root, target, &total)
	return total
}

func _target_path(root *TreeNode, target int, total *int) {
	if root != nil {
		_target_path_dfs(root, target, total, 0)
		_target_path(root.Left, target, total)
		_target_path(root.Right, target, total)
	}
}

func _target_path_dfs(root *TreeNode, target int, total *int, sum int) {
	if root == nil {
		return
	}
	sum += root.Val
	if sum == target {
		*total += 1
	}
	_target_path_dfs(root.Left, target, total, sum)
	_target_path_dfs(root.Right, target, total, sum)
}

//---------------------using recursion---------------------------

//----------------------using prefix sum------------------------- (refer LC solution)
func target_path_prefix_sum(root *TreeNode, target int) int {
	h := make(map[int]int)
	count := 0
	preorder(root, &h, &count, 0, target)
	return count
}

func preorder(root *TreeNode, h *map[int]int, count *int, curr_sum int, target int) {
	if root == nil {
		return
	}

	curr_sum += root.Val

	if curr_sum == target {
		(*count)++
	}

	*count += (*h)[curr_sum-target]
	(*h)[curr_sum]++
	preorder(root.Left, h, count, curr_sum, target)
	preorder(root.Right, h, count, curr_sum, target)
	(*h)[curr_sum]--
}

//----------------------using prefix sum-------------------------

func main() {
	root := New(10)
	root.Left = New(5)
	root.Right = New(-3)
	root.Left.Left = New(3)
	root.Left.Right = New(2)
	root.Left.Right.Right = New(1)
	root.Left.Left.Left = New(3)
	root.Left.Left.Right = New(-2)
	root.Right.Right = New(11)
	fmt.Println(target_path_prefix_sum(root, 8))
}
