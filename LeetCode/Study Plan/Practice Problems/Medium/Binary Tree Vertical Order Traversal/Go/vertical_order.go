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

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type Pair struct {
	node *TreeNode
	d    int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func vertical_order_traversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	mp := make(map[int][]int)
	q := []Pair{{root, 0}}
	l_min, r_max := math.MaxInt64, math.MinInt64
	for len(q) > 0 {
		curr := q[0].node
		d := q[0].d
		q = q[1:]

		mp[d] = append(mp[d], curr.Val)
		l_min = min(l_min, d)
		r_max = max(r_max, d)
		if curr.Left != nil {
			q = append(q, Pair{curr.Left, d - 1})
		}

		if curr.Right != nil {
			q = append(q, Pair{curr.Right, d + 1})
		}
	}

	res := [][]int{}
	for i := l_min; i <= r_max; i++ {
		if _, ok := mp[i]; ok {
			res = append(res, mp[i])
		}
	}
	return res
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(vertical_order_traversal(root))
}
