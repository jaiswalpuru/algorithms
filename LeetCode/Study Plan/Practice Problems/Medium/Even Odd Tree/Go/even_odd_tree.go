package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func even_odd_tree(root *TreeNode) bool {
	i := 0

	type Pair struct {
		node *TreeNode
		lvl  int
	}

	q := []Pair{{root, i}}
	for len(q) > 0 {
		n := len(q)
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				curr := q[j]
				if (curr.node.Val%2 != 1) || (j > 0 && q[j].node.Val <= q[j-1].node.Val) {
					return false
				}
				if curr.node.Left != nil {
					q = append(q, Pair{node: curr.node.Left, lvl: curr.lvl + 1})
				}
				if curr.node.Right != nil {
					q = append(q, Pair{node: curr.node.Right, lvl: curr.lvl + 1})
				}
			}
		} else {
			for j := 0; j < n; j++ {
				curr := q[j]
				if (curr.node.Val%2 != 0) || (j > 0 && q[j].node.Val >= q[j-1].node.Val) {
					return false
				}
				if curr.node.Left != nil {
					q = append(q, Pair{node: curr.node.Left, lvl: curr.lvl + 1})
				}
				if curr.node.Right != nil {
					q = append(q, Pair{node: curr.node.Right, lvl: curr.lvl + 1})
				}
			}
		}
		i++
		q = q[n:]
	}
	return true
}

func main() {
	root := New(5)
	root.Left = New(4)
	root.Right = New(2)
	root.Right.Left = New(7)
	root.Left.Left = New(3)
	root.Left.Right = New(3)
	// root.Left.Left.Left = New(12)
	// root.Left.Left.Right = New(8)
	// root.Right.Right = New(9)
	// root.Right.Left = New(7)
	// root.Right.Right.Right = New(2)
	// root.Right.Left.Left = New(6)
	fmt.Println(even_odd_tree(root))
}
