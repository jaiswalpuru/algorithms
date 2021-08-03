/*
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the
sequence has an edge connecting them. A node can only appear in the sequence at most once.
Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any path.

Example 1:
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

Example 2:
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
*/

package main

import (
	"fmt"
	"math"
)

type Node struct {
	Left, Right *Node
	Val         int
}

func New(data int) *Node { return &Node{Val: data, Left: nil, Right: nil} }

func max_path_sum(root *Node) int {
	res := int(math.MinInt32)
	_max_path_sum(root, &res)
	return res
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
		update_val(root.Val+left, res)
	}

	if root.Right != nil {
		right = _max_path_sum(root.Right, res)
		update_val(root.Val+right, res)
	}

	if root.Left != nil && root.Right != nil {
		update_val(root.Val+left+right, res)
	}

	return max(max(left, right)+root.Val, root.Val)
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
	fmt.Println(max_path_sum(root))
	root_one := New(-10)
	root_one.Left = New(9)
	root_one.Right = New(20)
	root_one.Right.Right = New(7)
	root_one.Right.Left = New(15)
	fmt.Println(max_path_sum(root_one))
	fmt.Println(max_path_sum(New(-3)))
}
