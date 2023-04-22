package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func find_lca(root *TreeNode, p, q int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p || root.Val == q {
		return root
	}

	left := find_lca(root.Left, p, q)
	right := find_lca(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

func find_distance_bfs(root *TreeNode, p, q int) int {
	head := find_lca(root, p, q)

	type Pair struct {
		node *TreeNode
		val  int
	}

	qu := []Pair{{head, 0}}
	a, b := -1, -1
	for len(qu) > 0 {
		curr := qu[0].node
		dis := qu[0].val
		qu = qu[1:]
		if curr.Val == p {
			a = dis
		}
		if curr.Val == q {
			b = dis
		}
		if a != -1 && b != -1 {
			break
		}
		if curr.Left != nil {
			qu = append(qu, Pair{node: curr.Left, val: dis + 1})
		}
		if curr.Right != nil {
			qu = append(qu, Pair{node: curr.Right, val: dis + 1})
		}
	}

	return a + b
}

func main() {
	root := New(3)
	root.Left = New(5)
	root.Right = New(1)
	root.Left.Left = New(6)
	root.Left.Right = New(2)
	root.Left.Right.Left = New(7)
	root.Left.Right.Right = New(4)
	root.Right.Right = New(8)
	root.Right.Left = New(0)
	fmt.Println(find_distance_bfs(root, 5, 0))
}
