package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	order := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		qSize := len(q)
		temp := []int{}
		for i := 0; i < qSize; i++ {
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
		order = append(order, temp)
	}

	size := len(order)
	//reverse the order
	for i := 0; i < size/2; i++ {
		order[i], order[size-i-1] = order[size-1-i], order[i]
	}

	return order
}

func main() {
	root := New(3)
	root.Left = New(9)
	root.Right = New(20)
	root.Right.Left = New(15)
	root.Right.Right = New(7)
	fmt.Println(levelOrderBottom(root))
}
