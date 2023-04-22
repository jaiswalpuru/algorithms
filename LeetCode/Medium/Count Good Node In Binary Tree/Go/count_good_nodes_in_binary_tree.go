package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

//----------------------------dfs-------------------------------
type Pair struct {
	node       *TreeNode
	max_so_far int
}

func count_good_nodes_in_binary_tree_dfs(root *TreeNode) int {
	cnt := 0
	preorder(root, math.MinInt64, &cnt)
	return cnt
}

func preorder(root *TreeNode, max_so_far int, cnt *int) {
	if max_so_far <= root.Val {
		*cnt++
	}
	if root.Left != nil {
		preorder(root.Left, max(max_so_far, root.Val), cnt)
	}
	if root.Right != nil {
		preorder(root.Right, max(max_so_far, root.Val), cnt)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//----------------------------dfs-------------------------------

//----------------------------bfs-------------------------------
func count_good_nodes_in_binary_tree_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cnt := 1
	q := []Pair{{root, root.Val}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.node.Left != nil {
			if curr.max_so_far <= curr.node.Left.Val {
				cnt++
			}
			q = append(q, Pair{curr.node.Left, max(curr.max_so_far, curr.node.Left.Val)})
		}
		if curr.node.Right != nil {
			if curr.max_so_far <= curr.node.Right.Val {
				cnt++
			}
			q = append(q, Pair{curr.node.Right, max(curr.max_so_far, curr.node.Right.Val)})
		}
	}
	return cnt
}

//----------------------------bfs-------------------------------

func main() {
	root := New(3)
	root.Left = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(2)
	fmt.Println(count_good_nodes_in_binary_tree_bfs(root))
}
