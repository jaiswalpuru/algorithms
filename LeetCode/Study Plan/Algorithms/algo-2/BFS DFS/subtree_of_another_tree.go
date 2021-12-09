/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure
and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
The tree tree could also be considered as a subtree of itself.
*/

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func _is_sub_tree(root *Node, subroot *Node) bool {
	if root == nil && subroot == nil {
		return true
	}
	if root == nil && subroot != nil || root != nil && subroot == nil {
		return false
	}
	if root.Val != subroot.Val {
		return false
	}
	return _is_sub_tree(root.Left, subroot.Left) && _is_sub_tree(root.Right, subroot.Right)

}

func is_subtree(root *Node, sub_root *Node) bool {
	if root == nil {
		return false
	}

	if root.Val == sub_root.Val {
		if _is_sub_tree(root, sub_root) {
			return true
		}
	}
	return is_subtree(root.Left, sub_root) && is_subtree(root.Right, sub_root)
}

func main() {
	root := &Node{Val: 3, Left: nil, Right: nil}
	root.Left = &Node{Val: 4, Left: nil, Right: nil}
	root.Right = &Node{Val: 5, Left: nil, Right: nil}
	root.Left.Left = &Node{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &Node{Val: 2, Left: nil, Right: nil}
	sub_tree := &Node{Val: 4, Left: nil, Right: nil}
	sub_tree.Left = &Node{Val: 1, Left: nil, Right: nil}
	sub_tree.Right = &Node{Val: 2, Left: nil, Right: nil}
	fmt.Println(is_subtree(root, sub_tree))
}
