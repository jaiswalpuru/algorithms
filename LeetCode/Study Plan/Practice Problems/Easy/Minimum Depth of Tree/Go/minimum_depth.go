package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type Pair struct {
	node  *TreeNode
	depth int
}

func minimum_depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []Pair{{node: root, depth: 1}}
	depth := 0
	for len(queue) != 0 {
		curr := queue[0].node
		depth = queue[0].depth
		queue = queue[1:]
		if curr.Right == nil && curr.Left == nil {
			return depth
		}
		if curr.Left != nil {
			queue = append(queue, Pair{node: curr.Left, depth: depth + 1})
		}
		if curr.Right != nil {
			queue = append(queue, Pair{node: curr.Right, depth: depth + 1})
		}
	}
	return depth
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(minimum_depth(root))
}
