/*
Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Invert a binary tree.

For example, given the following tree:

    a
   / \
  b   c
 / \  /
d   e f
should become:

  a
 / \
 c  b
 \  / \
  f e  d

*/

package main

import "fmt"

type Node struct {
	val   rune
	left  *Node
	right *Node
}

func New(val rune) *Node { return &Node{val: val, left: nil, right: nil} }

func invert_binary_tree(root *Node) {
	if root == nil {
		return
	}
	root.left, root.right = root.right, root.left
	if root.left != nil {
		fmt.Println(string(root.left.val))
	}
	if root.right != nil {
		fmt.Println(string(root.right.val))
	}
	invert_binary_tree(root.left)
	invert_binary_tree(root.right)
	return
}

func invert_binary_tree_recursive(root *Node) *Node {
	if root != nil {
		root.left, root.right = invert_binary_tree_recursive(root.right), invert_binary_tree_recursive(root.left)
	}
	return root
}

func main() {
	root := New('a')
	root.left = New('b')
	root.right = New('c')
	root.left.left = New('d')
	root.left.right = New('e')
	root.right.left = New('f')
	root_one := invert_binary_tree_recursive(root)
	fmt.Println(string(root_one.left.val))
	fmt.Println()
	invert_binary_tree(root)
}
