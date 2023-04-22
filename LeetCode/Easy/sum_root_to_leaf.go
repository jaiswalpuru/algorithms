/*
You are given the root of a binary tree where each node has a value 0 or 1. Each root-to-leaf path represents a
binary number starting with the most significant bit.

For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.
For all leaves in the tree, consider the numbers represented by the path from the root to that leaf. Return the sum of
these numbers.

The test cases are generated so that the answer fits in a 32-bits integer.
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func sum_root_to_leaf(root *TreeNode) int {
	sum := 0
	_sum_root_to_leaf(root, 0, &sum)
	return sum
}

func _sum_root_to_leaf(root *TreeNode, sum int, ans *int) {
	if root == nil {
		return
	}
	sum = sum<<1 | root.Val
	if root.Left == nil && root.Right == nil {
		*ans += sum
	}
	_sum_root_to_leaf(root.Left, sum, ans)
	_sum_root_to_leaf(root.Right, sum, ans)
}

func main() {
	root := New(1)
	root.Left = New(0)
	root.Left.Left = New(0)
	root.Left.Right = New(1)
	root.Right = New(1)
	root.Right.Left = New(0)
	root.Right.Right = New(1)
	fmt.Println(sum_root_to_leaf(root))
}
