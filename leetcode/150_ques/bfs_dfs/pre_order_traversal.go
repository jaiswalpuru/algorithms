/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,2,3]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
Input: root = [1,2]
Output: [1,2]

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

func pre_order_recursive(root *Node, res *[]int) []int {
	if root == nil {
		return *res
	}

	*res = append(*res, root.Val)
	pre_order_recursive(root.Left, res)
	pre_order_recursive(root.Right, res)
	return *res
}

func pre_order(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := make([]*Node, 0)
	stack = append(stack, root)
	res := make([]int, 0)

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, temp.Val)
		if temp.Right != nil {
			stack = append(stack, temp.Right)
		}
		if temp.Left != nil {
			stack = append(stack, temp.Left)
		}
	}

	return res
}

func main() {
	root := New(1)
	root.Left = New(4)
	root.Right = New(3)
	root.Left.Left = New(2)
	fmt.Println(pre_order_recursive(root, &[]int{}))
	fmt.Println(pre_order(root))
}
