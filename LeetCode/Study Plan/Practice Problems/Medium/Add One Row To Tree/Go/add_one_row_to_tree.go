package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

type Pair struct {
	node  *TreeNode
	depth int
}

func add_one_row(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := New(val, root, nil)
		return node
	}
	q := []Pair{Pair{root, 1}}
	f := false
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			v := q[i]
			if v.depth+1 == depth {
				f = true
				n1 := New(val, v.node.Left, nil)
				n2 := New(val, nil, v.node.Right)
				v.node.Left = n1
				v.node.Right = n2
			} else {
				if v.node.Left != nil {
					q = append(q, Pair{v.node.Left, v.depth + 1})
				}
				if v.node.Right != nil {
					q = append(q, Pair{v.node.Right, v.depth + 1})
				}
			}
		}
		if f {
			break
		}
		q = q[n:]
	}
	return root
}

func main() {
	root := New(4, nil, nil)
	root.Left = New(2, nil, nil)
	root.Right = New(5, nil, nil)
	root.Left.Left = New(3, nil, nil)
	root.Left.Right = New(1, nil, nil)
	root.Right.Left = New(8, nil, nil)
	root.Right.Right = New(9, nil, nil)
	fmt.Println(add_one_row(root, 1, 3))
}
