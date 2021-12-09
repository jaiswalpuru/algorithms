/*
Given the root of a binary tree, return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
*/

package main

import "fmt"

type Node struct {
	Left, Right *Node
	Val         int
}

func New(data int) *Node { return &Node{Val: data, Left: nil, Right: nil} }

func level_order_traverse(root *Node) [][]int {

	if root == nil {
		return nil
	}

	q := make([]*Node, 0)
	q = append(q, root)
	res := make([][]int, 0)
	for len(q) > 0 {
		n := len(q)
		temp := []int{}
		for i := 0; i < n; i++ {
			temp = append(temp, q[0].Val)
			curr := q[0]
			q = q[1:]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(level_order_traverse(root))
	fmt.Println(level_order_traverse(New(1)))
	fmt.Println(level_order_traverse(nil))
}
