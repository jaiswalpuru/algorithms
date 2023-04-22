package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func trim_binary_tree(root *TreeNode, low, high int) *TreeNode {
	prev := New(-1)
	prev.Left = root
	inorder(prev, low, high)
	return prev.Left
}

func inorder(root *TreeNode, l, h int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		if root.Left.Val < l {
			root.Left = root.Left.Right
			inorder(root, l, h)
		} else if root.Left.Val > h {
			root.Left = root.Left.Left
			inorder(root, l, h)
		} else {
			inorder(root.Left, l, h)
		}
	}

	if root.Right != nil {
		if root.Right.Val < l {
			root.Right = root.Right.Right
			inorder(root, l, h)
		} else if root.Right.Val > h {
			root.Right = root.Right.Left
			inorder(root, l, h)
		} else {
			inorder(root.Right, l, h)
		}
	}
}


func main() {
	root := New(3)
	root.Left = New(1)
	root.Right = New(4)
	root.Left.Right = New(2)
	fmt.Println(trim_binary_tree(root, 1, 2))
}
