package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func inorder(root *TreeNode, mp *map[int]struct{}, k int, flag *bool) {
	if root == nil {
		return
	}
	inorder(root.Left, mp, k, flag)
	if _, ok := (*mp)[k-root.Val]; ok {
		*flag = true
		return
	} else {
		(*mp)[root.Val] = struct{}{}
	}
	inorder(root.Right, mp, k, flag)
}

func two_sum(root *TreeNode, k int) bool {
	mp := make(map[int]struct{})
	flag := false
	inorder(root, &mp, k, &flag)
	return flag
}

func main() {
	root := New(5)
	root.Left = New(3)
	root.Left.Left = New(2)
	root.Left.Right = New(4)
	root.Right = New(6)
	root.Right.Right = New(7)
	fmt.Println(two_sum(root, 9))
}
