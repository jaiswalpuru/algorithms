/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]
    1
     \
      2
     /
    3

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
Input: root = [1,2]
Output: [2,1]

Example 5:
Input: root = [1,null,2]
Output: [1,2]
*/
package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func in_order_recursive(root *Node) {
	if root == nil {
		return
	}
	in_order_recursive(root.Left)
	fmt.Println(root.Val)
	in_order_recursive(root.Right)
}

func in_order(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := make([]*Node, 0)
	res := make([]int, 0)
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}

func main() {
	root := New(1)
	root.Left = New(4)
	root.Right = New(3)
	root.Left.Left = New(2)
	in_order_recursive(root)
	fmt.Println(in_order(root))
}
