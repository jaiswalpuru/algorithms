package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func binary_tree_right_side_view(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	level := []*TreeNode{root}
	right_side := []int{root.Val}
	for len(level) > 0 {
		curr_level := []*TreeNode{}
		n := len(level)
		for i := 0; i < n; i++ {
			curr := level[i]
			if curr.Left != nil {
				curr_level = append(curr_level, curr.Left)
			}
			if curr.Right != nil {
				curr_level = append(curr_level, curr.Right)
			}
		}
		level = curr_level
		if len(curr_level) > 0 {
			right_side = append(right_side, curr_level[len(curr_level)-1].Val)
		}
	}
	return right_side
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Right = New(5)
	root.Right.Right = New(4)
	fmt.Println(binary_tree_right_side_view(root))
}
