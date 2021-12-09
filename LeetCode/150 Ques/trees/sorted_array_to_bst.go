/*
Given an integer array nums where the elements are sorted in ascending order,
convert it to a height-balanced binary search tree.

A height-balanced binary tree is a binary tree in which the depth of the two
subtrees of every node never differs by more than one.

Example 1:
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,3] and [3,1] are both a height-balanced BSTs.
*/

package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func sorted_arr_to_bst(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	if len(arr) == 1 {
		return &Node{Val: arr[0]}
	}

	var root *Node
	n := len(arr) / 2
	root = New(arr[n])
	root.Left = sorted_arr_to_bst(arr[:n])
	root.Right = sorted_arr_to_bst(arr[n+1:])
	return root
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	inorder(root.Left)
	inorder(root.Right)
}

func main() {
	root := sorted_arr_to_bst([]int{-10, -3, 0, 5, 9})
	inorder(root)
	fmt.Println()
}
