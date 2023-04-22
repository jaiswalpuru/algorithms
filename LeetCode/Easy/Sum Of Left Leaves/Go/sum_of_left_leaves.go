package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func sum_of_left_leaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.Left != nil {
			q = append(q, curr.Left)
			if curr.Left.Left == nil && curr.Left.Right == nil {
				sum += curr.Left.Val
			}
		}
		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}
	return sum
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(sum_of_left_leaves(root))
}
