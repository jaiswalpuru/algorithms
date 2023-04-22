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
	lvl  int
}

func max_lvl_sum(root *TreeNode) int {

	q := []Pair{}
	q = append(q, Pair{node: root, lvl: 1})
	mp := make(map[int]int)
	for len(q) > 0 {
		curr := q[0].node
		lvl := q[0].lvl
		q = q[1:]
		mp[lvl] += curr.Val
		if curr.Left != nil {
			q = append(q, Pair{node: curr.Left, lvl: lvl + 1})
		}
		if curr.Right != nil {
			q = append(q, Pair{node: curr.Right, lvl: lvl + 1})
		}
	}
	max_val := 1
	for k, v := range mp {
		if v > mp[max_val] {
			max_val = k
		}
	}
	return max_val
}

func main() {
	root := New(-100)
	root.Right = New(-300)
	root.Left = New(-200)
	root.Right.Left = New(-10)
	root.Left.Left = New(-20)
	root.Left.Right = New(-5)
	fmt.Println(max_lvl_sum(root))
}
