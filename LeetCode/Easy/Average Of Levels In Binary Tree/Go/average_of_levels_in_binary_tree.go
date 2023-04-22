package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type Pair struct {
	node *TreeNode
	pos  int
}

//----------------------------bfs---------------------------------------
func average_of_levels_in_binary_tree(root *TreeNode) []float64 {
	mp := make(map[int][]int)
	q := []Pair{{root, 0}}
	lvl := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		lvl = curr.pos
		mp[curr.pos] = append(mp[curr.pos], curr.node.Val)
		if curr.node.Left != nil {
			q = append(q, Pair{curr.node.Left, curr.pos + 1})
		}
		if curr.node.Right != nil {
			q = append(q, Pair{curr.node.Right, curr.pos + 1})
		}
	}
	res := make([]float64, lvl+1)
	for k, v := range mp {
		sum := 0
		for i := 0; i < len(v); i++ {
			sum += v[i]
		}
		res[k] = float64(sum) / float64(len(v))
	}
	return res
}

//----------------------------bfs---------------------------------------

//----------------------------dfs---------------------------------------
func average_of_levels_in_binary_tree_dfs(root *TreeNode) []float64 {
	mp := make(map[int][]int)
	lvl := 0
	dfs(root, &mp, 0, &lvl)
	res := make([]float64, lvl)
	for k, v := range mp {
		sum := 0
		for i := 0; i < len(v); i++ {
			sum += v[i]
		}
		res[k] = float64(sum) / float64(len(v))
	}
	return res
}

func dfs(root *TreeNode, mp *map[int][]int, lvl int, l *int) {
	if root == nil {
		*l = max(*l, lvl)
		return
	}
	(*mp)[lvl] = append((*mp)[lvl], root.Val)
	dfs(root.Left, mp, lvl+1, l)
	dfs(root.Right, mp, lvl+1, l)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//----------------------------dfs---------------------------------------

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(average_of_levels_in_binary_tree_dfs(root))
}
