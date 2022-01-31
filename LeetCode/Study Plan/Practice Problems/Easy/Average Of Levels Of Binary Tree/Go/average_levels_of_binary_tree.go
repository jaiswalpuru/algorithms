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

type Pair struct {
	node *TreeNode
	lvl  int
}

func average_levels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	ans := []float64{}
	mp := make(map[int]int)
	node_num := make(map[int]int)
	q := []Pair{{node: root, lvl: 0}}
	for len(q) > 0 {
		cur := q[0].node
		lvl := q[0].lvl
		q = q[1:]
		mp[lvl] += cur.Val
		node_num[lvl]++
		if cur.Left != nil {
			q = append(q, Pair{node: cur.Left, lvl: lvl + 1})
		}

		if cur.Right != nil {
			q = append(q, Pair{node: cur.Right, lvl: lvl + 1})
		}
	}
	for i := 0; i < len(mp); i++ {
		ans = append(ans, float64(mp[i])/float64(node_num[i]))
	}
	return ans
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(average_levels(root))
}
