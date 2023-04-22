package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func find_max(arr []int) int {
	max_val := math.MinInt64
	ind := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > max_val {
			max_val = arr[i]
			ind = i
		}
	}
	return ind
}

func generate_binary_tree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	ind := find_max(arr)
	root := New(arr[ind])
	root.Left = generate_binary_tree(arr[:ind])
	root.Right = generate_binary_tree(arr[ind+1:])
	return root
}

func main() {
	fmt.Println(generate_binary_tree([]int{3, 2, 1, 6, 0, 5}))
}
