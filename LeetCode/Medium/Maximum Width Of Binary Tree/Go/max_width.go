package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type Pair struct {
	node *TreeNode
	ind  int
}

func max_width(root *TreeNode) int {
	q := []Pair{{root, 0}}
	width := 1
	for len(q) > 0 {
		m := len(q)
		for i := 0; i < m; i++ {
			curr := q[i].node
			ind := q[i].ind
			if curr.Left != nil {
				q = append(q, Pair{curr.Left, 2 * ind})
			}
			if curr.Right != nil {
				q = append(q, Pair{curr.Right, 2*ind + 1})
			}
		}
		q = q[m:]
		if len(q) > 0 {
			width = max(width, q[len(q)-1].ind-q[0].ind+1)
		}
	}
	return width
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := New(1)
	root.Left = New(3)
	root.Right = New(2)
	root.Left.Left = New(5)
	root.Left.Left.Left = New(6)
	root.Right.Right = New(9)
	root.Right.Right.Right = New(7)
	fmt.Println(max_width(root))
}
