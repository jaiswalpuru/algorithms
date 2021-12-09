/*
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf
path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: false

Example 3:
Input: root = [1,2], targetSum = 0
Output: false
*/

package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func has_path_sum(root *Node, target int) bool {
	return _has_path_sum(root, target)
}

func _has_path_sum(root *Node, target int) bool {
	if root == nil {
		return false
	}
	target -= root.Val

	if root.Left == nil && root.Right == nil {
		if target == 0 {
			return true
		}
	}
	return _has_path_sum(root.Left, target) || _has_path_sum(root.Right, target)
}

func main() {
	root := New(5)
	root.Left = New(4)
	root.Left.Left = New(11)
	root.Left.Left.Left = New(7)
	root.Left.Left.Right = New(2)
	root.Right = New(8)
	root.Right.Left = New(13)
	root.Right.Right = New(4)
	root.Right.Right.Right = New(1)
	fmt.Println(has_path_sum(root, 22))
}
