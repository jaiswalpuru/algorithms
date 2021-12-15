/*
This problem was asked by Facebook.

Given a binary tree, return the level of the tree with minimum sum.
*/

package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func New(data int) *Node { return &Node{Val: data, Left: nil, Right: nil} }

func level_with_min_sum(root *Node) int {
	if root == nil {
		return 0
	}
	q := []*Node{root}
	sum := int(math.MaxInt32)
	for len(q) > 0 {
		n := len(q)
		temp := 0
		for i := 0; i < n; i++ {
			temp += q[0].Val
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			q = q[1:]
		}
		sum = min(sum, temp)
	}
	return sum
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	root := New(10)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(15)
	root.Right.Right = New(20)
	fmt.Println(level_with_min_sum(root))
}
