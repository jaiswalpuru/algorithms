/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
*/

package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data, left: nil, right: nil} }

func _is_valid(root, left, right *Node) bool {
	if root == nil {
		return true
	}

	if left != nil && root.data <= left.data {
		return false
	}

	if right != nil && root.data >= right.data {
		return false
	}

	return _is_valid(root.left, left, root) && _is_valid(root.right, root, right)
}

func is_valid(root *Node) bool {
	return _is_valid(root, nil, nil)
}

func main() {
	root := New(5)
	root.left = New(1)
	root.right = New(4)
	root.right.left = New(3)
	root.right.right = New(6)
	fmt.Println(is_valid(root))
	root_one := New(0)
	root_one.left = New(-1)
	fmt.Println(is_valid(root_one))
	root_two := New(5)
	root_two.left = New(4)
	root_two.right = New(6)
	root_two.right.left = New(3)
	root_two.right.right = New(7)
	fmt.Println(is_valid(root_two))
	root_three := New(1)
	root_three.right = New(1)
	fmt.Println(is_valid(root_three))
}
