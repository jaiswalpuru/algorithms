package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right,
	Parent *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func lcs_binary_tree(p, q *TreeNode) *TreeNode {
	mp := make(map[int]struct{})
	update_root_to_path(p, &mp)
	for q != nil {
		if _, ok := mp[q.Val]; ok {
			return q
		}
		q = q.Parent
	}
	return nil
}

func update_root_to_path(p *TreeNode, mp *map[int]struct{}) {
	for p != nil {
		(*mp)[p.Val] = struct{}{}
		p = p.Parent
	}
}

func main() {
	root := New(3)
	root.Right = New(1)
	root.Left = New(5)
	root.Right.Parent = root
	root.Left.Parent = root
	root.Left.Left = New(6)
	root.Left.Left.Parent = root.Left
	root.Left.Right = New(2)
	root.Left.Right.Parent = root.Left
	root.Left.Right.Right = New(4)
	root.Left.Right.Right.Parent = root.Left.Right
	root.Left.Right.Left = New(7)
	root.Left.Right.Left.Parent = root.Left.Right
	root.Right.Right = New(8)
	root.Right.Right.Parent = root.Right
	root.Right.Left = New(0)
	root.Right.Left.Parent = root.Right
	fmt.Println(lcs_binary_tree(root.Left, root.Right))
}
