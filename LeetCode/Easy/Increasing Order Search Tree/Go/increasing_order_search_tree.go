package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func increasing_order_search_tree(root *TreeNode) *TreeNode {
	res := New(0)
	temp := res
	_new_tree(root, &temp)
	return res.Right
}

func _new_tree(root *TreeNode, res **TreeNode) {
	if root == nil {
		return
	}

	_new_tree(root.Left, res)
	root.Left = nil
	(*res).Right = root
	*res = root
	_new_tree(root.Right, res)
}

func main() {
	root := New(5)
	root.Left = New(3)
	root.Left.Left = New(2)
	root.Left.Right = New(4)
	root.Left.Left.Left = New(1)
	root.Right = New(6)
	root.Right.Right = New(8)
	root.Right.Right.Left = New(7)
	root.Right.Right.Right = New(9)
	fmt.Println(increasing_order_search_tree(root))
}
