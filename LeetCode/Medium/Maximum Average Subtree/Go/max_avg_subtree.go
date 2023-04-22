package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func max_avg_subtree(root *TreeNode) float64 {
	q := []*TreeNode{root}
	max_avg := float64(0)
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		n, sum := 0, 0
		inorder(curr, &sum, &n)
		max_avg = max(max_avg, float64(sum)/float64(n))
		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}
	return max_avg
}

func inorder(root *TreeNode, sum *int, n *int) {
	if root == nil {
		return
	}
	*n += 1
	*sum += root.Val
	inorder(root.Left, sum, n)
	inorder(root.Right, sum, n)
}

func main() {
	root := New(5)
	root.Left = New(6)
	root.Right = New(1)
	fmt.Println(max_avg_subtree(root))
}
