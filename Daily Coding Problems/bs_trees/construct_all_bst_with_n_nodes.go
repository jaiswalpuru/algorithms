/*
	Given an integer n construct all possible binary search trees with n nodes where all values from [1...n]
	are used.
*/

package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data} }

func make_trees(low, high int) []*Node {
	trees := []*Node{}
	if low > high {
		trees = append(trees, nil)
		return trees
	}
	for i := low; i < high+1; i++ {
		left := make_trees(low, i-1)
		right := make_trees(i+1, high)
		for _, v := range left {
			for _, k := range right {
				node := New(i)
				node.left = v
				node.right = k
				trees = append(trees, node)
			}
		}
	}
	return trees
}

func pre_order(root *Node) {
	if root != nil {
		fmt.Print(root.data, " ")
		pre_order(root.left)
		pre_order(root.right)
	}
}

func main() {
	trees := make_trees(1, 3)
	for i := range trees {
		pre_order(trees[i])
		fmt.Println()
	}
}
