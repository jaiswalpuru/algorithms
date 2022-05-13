package main

import "fmt"

type Node struct {
	Val int
	Left,
	Right,
	Next *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil, Next: nil} }

func populating_next_right_pointers_in_each_node(root *Node) *Node {
	if root == nil {
        return nil
    }
	q := []*Node{root}
	for len(q) > 0 {
		n := len(q)
		if n == 1 {
			curr := q[0]
			q = q[1:]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		} else {
			for i := 0; i < n-1; i++ {
				curr := q[0]
				q = q[1:]
				curr.Next = q[0]
				if curr.Left != nil {
					q = append(q, curr.Left)
				}

				if curr.Right != nil {
					q = append(q, curr.Right)
				}
			}
			curr := q[0]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
			q = q[1:]
		}
	}
	return root
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Right.Right = New(7)
	fmt.Println(populating_next_right_pointers_in_each_node(root))
}
