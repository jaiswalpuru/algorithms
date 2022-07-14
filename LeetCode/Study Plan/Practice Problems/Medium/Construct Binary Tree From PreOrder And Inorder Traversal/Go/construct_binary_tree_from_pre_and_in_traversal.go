package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func get_index(in []int, val int) int {
	for i := 0; i < len(in); i++ {
		if in[i] == val {
			return i
		}
	}
	return -1
}

func pre_to_in(pre, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := New(pre[0])
	ind := get_index(in, pre[0])
	if ind != -1 {
		root.Left = pre_to_in(pre[1:ind+1], in[:ind])
		root.Right = pre_to_in(pre[ind+1:], in[ind+1:])
	}
	return root
}

func main() {
	fmt.Println(pre_to_in([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
