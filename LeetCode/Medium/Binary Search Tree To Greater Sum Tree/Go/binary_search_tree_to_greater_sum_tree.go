package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func inorder(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, order)
	*order = append(*order, root.Val)
	inorder(root.Right, order)
}

func inorder_mod(root *TreeNode, order []int, ind *int) {
	if root == nil {
		return
	}
	inorder_mod(root.Left, order, ind)
	root.Val = order[*ind]
	(*ind)++
	inorder_mod(root.Right, order, ind)
}

func binary_search_tree_to_greater_sum_tree(root *TreeNode) *TreeNode {
	order := []int{}
	inorder(root, &order)
	n := len(order)
	for i := n - 2; i >= 0; i-- {
		order[i] = order[i] + order[i+1]
	}
	ind := 0
	inorder_mod(root, order, &ind)
	return root
}

func main() {
	root := New(4)
	root.Left = New(1)
	root.Left.Left = New(0)
	root.Left.Right = New(2)
	root.Left.Right.Right = New(3)
	root.Right = New(6)
	root.Right.Right = New(7)
	root.Right.Right.Right = New(8)
	root.Right.Left = New(5)
	fmt.Println(binary_search_tree_to_greater_sum_tree(root))
}
