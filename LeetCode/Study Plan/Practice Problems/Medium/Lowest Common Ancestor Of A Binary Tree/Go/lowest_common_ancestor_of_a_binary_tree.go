package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func lowest_common_ancestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowest_common_ancestor(root.Left, p, q)
	right := lowest_common_ancestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}

func main() {
	root := New(3)
	root.Left = New(5)
	root.Left.Left = New(6)
	root.Left.Right = New(2)
	root.Left.Right.Left = New(7)
	root.Left.Right.Right = New(4)
	root.Right = New(1)
	root.Right.Left = New(0)
	root.Right.Right = New(8)
	fmt.Println(lowest_common_ancestor(root, root.Left, root.Right))
}
