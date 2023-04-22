package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type Pair struct {
	node *TreeNode
	lvl  int
}

func find_nearest_right_binary_tree(root *TreeNode, u *TreeNode) *TreeNode {
	q := []Pair{{node: root, lvl: 0}}
	for len(q) > 0 {
		curr := q[0].node
		lvl := q[0].lvl
		q = q[1:]
		
		if curr == u {
			if len(q) > 0 && q[0].lvl == lvl {
				return q[0].node
			}
			return nil
		}

		if curr.Left != nil {
			q = append(q, Pair{curr.Left, lvl + 1})
		}

		if curr.Right != nil {
			q = append(q, Pair{curr.Right, lvl + 1})
		}

	}
	return nil
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Right = New(4)
	root.Right = New(3)
	root.Right.Left = New(5)
	root.Right.Right = New(6)
	fmt.Println(find_nearest_right_binary_tree(root, root.Left.Right))
}
