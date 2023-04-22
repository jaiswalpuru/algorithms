package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func closest_leaf_in_binary_tree(root *TreeNode, k int) int {
	graph := make(map[*TreeNode][]*TreeNode)
	dfs(&graph, root, nil)

	q := []*TreeNode{}
	visited := make(map[*TreeNode]bool)

	for key := range graph {
		if key != nil && key.Val == k {
			q = append(q, key)
			visited[key] = true
			break
		}
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr != nil {
			if len(graph[curr]) <= 1 {
				return curr.Val
			}
			for i := 0; i < len(graph[curr]); i++ {
				if !visited[graph[curr][i]] {
					q = append(q, graph[curr][i])
					visited[graph[curr][i]] = true
				}
			}
		}
	}

	return -1
}

func dfs(graph *map[*TreeNode][]*TreeNode, root *TreeNode, parent *TreeNode) {
	if root != nil {
		(*graph)[root] = append((*graph)[root], parent)
		(*graph)[parent] = append((*graph)[parent], root)
		dfs(graph, root.Left, root)
		dfs(graph, root.Right, root)
	}
}

func main() {
	root := New(1)
	root.Left = New(3)
	root.Right = New(2)
	fmt.Println(closest_leaf_in_binary_tree(root, 2))
}
