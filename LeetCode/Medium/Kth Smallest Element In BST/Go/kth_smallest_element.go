package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func kth_smallest(root *TreeNode, k int) int {
	in := []int{-1}
	inorder(root, &in)
	return in[k]
}

func inorder(root *TreeNode, in *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, in)
	*in = append(*in, root.Val)
	inorder(root.Right, in)
}

func main() {
	root := New(5)
	root.Left = New(3)
	root.Left.Right = New(4)
	root.Left.Left = New(2)
	root.Left.Left.Left = New(1)
	root.Right = New(5)
	fmt.Println(kth_smallest(root, 3))
}
