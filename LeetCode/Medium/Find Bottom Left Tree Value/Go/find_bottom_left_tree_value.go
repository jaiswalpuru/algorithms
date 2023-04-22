package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val:val, Left:nil, Right :nil}}

func find_bottom_left_tree_val(root *TreeNode) int {
	q := []*TreeNode{root}
	left_bottom_val := root.Val
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		left_bottom_val = curr.Val
		n := len(q)

		if curr.Left != nil {
			q = append(q, curr.Left)
		}
		if curr.Right != nil {
			q = append(q, curr.Right)
		}

		for i:=0;i<n;i++{
			curr = q[0]
			q= q[1:]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
	}
	return left_bottom_val
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Left = New(4)
	root.Right = New(3)
	root.Right.Left = New(5)
	root.Right.Left.Left = New(7)
	root.Right.Right = New(6)
	fmt.Println(find_bottom_left_tree_val(root))
}