package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func tree_sum(root *TreeNode, s *[]int) int {
	if root == nil {
		return 0
	}
	left_sum := tree_sum(root.Left, s)
	right_sum := tree_sum(root.Right, s)
	tot_sum := root.Val + left_sum + right_sum
	*s = append(*s, tot_sum)
	return tot_sum
}

func equal_tree_partition(root *TreeNode) bool {
	s := []int{}
	tot_sum := tree_sum(root, &s)
	s= s[:len(s)-1]
	if tot_sum %2==0 {
		for i:=0;i<len(s);i++{
			if s[i] == tot_sum/2 {
				return true
			}
		}
	}
	return false
}

func main() {
	root := New(5)
	root.Left = New(10)
	root.Right = New(10)
	root.Right.Right = New(3)
	root.Right.Left = New(2)
	fmt.Println(equal_tree_partition(root))
}
