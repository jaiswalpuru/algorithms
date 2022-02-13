package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func generate_trees(start, end int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if start > end {
		trees = append(trees, nil)
		return trees
	}

	for i := start; i <= end; i++ {
		left := generate_trees(start, i-1)
		right := generate_trees(i+1, end)

		for _, l := range left {
			for _, r := range right {
				root := New(i)
				root.Left = l
				root.Right = r
				trees = append(trees, root)
			}
		}
	}
	return trees
}

func unique_bst(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate_trees(1, n)
}

func main() {
	fmt.Println(unique_bst(3))
}
