package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Right: nil, Left: nil} }

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	even := true
	q := []*TreeNode{root}
	res := [][]int{}
	for len(q) > 0 {
		temp := []int{}
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]
			temp = append(temp, curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		if !even {
			reverse(&temp)
		}
		even = !even
		res = append(res, temp)
	}
	return res
}

func reverse(t *[]int) {
	n := len(*t)
	for i := 0; i < n/2; i++ {
		(*t)[i], (*t)[n-i-1] = (*t)[n-1-i], (*t)[i]
	}
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(zigzagLevelOrder(root))
}
