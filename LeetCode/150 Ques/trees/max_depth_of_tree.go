/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2
Example 3:

Input: root = []
Output: 0
Example 4:

Input: root = [0]
Output: 1
*/

package main

import "fmt"

type Node struct {
	Left, Right *Node
	Val         int
}

func New(data int) *Node { return &Node{Val: data, Left: nil, Right: nil} }

func max_depth(root *Node) int {
	if root == nil {
		return 0
	}

	q := make([]*Node, 0)
	q = append(q, root)
	ans := 0
	for len(q) > 0 {
		ans++
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
	}
	return ans
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(max_depth(root))
	root_two := New(1)
	root_two.Right = New(2)
	fmt.Println(max_depth(root_two))
}
