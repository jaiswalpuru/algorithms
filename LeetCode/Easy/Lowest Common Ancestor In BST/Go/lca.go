package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if (p.Val < root.Val && q.Val > root.Val) || (p.Val > root.Val && root.Val > q.Val) {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lca(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lca(root.Right, p, q)
	}
	return root
}

//For recursive approach just traverse the tree until the condition is met

func main() {
	root := New(6)
	root.Left = New(2)
	root.Right = New(8)
	root.Left.Left = New(0)
	root.Left.Right = New(4)
	root.Left.Right.Left = New(3)
	root.Left.Right.Right = New(5)
	root.Right.Left = New(7)
	root.Right.Right = New(9)
	fmt.Println(lca(root, root.Left, root.Right))
}
