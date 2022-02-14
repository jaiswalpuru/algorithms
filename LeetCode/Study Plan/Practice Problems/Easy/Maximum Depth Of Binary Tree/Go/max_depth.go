package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func max_depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val := 1
	type Pair struct {
		node *TreeNode
		lvl  int
	}

	q := []Pair{{root, 1}}
	for len(q) > 0 {
		curr := q[0].node
		val = q[0].lvl
		q = q[1:]
		if curr.Left != nil {
			q = append(q, Pair{curr.Left, val + 1})
		}

		if curr.Right != nil {
			q = append(q, Pair{curr.Right, val + 1})
		}
	}
	return val
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Right = New(7)
	root.Right.Left = New(15)
	fmt.Println(max_depth(root))
}
