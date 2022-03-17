package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func main() {
	root := New(1)
	root.Right = New(2)
	root.Right.Left = New(3)
	fmt.Println(inorder(root))
}

func inorder(root *TreeNode) []int {
	output := []int{}
	if root == nil {
		return output
	}

	temp := root
	stack := []*TreeNode{}
	for temp != nil || len(stack) > 0 {
		for temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		}

		temp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output = append(output, temp.Val)
		temp = temp.Right

	}
	return output
}

func inorder_recursive(root *TreeNode) []int {
	output := make([]int, 0)
	_inorder_recursive(root, &output)
	return output
}

func _inorder_recursive(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}

	_inorder_recursive(root.Left, output)
	*output = append(*output, root.Val)
	_inorder_recursive(root.Right, output)
}
