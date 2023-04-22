package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func convert_sorted_array_to_bst(arr []int) *TreeNode { return recur(arr, 0, len(arr)-1) }

func recur(arr []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := (l + r) / 2
	root := New(arr[mid])
	root.Left = recur(arr, l, mid-1)
	root.Right = recur(arr, mid+1, r)
	return root
}

func main() {
	fmt.Println(convert_sorted_array_to_bst([]int{-10, -3, 0, 5, 9}))
}
