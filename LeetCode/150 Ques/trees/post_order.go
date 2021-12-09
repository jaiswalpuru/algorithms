/*
Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [3,2,1]

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
Output: [2,1]
*/

package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func post_order(root *Node) {
	if root == nil {
		return
	}
	post_order(root.Left)
	post_order(root.Right)
	fmt.Println(root.Val)
}

func reverse(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func post_order_non_recursive(root *Node) []int {
	if root == nil {
		return nil
	}

	s1, s2 := []*Node{}, []int{}
	s1 = append(s1, root)

	for len(s1) > 0 {
		curr := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]

		s2 = append(s2, curr.Val)

		if curr.Left != nil {
			s1 = append(s1, curr.Left)
		}

		if curr.Right != nil {
			s1 = append(s1, curr.Right)
		}
	}
	return reverse(s2)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Right.Left = New(6)
	root.Right.Right = New(7)
	fmt.Println(post_order_non_recursive(root))
	post_order(root)
}
