package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func longest_consecutive_subsequence(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	traverse(root, &res)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func traverse(root *TreeNode, res *int) []int {
	if root == nil {
		return []int{0, 0}
	}

	inc, dec := 1, 1
	if root.Left != nil {
		left := traverse(root.Left, res)
		if root.Val == root.Left.Val+1 {
			dec = left[1] + 1
		} else if root.Val == root.Left.Val-1 {
			inc = left[0] + 1
		}
	}

	if root.Right != nil {
		right := traverse(root.Right, res)
		if root.Val == root.Right.Val+1 {
			dec = max(dec, right[1]+1)
		} else if root.Val == root.Right.Val-1 {
			inc = max(inc, right[0]+1)
		}
	}

	*res = max(*res, inc+dec-1)
	return []int{inc, dec}
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	fmt.Println(longest_consecutive_subsequence(root))
}
