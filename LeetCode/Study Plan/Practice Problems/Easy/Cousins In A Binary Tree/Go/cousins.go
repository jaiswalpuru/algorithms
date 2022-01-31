package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type Pair struct {
	node   *TreeNode
	depth  int
	parent int
}

func cousins(root *TreeNode, x, y int) bool {
	if root == nil {
		return false
	}
	q := []Pair{{node: root, depth: 0, parent: -1}}
	m, n := Pair{}, Pair{}
	for len(q) > 0 {
		curr := q[0].node
		depth := q[0].depth
		parent := q[0].parent
		q = q[1:]
		if curr.Val == x {
			m.parent = parent
			m.depth = depth
			m.node = curr
		}

		if curr.Val == y {
			n.parent = parent
			n.depth = depth
			n.node = curr
		}

		if curr.Left != nil {
			q = append(q, Pair{node: curr.Left, depth: depth + 1, parent: curr.Val})
		}

		if curr.Right != nil {
			q = append(q, Pair{node: curr.Right, depth: depth + 1, parent: curr.Val})
		}
	}

	if m.depth == n.depth {
		return m.parent != n.parent
	} else {
		return false
	}
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	fmt.Println(cousins(root, 3, 4))
}
