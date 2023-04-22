package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func preorder_recursive(root *TreeNode) []int {
	output := make([]int, 0)
	_preorder_recursive(root, &output)
	return output
}

func _preorder_recursive(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	*output = append(*output, root.Val)
	_preorder_recursive(root.Left, output)
	_preorder_recursive(root.Right, output)
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	output := []int{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		output = append(output, curr.Val)
		stack = stack[:len(stack)-1]
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return output
}

func main() {
	root := New(1)
	root.Right = New(2)
	root.Right.Left = New(3)
	fmt.Println(preorder(root))
}
