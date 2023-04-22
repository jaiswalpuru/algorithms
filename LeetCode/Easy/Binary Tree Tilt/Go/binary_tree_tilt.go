package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func tilt(root *TreeNode) int {
	sum := 0
	traverse(root, &sum)
	return sum
}

func traverse(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	left := traverse(root.Left, sum)
	right := traverse(root.Right, sum)
	*sum += abs(left - right)
	return left + right + root.Val
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Left.Left = New(3)
	root.Left.Right = New(5)
	root.Right = New(9)
	root.Right.Right = New(7)
	fmt.Println(tilt(root))
}
