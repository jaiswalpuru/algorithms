package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

var mod = int(1e9+7)

func New(val int) *TreeNode {return &TreeNode{Val:val, Left:nil, Right:nil}}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max_product_split_binary_tree(root *TreeNode) int {
	sum_arr := make([]int, 0)
	sum := _subtree_sum(root, &sum_arr)
	res := 0
	for i:=0;i<len(sum_arr); i++{
		res = max(res, sum_arr[i]*(sum-sum_arr[i]))
	}
	return res%mod
}

func _subtree_sum(root *TreeNode, sum_arr *[]int) int {
	if root == nil {
		return 0
	}
	left_sum := _subtree_sum(root.Left, sum_arr)
	right_sum := _subtree_sum(root.Right, sum_arr)
	tot_sum := root.Val+left_sum+right_sum
	*sum_arr = append(*sum_arr, tot_sum)
	return tot_sum
}

func main () {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Right.Left = New(6)
	fmt.Println(max_product_split_binary_tree(root))
}
