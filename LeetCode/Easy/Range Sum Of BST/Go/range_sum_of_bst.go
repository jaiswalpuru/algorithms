package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func range_sum_of_bst(root *TreeNode, l, h int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val >= l && root.Val <= h {
		sum += root.Val + range_sum_of_bst(root.Left, l, h) + range_sum_of_bst(root.Right, l, h)
	} else {
		sum += range_sum_of_bst(root.Left, l, h) + range_sum_of_bst(root.Right, l, h)
	}
	return sum
}

func main() {
	root := New(10)
	root.Left = New(5)
	root.Right = New(15)
	root.Left.Left = New(3)
	root.Left.Right = New(7)
	root.Right.Right = New(18)
	fmt.Println(range_sum_of_bst(root, 7, 15))
}
