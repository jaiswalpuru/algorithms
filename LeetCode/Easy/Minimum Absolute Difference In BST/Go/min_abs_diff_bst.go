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

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func min_abs_diff(root *TreeNode) int {
	res := []int{}
	inorder(root, &res)
	n := len(res)
	min_val := math.MaxInt64
	for i := 0; i < n-1; i++ {
		min_val = min(min_val, abs(res[i]-res[i+1]))
	}
	return min_val
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	root.Right = New(6)
	fmt.Println(min_abs_diff(root))
}
