/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]
*/

package main

import "fmt"

type Node struct {
	Left, Right *Node
	Val         int
}

func New(data int) *Node { return &Node{Val: data, Left: nil, Right: nil} }

func construct_tree(preorder, inorder []int) *Node {
	if len(preorder) == 0 {
		return nil
	}
	root := &Node{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			root.Left = construct_tree(preorder[1:i+1], inorder[:i])
			root.Right = construct_tree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return root
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

func main() {
	root := construct_tree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	inorder(root)
	fmt.Println()
}
