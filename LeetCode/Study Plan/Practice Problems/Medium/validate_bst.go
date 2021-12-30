/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func _validate_bst(node, left, right *Node) bool {
	if node == nil {
		return true
	}

	if left != nil && left.Val >= node.Val {
		return false
	}

	if right != nil && right.Val <= node.Val {
		return false
	}

	return _validate_bst(node.Left, left, node) && _validate_bst(node.Right, node, right)
}

func validate_bst(node *Node) bool {
	return _validate_bst(node, nil, nil)
}

func main() {
	node := New(5)
	node.Left = New(4)
	node.Right = New(6)
	node.Right.Left = New(3)
	node.Right.Right = New(7)
	fmt.Println(validate_bst(node))
}
