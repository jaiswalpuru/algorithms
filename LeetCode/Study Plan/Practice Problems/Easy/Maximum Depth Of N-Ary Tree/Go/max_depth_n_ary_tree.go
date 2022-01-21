package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: make([]*Node, 0)} }

type Pair struct {
	node  *Node
	depth int
}

func depth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []Pair{{node: root, depth: 1}}

	for len(queue) > 0 {
		curr := queue[0].node
		depth = queue[0].depth
		queue = queue[1:]
		n := len(curr.Children)
		for i := 0; i < n; i++ {
			queue = append(queue, Pair{node: curr.Children[i], depth: depth + 1})
		}
	}
	return depth
}

//using recursion
func max_depth(root *Node) int {
	if root == nil {
		return 0
	} else if len(root.Children) == 0 {
		return 1
	} else {
		heights := []int{}
		for i := 0; i < len(root.Children); i++ {
			heights = append(heights, max_depth(root.Children[i]))
		}
		max_val := 0
		for j := 0; j < len(heights); j++ {
			max_val = max(max_val, heights[j])
		}
		return max_val + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := New(1)
	root.Children = append(root.Children, New(3))
	root.Children = append(root.Children, New(2))
	root.Children = append(root.Children, New(4))
	root.Children[0].Children = append(root.Children[0].Children, New(5))
	root.Children[0].Children = append(root.Children[0].Children, New(6))
	fmt.Println(depth(root))
}
