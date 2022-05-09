package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func is_leaf(root *TreeNode) bool { return root.Left == nil && root.Right == nil }

func boundary_of_binary_tree(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	if !is_leaf(root) {
		res = append(res, root.Val)
	}

	left := root.Left
	for left != nil {
		if !is_leaf(left) {
			res = append(res, left.Val)
		}
		if left.Left != nil {
			left = left.Left
		} else {
			left = left.Right
		}
	}
	add_leaves(root, &res)
	st := []int{}
	right := root.Right
	for right != nil {
		if !is_leaf(right) {
			st = append(st, right.Val)
		}

		if right.Right != nil {
			right = right.Right
		} else {
			right = right.Left
		}
	}

	for len(st) > 0 {
		res = append(res, st[len(st)-1])
		st = st[:len(st)-1]
	}
	return res
}

func add_leaves(root *TreeNode, res *[]int) {
	if is_leaf(root) {
		*res = append(*res, root.Val)
	} else {
		if root.Left != nil {
			add_leaves(root.Left, res)
		}
		if root.Right != nil {
			add_leaves(root.Right, res)
		}
	}
}

func main() {
	root := New(1)
	root.Right = New(2)
	root.Right.Right = New(4)
	root.Right.Left = New(3)
	fmt.Println(boundary_of_binary_tree(root))
}
