package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largest_val_in_each_row(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		n := len(q)
		max_val := math.MinInt64
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
			max_val = max(max_val, curr.Val)
		}
		res = append(res, max_val)
	}

	return res
}

func main() {
	root := New(1)
	root.Left = New(3)
	root.Right = New(2)
	root.Left.Left = New(5)
	root.Left.Right = New(3)
	root.Right.Right = New(9)
	fmt.Println(largest_val_in_each_row(root))
}
