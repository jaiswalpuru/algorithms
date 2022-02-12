package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func sum_to_root(root *TreeNode) int {
	sum := 0
	_sum_to_root(root, 0, &sum)
	return sum
}

func _sum_to_root(root *TreeNode, num int, sum *int) {
	if root == nil {
		return
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += num
	}

	_sum_to_root(root.Left, num, sum)
	_sum_to_root(root.Right, num, sum)
}

func main() {
	root := New(1)
	root.Left = New(0)
	// root.Right = New(0)
	// root.Left.Left = New(5)
	// root.Left.Right = New(1)
	fmt.Println(sum_to_root(root))
}
