/*
Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

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

func right_view(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{}

	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]

			if i == n-1 {
				ans = append(ans, curr.Val)
			}
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
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Right = New(5)
	root.Right.Right = New(4)
	fmt.Println(right_view(root))
}
