package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func construct_string_from_binary_tree(root *TreeNode) string {
	return pre_order(root)
}

func pre_order(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}
	if root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + pre_order(root.Left) + ")"
	}
	return strconv.Itoa(root.Val) + "(" + pre_order(root.Left) + ")" + "(" + pre_order(root.Right) + ")"
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	fmt.Println(construct_string_from_binary_tree(root))
}
