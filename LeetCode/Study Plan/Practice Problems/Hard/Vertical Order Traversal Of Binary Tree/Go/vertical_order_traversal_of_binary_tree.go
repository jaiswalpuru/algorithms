package main

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type Pair struct {
	val int
	lvl int
}
type P []Pair

func (p P) Len() int { return len(p) }
func (p P) Less(i, j int) bool {
	if p[i].lvl == p[j].lvl {
		return p[i].val < p[j].val
	}
	return p[i].lvl < p[j].lvl
}
func (p P) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func vertical_order_traversal(root *TreeNode) [][]int {
	res := [][]int{}
	mp := make(map[int]P)
	min_val, max_val := math.MaxInt64, math.MinInt64
	traverse(root, &mp, 0, 0, &min_val, &max_val)
	for i := min_val; i <= max_val; i++ {
		val := mp[i]
		sort.Sort(val)
		v := make([]int, len(val))
		for j := 0; j < len(val); j++ {
			v[j] = val[j].val
		}
		res = append(res, v)
	}
	return res
}

func traverse(root *TreeNode, mp *map[int]P, lvl, r int, min_val, max_val *int) {
	if root == nil {
		return
	}
	(*mp)[lvl] = append((*mp)[lvl], Pair{root.Val, r})
	*min_val = min(*min_val, lvl)
	*max_val = max(*max_val, lvl)
	traverse(root.Left, mp, lvl-1, r+1, min_val, max_val)
	traverse(root.Right, mp, lvl+1, r+1, min_val, max_val)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Right = New(3)
	root.Right.Right = New(7)
	root.Right.Left = New(6)
	fmt.Println(vertical_order_traversal(root))
}
