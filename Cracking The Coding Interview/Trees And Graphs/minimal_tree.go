package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func minimal_tree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	l, r := 0, len(arr)
	mid := (l + r) / 2
	root := New(arr[mid])
	root.Left = minimal_tree(arr[:mid])
	root.Right = minimal_tree(arr[mid+1:])
	return root
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " -> ")
	inorder(root.Right)
}

func main() {
	root := minimal_tree([]int{1, 2, 3, 4, 5})
	inorder(root)
	fmt.Println()
}
