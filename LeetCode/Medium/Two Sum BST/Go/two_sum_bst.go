package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func two_sum(r1, r2 *TreeNode, target int) bool {
	stack := []*TreeNode{r1}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if dfs(r2, target-curr.Val) {
			return true
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}
	return false
}

func dfs(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if target == root.Val {
		return true
	}

	l := dfs(root.Left, target)
	r := dfs(root.Right, target)
	return l || r
}

func main() {
	r1 := New(2)
	r1.Left = New(1)
	r1.Right = New(4)
	r2 := New(1)
	r2.Left = New(0)
	r2.Right = New(3)
	fmt.Println(two_sum(r1, r2, 5))
}
