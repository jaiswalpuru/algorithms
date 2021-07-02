/*
You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node.
If such a node does not exist, return null.

Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]
*/

package main

import "fmt"

type BST struct {
	data        int
	left, right *BST
}

func New(data int) *BST { return &BST{data: data, left: nil, right: nil} }

func search_in_bst(root *BST, val int) *BST {
	if root == nil {
		return nil
	}
	if root.data == val {
		return root
	}
	left := search_in_bst(root.left, val)
	right := search_in_bst(root.right, val)
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func pre_order(root *BST) {
	if root == nil {
		return
	}
	fmt.Print(root.data, "->")
	pre_order(root.left)
	pre_order(root.right)
}

func main() {
	root := New(4)
	root.left = New(2)
	root.right = New(7)
	root.left.left = New(1)
	root.left.right = New(3)
	r := search_in_bst(root, 2)
	pre_order(r)
}
