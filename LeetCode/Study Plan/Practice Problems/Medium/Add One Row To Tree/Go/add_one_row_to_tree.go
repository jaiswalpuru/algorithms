package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func add_one_row(root *TreeNode, val int, d int) *TreeNode {
	if d == 1 {
		node := New(val)
		node.Left = root
		return node
	}

	depth := 1
	q := []*TreeNode{root}

	for depth < d-1 {
		temp := []*TreeNode{}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			if curr.Left != nil {
				temp = append(temp, curr.Left)
			}
			if curr.Right != nil {
				temp = append(temp, curr.Right)
			}
		}
		q = temp
		depth++
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		temp := curr.Left
		curr.Left = New(val)
		curr.Left.Left = temp
		temp = curr.Right
		curr.Right = New(val)
		curr.Right.Right = temp
	}
	return root
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(5)
	root.Left.Left = New(3)
	root.Left.Right = New(1)
	root.Right.Left = New(8)
	root.Right.Right = New(9)
	fmt.Println(add_one_row(root, 1, 3).Right.Right.Right)
}
