/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
(i.e., from left to right, then right to left for the next level and alternate between).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]

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

func New(data int) *TreeNode { return &TreeNode{Val: data, Left: nil, Right: nil} }

func zig_zag_traversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := []*TreeNode{}
	q = append(q, root)
	even := true
	ans := [][]int{}
	for len(q) > 0 {
		temp := []int{}
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]
			temp = append(temp, curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

		if !even {
			reverse(&temp)
		}
		even = !even
		ans = append(ans, temp)
	}
	return ans
}

func reverse(t *[]int) {
	n := len(*t)
	for i := 0; i < n/2; i++ {
		(*t)[i], (*t)[n-1-i] = (*t)[n-i-1], (*t)[i]
	}
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(zig_zag_traversal(root))
	root_one := New(1)
	root_one.Left = New(2)
	root_one.Right = New(3)
	root_one.Left.Left = New(4)
	root_one.Right.Right = New(5)
	fmt.Println(zig_zag_traversal(root_one))
}
