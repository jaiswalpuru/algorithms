/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:
Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
Input: root = [1,2,2,null,3,null,3]
Output: false
*/

package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func _is_symmetric(left, right *Node) bool {
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	if left == nil && right == nil {
		return true
	}

	if left.Val != right.Val {
		return false
	}

	return _is_symmetric(left.Left, right.Right) && _is_symmetric(left.Right, right.Left)
}

func is_symmetric(root *Node) bool {
	if root == nil {
		return false
	}
	return _is_symmetric(root.Left, root.Right)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(2)
	root.Left.Left = New(3)
	root.Left.Right = New(4)
	root.Right.Left = New(4)
	root.Right.Right = New(3)
	fmt.Println(is_symmetric(root))
	root_one := New(1)
	root_one.Left = New(2)
	root_one.Right = New(2)
	root_one.Left.Right = New(3)
	root_one.Right.Right = New(3)
	fmt.Println(is_symmetric(root_one))

}
