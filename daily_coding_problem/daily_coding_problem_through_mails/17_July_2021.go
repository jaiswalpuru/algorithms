/*
Given a binary tree of integers, find the maximum path sum between two nodes.
The path must go through at least one node, and does not need to go through the root.
*/

package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func New(data int) *Node { return &Node{Val: data, Left: nil, Right: nil} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func update_val(val int, res *int) {
	if val > *res {
		*res = val
	}
}

func _max_path_sum(root *Node, res *int) int {
	update_val(root.Val, res)
	if root == nil {
		return root.Val
	}

	var left, right int
	if root.Left != nil {
		left = _max_path_sum(root.Left, res)
		update_val(left+root.Val, res)
	}

	if root.Right != nil {
		right = _max_path_sum(root.Right, res)
		update_val(right+root.Val, res)
	}

	if root.Left != nil && root.Right != nil {
		update_val(left+right+root.Val, res)
	}

	return max(max(left, right)+root.Val, root.Val)
}

func max_path_sum(root *Node) int {
	res := 0
	_max_path_sum(root, &res)
	return res
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	fmt.Println(max_path_sum(root))
}
