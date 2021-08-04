/*
Given the root of a Binary Search Tree (BST), return the minimum difference between the values
of any two different nodes in the tree.

Example 1:
Input: root = [4,2,6,1,3]
Output: 1

Example 2:
Input: root = [1,0,48,null,null,12,49]
Output: 1
*/

package main

import (
	"fmt"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func min_diff(root *Node) int {
	res := []int{}
	inorder(root, &res)
	diff := abs(res[0] - res[1])
	n := len(res)
	for i := 2; i < n; i++ {
		diff = min(diff, abs(res[i]-res[i-1]))
	}
	return diff
}

func inorder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(6)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	fmt.Println(min_diff(root))
}
