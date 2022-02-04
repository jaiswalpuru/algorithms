package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func leaf_similar_trees(r1, r2 *TreeNode) bool {
	l1, l2 := []int{}, []int{}
	traverse(r1, &l1)
	traverse(r2, &l2)
	if len(l1) != len(l2) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func traverse(r *TreeNode, path *[]int) {
	if r == nil {
		return
	}

	if r.Left == nil && r.Right == nil {
		*path = append(*path, r.Val)
	}
	traverse(r.Left, path)
	traverse(r.Right, path)
}

func main() {
	r1 := New(3)
	r1.Left = New(2)
	r1.Right = New(1)
	r2 := New(5)
	r2.Left = New(2)
	r2.Right = New(1)
	fmt.Println(leaf_similar_trees(r1, r2))
}
