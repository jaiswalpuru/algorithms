/*
Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values.
(i.e., from left to right, level by level from leaf to root).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
*/

package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func reverse(arr *[][]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}

func level_order_traversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := make([]*TreeNode, 0)
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
	reverse(&res)
	return res
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(level_order_traversal(root))
}
